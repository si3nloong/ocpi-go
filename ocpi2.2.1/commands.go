package ocpi221

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (c *Client) StartSessionCommands(ctx context.Context, req StartSessionRequest) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTypeCommands, RoleReceiver)
	if err != nil {
		return nil, err
	}

	req.Token.LastUpdated = req.Token.LastUpdated.UTC()
	var res json.RawMessage
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStartSession), req, &res); err != nil {
		return nil, err
	}
	return string(res), nil
}

func (c *Client) StopSessionCommands(ctx context.Context, req StopSessionRequest) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTypeCommands, RoleReceiver)
	if err != nil {
		return nil, err
	}
	var res json.RawMessage
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStopSession), req, &res); err != nil {
		return nil, err
	}
	return string(res), nil
}

func (c *Client) ReserveNowCommands(ctx context.Context, req ReserveNowRequest) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTypeCommands, RoleReceiver)
	if err != nil {
		return nil, err
	}
	var res json.RawMessage
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeReserveNow), req, &res); err != nil {
		return nil, err
	}
	log.Println(string(res))
	return res, nil
}
