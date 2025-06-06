package ocpi221

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"unsafe"

	"github.com/samber/lo"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Option func(*Client)

type Client struct {
	rw           sync.RWMutex
	tokenA       string
	tokenC       string
	versionUrl   string
	httpClient   *http.Client
	endpointDict map[string]DetailsDataEndpoints
}

func WithTokenC(tokenC string) Option {
	return func(c *Client) {
		c.tokenC = tokenC
	}
}

func NewClient(versionUrl string, options ...Option) *Client {
	c := new(Client)
	c.versionUrl = versionUrl
	c.httpClient = &http.Client{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func (c *Client) getEndpoint(ctx context.Context, mod ModuleIDType, role InterfaceRoleType) (string, error) {
	c.rw.RLock()
	if c.endpointDict == nil {
		c.rw.RUnlock()
		versions, err := c.Versions(ctx)
		if err != nil {
			return "", err
		}

		version, _ := lo.First(versions)
		var o ocpi.Response[DetailsData]
		if err := c.do(ctx, http.MethodGet, version.Url, nil, &o); err != nil {
			return "", err
		}

		c.rw.Lock()
		c.endpointDict = make(map[string]DetailsDataEndpoints)
		for _, v := range o.Data.Endpoints {
			c.endpointDict[string(v.Identifier)+":"+string(v.Role)] = v
		}
		c.rw.Unlock()
		c.rw.RLock()
	}
	defer c.rw.RUnlock()
	v, ok := c.endpointDict[string(mod)+":"+string(role)]
	if ok {
		return v.Url, nil
	}
	return "", fmt.Errorf(`ocpi: missing endpoint for module id %q (%s)`, mod, role)
}

func (c *Client) newRequest(
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

func (c *Client) do(
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

	dir := filepath.Base(endpoint)
	f, err := os.OpenFile("./examples/"+dir, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	b, _ := io.ReadAll(res.Body)
	f.Write(b)
	// b, _ := httputil.DumpResponse(res, true)
	// log.Println(string(b))
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf(`ocpi221: encounter status code (%d) due to %s`, res.StatusCode, unsafe.String(unsafe.SliceData(b), len(b)))
	}
	return json.NewDecoder(bytes.NewBuffer(b)).Decode(dst)
}
