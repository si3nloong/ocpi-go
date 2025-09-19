package ocpi230

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetTariffs(ctx context.Context, params ...GetTariffsParams) (*ocpi.PaginationResponse[Tariff], error) {
	query := make(url.Values)
	query.Set("limit", "100")
	if len(params) > 0 {
		p := params[0]
		if p.DateFrom != nil && p.DateFrom.IsZero() {
			query.Set("date_from", p.DateFrom.String())
		}
		if p.DateTo != nil && p.DateTo.IsZero() {
			query.Set("date_to", p.DateTo.String())
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Set("offset", strconv.Itoa(*p.Offset))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Set("limit", strconv.Itoa(*p.Limit))
		}
	}
	var res ocpi.Response[[]Tariff]
	if err := c.CallEndpoint(ctx, ModuleIDTariffs, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &ocpi.PaginationResponse[Tariff]{
		Response: res,
	}, nil
}

func (c *ClientConn) GetClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string) (*ocpi.Response[Tariff], error) {
	var res ocpi.Response[Tariff]
	if err := c.CallEndpoint(ctx, ModuleIDTariffs, InterfaceRoleReceiver, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + tariffID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string, tariff Tariff) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDTariffs, InterfaceRoleReceiver, http.MethodPut, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + tariffID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) DeleteClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDTariffs, InterfaceRoleReceiver, http.MethodDelete, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + tariffID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
