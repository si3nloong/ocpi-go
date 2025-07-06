package ocpi221

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetLocations(ctx context.Context, params ...GetLocationsParams) (*ocpi.PaginationResponse[Location], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
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

	var o ocpi.Response[[]Location]
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return &ocpi.PaginationResponse[Location]{
		Response: o,
	}, nil
}

func (c *ClientConn) GetLocation(ctx context.Context, locationID string) (*ocpi.Response[Location], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[Location]
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+locationID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*ocpi.Response[Location], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[Location]
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+locationID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, location Location) (*ocpi.Response[any], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[any]
	if err := c.do(ctx, http.MethodPut, endpoint+"/"+countryCode+"/"+partyID+"/"+locationID, location, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, location PartialLocation) (*ocpi.Response[any], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[any]
	if err := c.do(ctx, http.MethodPatch, endpoint+"/"+countryCode+"/"+partyID+"/"+locationID, location, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
