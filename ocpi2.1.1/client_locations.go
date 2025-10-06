package ocpi211

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetLocations(ctx context.Context, params ...GetLocationsParams) (*ocpi.PaginationResponse[Location], error) {
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
	var res ocpi.Response[[]Location]
	if err := c.CallEndpoint(ctx, ModuleIDLocations, http.MethodGet, func(endpoint string) string {
		return endpoint + "?" + query.Encode()
	}, nil, &res); err != nil {
		return nil, err
	}
	return &ocpi.PaginationResponse[Location]{
		Response: res,
	}, nil
}

func (c *ClientConn) GetLocation(ctx context.Context, locationID string) (*ocpi.Response[Location], error) {
	var res ocpi.Response[Location]
	if err := c.CallEndpoint(ctx, ModuleIDLocations, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + locationID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*ocpi.Response[Location], error) {
	var res ocpi.Response[Location]
	if err := c.CallEndpoint(ctx, ModuleIDLocations, http.MethodGet, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + locationID
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, location Location) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDLocations, http.MethodPut, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + locationID
	}, location, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, location PartialLocation) (*ocpi.Response[any], error) {
	var res ocpi.Response[any]
	if err := c.CallEndpoint(ctx, ModuleIDLocations, http.MethodPatch, func(endpoint string) string {
		return endpoint + "/" + countryCode + "/" + partyID + "/" + locationID
	}, location, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
