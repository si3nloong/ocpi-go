package ocpi211

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (c *ClientConn) GetSessions(ctx context.Context, dateFrom time.Time, params ...GetSessionsParams) (*SessionsResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions)
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
		if p.DateTo != nil && !p.DateTo.IsZero() {
			query.Add("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Add("offset", strconv.FormatUint(uint64(*p.Offset), 10))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Add("limit", strconv.FormatUint(uint64(*p.Limit), 10))
		}
	}
	u.RawQuery = query.Encode()

	var res SessionsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*SessionResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions)
	if err != nil {
		return nil, err
	}

	var res SessionResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+sessionID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
