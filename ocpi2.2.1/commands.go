package ocpi221

import (
	"context"
	"net/http"
)

func (c *Client) StartSessionCommands(ctx context.Context, req StartSessionRequest) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	req.Token.LastUpdated = req.Token.LastUpdated.UTC()
	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStartSession), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) StopSessionCommands(ctx context.Context, req StopSessionRequest) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStopSession), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ReserveNowCommands(ctx context.Context, req ReserveNowRequest) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeReserveNow), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
