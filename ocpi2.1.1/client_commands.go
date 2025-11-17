package ocpi211

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) StartSession(ctx context.Context, req StartSession) (*ocpi.Response[DateTime, CommandResponse], error) {
	var res ocpi.Response[DateTime, CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeStartSession)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) StopSession(ctx context.Context, req StopSession) (*ocpi.Response[DateTime, CommandResponse], error) {
	var res ocpi.Response[DateTime, CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeStopSession)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) ReserveNow(ctx context.Context, req ReserveNow) (*ocpi.Response[DateTime, CommandResponse], error) {
	var res ocpi.Response[DateTime, CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeReserveNow)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) UnlockConnector(ctx context.Context, req UnlockConnector) (*ocpi.Response[DateTime, CommandResponse], error) {
	var res ocpi.Response[DateTime, CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeUnlockConnector)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
