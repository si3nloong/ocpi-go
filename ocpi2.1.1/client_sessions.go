package ocpi211

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetSessions(ctx context.Context, dateFrom time.Time, params ...GetSessionsParams) (*ocpi.PaginationResponse[Session], error) {
	query := make(url.Values)
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
	if err := c.CallEndpoint(ctx, ModuleIDTariffs, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + sessionID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
