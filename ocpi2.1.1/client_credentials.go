package ocpi211

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetCredential(ctx context.Context) (*ocpi.Response[Credentials], error) {
	var res ocpi.Response[Credentials]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, http.MethodGet, func(endpoint string) string {
		return endpoint
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PostCredential(ctx context.Context, req Credentials) (*ocpi.Response[Credentials], error) {
	var res ocpi.Response[Credentials]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, http.MethodPost, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
