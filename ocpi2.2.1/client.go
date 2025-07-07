package ocpi221

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"unsafe"

	"github.com/gofrs/uuid/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type EndpointResolver func(endpoint string) string

type OCPIClient interface {
	GetCredentialsTokenC(ctx context.Context) (string, error)
	GetEndpoint(ctx context.Context, mod ModuleID, role InterfaceRole) (string, error)
}

type ClientTokenA interface {
	GetVersions(ctx context.Context) (*ocpi.Response[ocpi.Versions], error)
	GetVersionDetails(ctx context.Context, version ocpi.Version) (*ocpi.Response[VersionDetails], error)
}

type Client interface {
	CallEndpoint(ctx context.Context, mod ModuleID, role InterfaceRole, method string, endpointResolver EndpointResolver, src, dst any) error
	GetLocations(ctx context.Context, params ...GetLocationsParams) (*ocpi.PaginationResponse[Location], error)
	GetLocation(ctx context.Context, locationID string) (*ocpi.Response[Location], error)
	GetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*ocpi.Response[Location], error)
	PutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, location Location) (*ocpi.Response[any], error)
	PatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, location PartialLocation) (*ocpi.Response[any], error)
	GetSessions(ctx context.Context, params ...GetSessionsParams) (*ocpi.PaginationResponse[Session], error)
	GetSession(ctx context.Context, sessionID string) (*ocpi.Response[Session], error)
	GetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*ocpi.Response[Session], error)
	PutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session Session) (*ocpi.Response[any], error)
	PatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session PartialSession) (*ocpi.Response[any], error)
	StartSession(ctx context.Context, req StartSession) (*ocpi.Response[CommandResponse], error)
	StopSession(ctx context.Context, req StopSession) (*ocpi.Response[CommandResponse], error)
	ReserveNow(ctx context.Context, req ReserveNow) (*ocpi.Response[CommandResponse], error)
	CancelReservation(ctx context.Context, req CancelReservation) (*ocpi.Response[CommandResponse], error)
	UnlockConnector(ctx context.Context, req UnlockConnector) (*ocpi.Response[CommandResponse], error)
	SetSessionChargingPreferences(ctx context.Context, sessionID string) (*ocpi.Response[ChargingPreferencesResponse], error)
}

type ClientConn struct {
	tokenA     string
	ocpi       OCPIClient
	versionUrl string
	httpClient *http.Client
}

var _ Client = (*ClientConn)(nil)

func NewClient(versionUrl string, ocpi OCPIClient) *ClientConn {
	c := new(ClientConn)
	c.versionUrl = versionUrl
	c.ocpi = ocpi
	c.httpClient = &http.Client{}
	return c
}

func NewClientWithTokenA(versionUrl string, tokenA string) ClientTokenA {
	c := new(ClientConn)
	c.tokenA = tokenA
	c.versionUrl = versionUrl
	c.httpClient = &http.Client{}
	return c
}

func (c *ClientConn) CallEndpoint(ctx context.Context, mod ModuleID, role InterfaceRole, method string, resolver EndpointResolver, src, dst any) error {
	endpoint, err := c.ocpi.GetEndpoint(ctx, mod, role)
	if err != nil {
		return err
	}

	tokenC, err := c.ocpi.GetCredentialsTokenC(ctx)
	if err != nil {
		return err
	}
	if err := c.do(ctx, tokenC, method, resolver(endpoint), src, dst); err != nil {
		return err
	}
	return nil
}

func (c *ClientConn) do(ctx context.Context, token, method, endpoint string, src, dst any) error {
	var body io.Reader
	if src != nil {
		b, err := json.Marshal(src)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set(ocpi.HttpHeaderXRequestID, uuid.Must(uuid.NewV7()).String())
	reqCtx := GetRequestContext(ctx)
	if reqCtx.FromCountryCode != "" && reqCtx.FromPartyID != "" {
		req.Header.Set(ocpi.HttpHeaderOCPIFromCountryCode, reqCtx.FromCountryCode)
		req.Header.Set(ocpi.HttpHeaderOCPIFromPartyID, reqCtx.FromPartyID)
	}
	if reqCtx.ToCountryCode != "" && reqCtx.ToPartyID != "" {
		req.Header.Set(ocpi.HttpHeaderOCPIToCountryCode, reqCtx.ToCountryCode)
		req.Header.Set(ocpi.HttpHeaderOCPIToPartyID, reqCtx.ToPartyID)
	}
	if reqCtx.requestID != "" {
		req.Header.Set(ocpi.HttpHeaderXCorrelationID, reqCtx.requestID)
	} else {
		req.Header.Set(ocpi.HttpHeaderXCorrelationID, uuid.Must(uuid.NewV7()).String())
	}

	req.Header.Set("Authorization", "Token "+base64.StdEncoding.EncodeToString(unsafe.Slice(unsafe.StringData(token), len(token))))

	b, _ := httputil.DumpRequest(req, true)
	log.Println(string(b))
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, _ = httputil.DumpResponse(res, true)
	log.Println(string(b))
	if scanner, ok := dst.(ocpi.HeaderScanner); ok {
		if err := scanner.ScanHeader(res.Header); err != nil {
			return fmt.Errorf(`ocpi221: unable to scan header: %w`, err)
		}
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf(`ocpi221: encounter status code (%d) due to %s`, res.StatusCode, unsafe.String(unsafe.SliceData(b), len(b)))
	}
	return json.NewDecoder(res.Body).Decode(dst)
}
