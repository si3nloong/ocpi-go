package ocpi211

import (
	"context"
	"net/http"
)

func (c *ClientConn) GetLocation(ctx context.Context, locationID string) (*LocationResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations)
	if err != nil {
		return nil, err
	}

	var res LocationResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+locationID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
