package ocpi211

import (
	"context"
	"net/http"
)

func (c *ClientConn) StartSession(ctx context.Context, req StartSession) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands)
	if err != nil {
		return nil, err
	}
	req.Token.LastUpdated = DateTime{Time: req.Token.LastUpdated.UTC()}

	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStartSession), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) StopSession(ctx context.Context, req StopSession) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands)
	if err != nil {
		return nil, err
	}

	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStopSession), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) ReserveNow(ctx context.Context, req ReserveNow) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands)
	if err != nil {
		return nil, err
	}

	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeReserveNow), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) UnlockConnector(ctx context.Context, req UnlockConnector) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands)
	if err != nil {
		return nil, err
	}

	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeUnlockConnector), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
