package ocpi221

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) RegisterCredentials(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) UpdateCredentials(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutCredentials(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
