package ocpi211

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"unsafe"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type EndpointResolver func(endpoint string) string

type Client interface {
	CallEndpoint(ctx context.Context, mod ModuleID, endpointResolver EndpointResolver, src, dst any) error
}

type Option func(*ClientConn)

type ClientConn struct {
	rw           sync.RWMutex
	tokenA       string
	tokenC       string
	versionUrl   string
	httpClient   *http.Client
	endpointDict map[string]Endpoint
}

var _ Client = (*ClientConn)(nil)

func WithTokenC(tokenC string) Option {
	return func(c *ClientConn) {
		c.tokenC = tokenC
	}
}

func NewClient(versionUrl string, options ...Option) *ClientConn {
	c := new(ClientConn)
	c.versionUrl = versionUrl
	c.httpClient = &http.Client{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func (c *ClientConn) CallEndpoint(
	ctx context.Context,
	mod ModuleID,
	endpointResolver EndpointResolver,
	src, dst any,
) error {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations)
	if err != nil {
		return err
	}

	if err := c.do(ctx, http.MethodGet, endpointResolver(endpoint), src, dst); err != nil {
		return err
	}
	return nil
}

func (c *ClientConn) getEndpoint(ctx context.Context, mod ModuleID) (string, error) {
	c.rw.RLock()
	if c.endpointDict == nil {
		c.rw.RUnlock()
		versions, err := c.Versions(ctx)
		if err != nil {
			return "", err
		}

		if len(versions) == 0 {
			return "", fmt.Errorf("ocpi211: no versions found at %s", c.versionUrl)
		}

		version := versions[0]
		var o ocpi.Response[VersionDetails]
		if err := c.do(ctx, http.MethodGet, version.URL, nil, &o); err != nil {
			return "", err
		}

		c.rw.Lock()
		c.endpointDict = make(map[string]Endpoint)
		for _, v := range o.Data.Endpoints {
			c.endpointDict[string(v.Identifier)] = v
		}
		c.rw.Unlock()
		c.rw.RLock()
	}
	defer c.rw.RUnlock()
	v, ok := c.endpointDict[string(mod)]
	if ok {
		return v.URL, nil
	}
	return "", fmt.Errorf(`ocpi211: missing endpoint for module id %q`, mod)
}

func (c *ClientConn) newRequest(
	ctx context.Context,
	method, endpoint string,
	src any,
) (*http.Request, error) {
	var body io.Reader
	if src != nil {
		b, err := json.Marshal(src)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	c.rw.RLock()
	if c.tokenC != "" {
		req.Header.Set("Authorization", "Token "+c.tokenC)
	} else {
		req.Header.Set("Authorization", "Token "+c.tokenA)
	}
	c.rw.RUnlock()
	return req, nil
}

func (c *ClientConn) do(
	ctx context.Context,
	method, endpoint string,
	src, dst any,
) error {
	req, err := c.newRequest(ctx, method, endpoint, src)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf(`ocpi211: encounter status code (%d) due to %s`, res.StatusCode, unsafe.String(unsafe.SliceData(b), len(b)))
	}
	return json.NewDecoder(res.Body).Decode(dst)
}
