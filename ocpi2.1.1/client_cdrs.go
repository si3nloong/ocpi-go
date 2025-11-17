package ocpi211

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetCDRs(ctx context.Context, params ...GetCDRsParams) (*ocpi.PaginatedResponse[DateTime, CDR], error) {
	query := make(url.Values)
	query.Set("limit", "100")
	if len(params) > 0 {
		p := params[0]
		if p.DateFrom != nil && !p.DateFrom.IsZero() {
			query.Set("date_from", p.DateFrom.String())
		}
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
	var res ocpi.PaginatedResponse[DateTime, CDR]
	if err := c.CallEndpoint(ctx, ModuleIDCdrs, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetCDR(ctx context.Context, cdrID string) (*ocpi.Response[DateTime, CDR], error) {
	var res ocpi.Response[DateTime, CDR]
	if err := c.CallEndpoint(ctx, ModuleIDCdrs, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + cdrID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PostCDR(ctx context.Context, endpoint string, req CDR) (*ocpi.Response[DateTime, ChargeDetailRecordResponse], error) {
	var res ocpi.Response[DateTime, ChargeDetailRecordResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCdrs, http.MethodPost, func(_ string) string {
		return endpoint
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
