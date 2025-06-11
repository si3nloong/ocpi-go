package ocpi221

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type GetSessionsParams struct {
	DateTo time.Time
	Offset uint32
	Limit  uint8
}

func (c *Client) GetSessions(
	ctx context.Context,
	dateFrom time.Time,
	params ...GetSessionsParams,
) (*SessionsResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("date_from", dateFrom.Format(time.RFC3339))
	if len(params) > 0 {
		p := params[0]
		if !p.DateTo.IsZero() {
			query.Add("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset > 0 {
			query.Add("offset", strconv.FormatUint(uint64(p.Offset), 10))
		}
		if p.Limit > 0 {
			query.Add("limit", strconv.FormatUint(uint64(p.Limit), 10))
		}
	}
	u.RawQuery = query.Encode()

	var o SessionsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *Client) GetSession(
	ctx context.Context,
	countryCode string,
	partyId string,
	sessionId string,
) (*SessionResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var o SessionResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyId+"/"+sessionId, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *Client) SetSessionChargingPreferences(ctx context.Context, sessionId string) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var o SessionResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/sessions/"+sessionId+"/charging_preferences", nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
