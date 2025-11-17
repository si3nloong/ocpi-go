package ocpi221

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetSessions(ctx context.Context, dateFrom DateTime, params ...GetSessionsParams) (*ocpi.PaginatedResponse[DateTime, Session], error) {
	query := make(url.Values)
	query.Set("date_from", dateFrom.String())
	if len(params) > 0 {
		p := params[0]
		if p.DateTo != nil && !p.DateTo.IsZero() {
			query.Set("date_to", p.DateTo.String())
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Set("offset", strconv.Itoa(*p.Offset))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Set("limit", strconv.Itoa(*p.Limit))
		}
	}
	var res ocpi.PaginatedResponse[DateTime, Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetSession(ctx context.Context, sessionID string) (*ocpi.Response[DateTime, Session], error) {
	var res ocpi.Response[DateTime, Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + sessionID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) SetSessionChargingPreferences(ctx context.Context, sessionID string) (*ocpi.Response[DateTime, ChargingPreferencesResponse], error) {
	var res ocpi.Response[DateTime, ChargingPreferencesResponse]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "/sessions/" + sessionID + "/charging_preferences"
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*ocpi.Response[DateTime, Session], error) {
	var res ocpi.Response[DateTime, Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session Session) (*ocpi.Response[DateTime, any], error) {
	var res ocpi.Response[DateTime, any]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver, http.MethodPut, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session PartialSession) (*ocpi.Response[DateTime, any], error) {
	var res ocpi.Response[DateTime, any]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver, http.MethodPatch, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
