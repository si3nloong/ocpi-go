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
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
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
	u.RawQuery = query.Encode()
	var o ocpi.PaginationResponse[Session]
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *ClientConn) GetSession(ctx context.Context, sessionID string) (*ocpi.Response[Session], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}
	var o ocpi.Response[Session]
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+sessionID, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *ClientConn) SetSessionChargingPreferences(ctx context.Context, sessionID string) (*ocpi.Response[ChargingPreferencesResponse], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var o ocpi.Response[ChargingPreferencesResponse]
	if err := c.do(ctx, http.MethodGet, endpoint+"/sessions/"+sessionID+"/charging_preferences", nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *ClientConn) GetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*ocpi.Response[Session], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res ocpi.Response[Session]
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+sessionID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session Session) (*ocpi.Response[any], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res ocpi.Response[any]
	if err := c.do(ctx, http.MethodPut, endpoint+"/"+countryCode+"/"+partyID+"/"+sessionID, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, session PartialSession) (*ocpi.Response[any], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res ocpi.Response[any]
	if err := c.do(ctx, http.MethodPatch, endpoint+"/"+countryCode+"/"+partyID+"/"+sessionID, session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
