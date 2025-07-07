package ocpi221

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetSessions(ctx context.Context, params ...GetSessionsParams) (*ocpi.PaginationResponse[Session], error) {
	query := make(url.Values)
	if len(params) > 0 {
		p := params[0]
		if p.DateFrom != nil && p.DateFrom.IsZero() {
			query.Set("date_from", p.DateFrom.Format(time.RFC3339))
		}
		if p.DateTo != nil && p.DateTo.IsZero() {
			query.Set("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Set("offset", strconv.FormatUint(uint64(*p.Offset), 10))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Set("limit", strconv.FormatUint(uint64(*p.Limit), 10))
		}
	}
	var res ocpi.PaginationResponse[Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetSession(ctx context.Context, sessionID string) (*ocpi.Response[Session], error) {
	var res ocpi.Response[Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + sessionID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) SetSessionChargingPreferences(ctx context.Context, sessionID string) (*ocpi.Response[ChargingPreferencesResponse], error) {
	var res ocpi.Response[ChargingPreferencesResponse]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "/sessions/" + sessionID + "/charging_preferences"
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*ocpi.Response[Session], error) {
	var res ocpi.Response[Session]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session Session) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver, http.MethodPut, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session PartialSession) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver, http.MethodPatch, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
