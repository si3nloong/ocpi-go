package ocpi221

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetTariffs(ctx context.Context, params ...GetTariffsParams) (*ocpi.PaginationResponse[Tariff], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
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
	u.RawQuery = query.Encode()
	var o ocpi.Response[[]Tariff]
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return &ocpi.PaginationResponse[Tariff]{
		Response: o,
	}, nil
}

func (c *ClientConn) GetClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string) (*ocpi.Response[Tariff], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[Tariff]
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+tariffID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string, tariff Tariff) (*ocpi.Response[any], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[any]
	if err := c.do(ctx, http.MethodPut, endpoint+"/"+countryCode+"/"+partyID+"/"+tariffID, tariff, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) DeleteClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string) (*ocpi.Response[any], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[any]
	if err := c.do(ctx, http.MethodDelete, endpoint+"/"+countryCode+"/"+partyID+"/"+tariffID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
