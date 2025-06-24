package ocpi221

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
	CallEndpoint(ctx context.Context, mod ModuleID, role InterfaceRole, endpointResolver EndpointResolver, src, dst any) error
	GetTariffs(ctx context.Context, params ...GetTariffsParams) (ocpi.Result[[]Tariff], error)
	GetLocations(ctx context.Context, params ...GetLocationsParams) (ocpi.Result[[]Location], error)
	GetLocation(ctx context.Context, locationID string) (*LocationResponse, error)
	GetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*LocationResponse, error)
	PutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, loc Location) error
	PatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, loc PatchedLocation) error
	GetSessions(ctx context.Context, params ...GetSessionsParams) (*SessionsResponse, error)
	GetSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*SessionResponse, error)
	StartSession(ctx context.Context, req StartSession) (*CommandResponse, error)
	StopSession(ctx context.Context, req StopSession) (*CommandResponse, error)
	ReserveNow(ctx context.Context, req ReserveNow) (*CommandResponse, error)
	CancelReservation(ctx context.Context, req CancelReservation) (*CommandResponse, error)
	UnlockConnector(ctx context.Context, req UnlockConnector) (*CommandResponse, error)
	SetSessionChargingPreferences(ctx context.Context, sessionID string) (*ocpi.Response[ChargingPreferencesResponse], error)
}

type Option func(*client)

type client struct {
	rw           sync.RWMutex
	tokenA       string
	tokenC       string
	versionUrl   string
	httpClient   *http.Client
	endpointDict map[string]Endpoint
}

func WithTokenC(tokenC string) Option {
	return func(c *client) {
		c.tokenC = tokenC
	}
}

func NewClient(versionUrl string, options ...Option) Client {
	c := new(client)
	c.versionUrl = versionUrl
	c.httpClient = &http.Client{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func (c *client) CallEndpoint(
	ctx context.Context,
	mod ModuleID,
	role InterfaceRole,
	endpointResolver EndpointResolver,
	src, dst any,
) error {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return err
	}

	if err := c.do(ctx, http.MethodGet, endpointResolver(endpoint), src, dst); err != nil {
		return err
	}
	return nil
}

func (c *client) getEndpoint(ctx context.Context, mod ModuleID, role InterfaceRole) (string, error) {
	c.rw.RLock()
	if c.endpointDict == nil {
		c.rw.RUnlock()
		versions, err := c.Versions(ctx)
		if err != nil {
			return "", err
		}

		if len(versions) == 0 {
			return "", fmt.Errorf("ocpi221: no versions found at %s", c.versionUrl)
		}

		version := versions[0]
		var o ocpi.Response[VersionDetails]
		if err := c.do(ctx, http.MethodGet, version.URL, nil, &o); err != nil {
			return "", err
		}

		c.rw.Lock()
		c.endpointDict = make(map[string]Endpoint)
		for _, v := range o.Data.Endpoints {
			c.endpointDict[string(v.Identifier)+":"+string(v.Role)] = v
		}
		c.rw.Unlock()
		c.rw.RLock()
	}
	defer c.rw.RUnlock()
	v, ok := c.endpointDict[string(mod)+":"+string(role)]
	if ok {
		return v.URL, nil
	}
	return "", fmt.Errorf(`ocpi221: missing endpoint for module id %q (%s)`, mod, role)
}

func (c *client) newRequest(
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

func (c *client) do(
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
		return fmt.Errorf(`ocpi221: encounter status code (%d) due to %s`, res.StatusCode, unsafe.String(unsafe.SliceData(b), len(b)))
	}
	return json.NewDecoder(res.Body).Decode(dst)
}
