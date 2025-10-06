package ocpi230

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetHubClientInfos(ctx context.Context, params ...GetHubClientInfoParams) (*ocpi.PaginatedResponse[ClientInfo], error) {
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
	var res ocpi.PaginatedResponse[ClientInfo]
	if err := c.CallEndpoint(ctx, ModuleIDHubClientInfo, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
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
