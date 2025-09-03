package ocpi230

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetHubClientInfos(ctx context.Context, params ...GetHubClientInfoParams) (*ocpi.PaginationResponse[ClientInfo], error) {
	query := make(url.Values)
	query.Set("limit", "100")
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
	var res ocpi.Response[[]ClientInfo]
	if err := c.CallEndpoint(ctx, ModuleIDHubClientInfo, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &ocpi.PaginationResponse[ClientInfo]{
		Response: res,
	}, nil
}

func (c *ClientConn) GetClientInfo(ctx context.Context, countryCode, partyID string) (*ocpi.Response[ClientInfo], error) {
	var res ocpi.Response[ClientInfo]
	if err := c.CallEndpoint(ctx, ModuleIDHubClientInfo, InterfaceRoleReceiver, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientInfo(ctx context.Context, countryCode, partyID string) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDHubClientInfo, InterfaceRoleReceiver, http.MethodPut, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
