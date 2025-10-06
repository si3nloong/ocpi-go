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

	"github.com/gofrs/uuid/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type EndpointResolver func(endpoint string) string

type OCPIClient interface {
	GetCredentialsToken(ctx context.Context) (string, error)
	GetEndpoint(ctx context.Context, module ModuleID) (string, error)
}

type TokenAClient interface {
	GetVersions(ctx context.Context) (*ocpi.Response[ocpi.Versions], error)
	GetVersionDetails(ctx context.Context) (*ocpi.Response[VersionDetails], error)
	GetCredential(ctx context.Context) (*ocpi.Response[Credentials], error)
	PostCredential(ctx context.Context, req Credentials) (*ocpi.Response[Credentials], error)
}

type Client interface {
	TokenAClient
	CallEndpoint(ctx context.Context, module ModuleID, method string, endpointResolver EndpointResolver, src, dst any) error
	GetLocation(ctx context.Context, locationID string) (*ocpi.Response[Location], error)
	GetSessions(ctx context.Context, dateFrom DateTime, params ...GetSessionsParams) (*ocpi.PaginatedResponse[Session], error)
	GetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*ocpi.Response[Session], error)
	PutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session Session) (*ocpi.Response[any], error)
	PatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session PartialSession) (*ocpi.Response[any], error)
	StartSession(ctx context.Context, req StartSession) (*ocpi.Response[CommandResponse], error)
	StopSession(ctx context.Context, req StopSession) (*ocpi.Response[CommandResponse], error)
	ReserveNow(ctx context.Context, req ReserveNow) (*ocpi.Response[CommandResponse], error)
	UnlockConnector(ctx context.Context, req UnlockConnector) (*ocpi.Response[CommandResponse], error)
	GetCDRs(ctx context.Context, params ...GetCDRsParams) (*ocpi.PaginatedResponse[CDR], error)
	GetCDR(ctx context.Context, cdrID string) (*ocpi.Response[CDR], error)
	PostCDR(ctx context.Context, endpoint string, req CDR) (*ocpi.Response[ChargeDetailRecordResponse], error)
}

var defaultClientOptions = ClientOptions{
	HttpClient: http.DefaultClient,
}

type ClientOptions struct {
	HttpClient *http.Client
}

type ClientConn struct {
	rw             sync.RWMutex
	ocpi           OCPIClient
	versionDetails *VersionDetails
	versions       ocpi.Versions
	versionUrl     string
	httpClient     *http.Client
}

var _ Client = (*ClientConn)(nil)

func NewClient(versionUrl string, ocpi OCPIClient, opts *ClientOptions) *ClientConn {
	if opts == nil {
		opts = &defaultClientOptions
	}
	c := new(ClientConn)
	c.versionUrl = versionUrl
	c.ocpi = ocpi
	if opts.HttpClient == nil {
		opts.HttpClient = &http.Client{}
	}
	c.httpClient = opts.HttpClient
	return c
}

func NewClientWithTokenA(versionUrl string, tokenA string, opts *ClientOptions) TokenAClient {
	c := &unregisteredClient{tokenA: tokenA}
	client := NewClient(versionUrl, c, opts)
	c.ClientConn = client
	return c
}

func (c *ClientConn) CallEndpoint(ctx context.Context, module ModuleID, method string, resolver EndpointResolver, src, dst any) error {
	endpoint, err := c.ocpi.GetEndpoint(ctx, module)
	if err != nil {
		return err
	}

	if err := c.do(ctx, method, resolver(endpoint), src, dst); err != nil {
		return err
	}
	return nil
}

func (c *ClientConn) do(ctx context.Context, method, endpoint string, src, dst any) error {
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

	token, err := c.ocpi.GetCredentialsToken(ctx)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set(ocpi.HttpHeaderXRequestID, uuid.Must(uuid.NewV7()).String())
	reqCtx := ocpi.GetRequestContext(ctx)
	// if reqCtx.FromCountryCode != "" && reqCtx.FromPartyID != "" {
	// 	req.Header.Set(ocpi.HttpHeaderOCPIFromCountryCode, reqCtx.FromCountryCode)
	// 	req.Header.Set(ocpi.HttpHeaderOCPIFromPartyID, reqCtx.FromPartyID)
	// }
	// if reqCtx.ToCountryCode != "" && reqCtx.ToPartyID != "" {
	// 	req.Header.Set(ocpi.HttpHeaderOCPIToCountryCode, reqCtx.ToCountryCode)
	// 	req.Header.Set(ocpi.HttpHeaderOCPIToPartyID, reqCtx.ToPartyID)
	// }
	if reqCtx.RequestID != "" {
		req.Header.Set(ocpi.HttpHeaderXCorrelationID, reqCtx.RequestID)
	} else {
		req.Header.Set(ocpi.HttpHeaderXCorrelationID, uuid.Must(uuid.NewV7()).String())
	}

	req.Header.Set("Authorization", "Token "+token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if scanner, ok := dst.(ocpi.HeaderScanner); ok {
		if err := scanner.ScanHeader(res.Header); err != nil {
			return fmt.Errorf(`ocpi211: unable to scan headers: %w`, err)
		}
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf(`ocpi211: encounter status code (%d) due to %s`, res.StatusCode, unsafe.String(unsafe.SliceData(b), len(b)))
	}
	return json.NewDecoder(res.Body).Decode(dst)
}
