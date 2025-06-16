package ocpi221

import (
	"context"
	"net/http"
)

func (c *Client) StartSession(ctx context.Context, req StartSessionRequest) (*CommandResponse, error) {
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

func (c *Client) StopSession(ctx context.Context, req StopSessionRequest) (*CommandResponse, error) {
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

func (c *Client) ReserveNow(ctx context.Context, req ReserveNowRequest) (*CommandResponse, error) {
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

func (c *Client) CancelReservation(ctx context.Context, req CancelReservation) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeCancelReservation), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) UnlockConnector(ctx context.Context, req UnlockConnector) (*CommandResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res CommandResponse
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeUnlockConnector), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
