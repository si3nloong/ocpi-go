package ocpi211

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetSessions(ctx context.Context, dateFrom DateTime, params ...GetSessionsParams) (*ocpi.PaginationResponse[Session], error) {
	query := make(url.Values)
	query.Add("date_from", dateFrom.String())
	if len(params) > 0 {
		p := params[0]
		if p.DateTo != nil && !p.DateTo.IsZero() {
			query.Add("date_to", p.DateTo.String())
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Add("offset", strconv.Itoa(*p.Offset))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Add("limit", strconv.Itoa(*p.Limit))
		}
	}

	var res ocpi.PaginationResponse[Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*ocpi.Response[Session], error) {
	var res ocpi.Response[Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session Session) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, http.MethodPut, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session PartialSession) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, http.MethodPatch, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
