package ocpi221

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) StartSession(ctx context.Context, req StartSession) (*ocpi.Response[CommandResponse], error) {
	endpoint, err := c.ocpi.GetEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	req.Token.LastUpdated = DateTime{Time: req.Token.LastUpdated.UTC()}
	reqCtx := GetRequestContext(ctx)
	reqCtx.FromCountryCode = req.Token.CountryCode
	reqCtx.FromPartyID = req.Token.PartyID

	var res ocpi.Response[CommandResponse]
	if err := c.do(WithRequestContext(ctx, reqCtx), http.MethodPost, endpoint+"/"+string(CommandTypeStartSession), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) StopSession(ctx context.Context, req StopSession) (*ocpi.Response[CommandResponse], error) {
	endpoint, err := c.ocpi.GetEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[CommandResponse]
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeStopSession), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) ReserveNow(ctx context.Context, req ReserveNow) (*ocpi.Response[CommandResponse], error) {
	endpoint, err := c.ocpi.GetEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	req.Token.LastUpdated = DateTime{Time: req.Token.LastUpdated.UTC()}

	var res ocpi.Response[CommandResponse]
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeReserveNow), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) CancelReservation(ctx context.Context, req CancelReservation) (*ocpi.Response[CommandResponse], error) {
	endpoint, err := c.ocpi.GetEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}
	var res ocpi.Response[CommandResponse]
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeCancelReservation), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) UnlockConnector(ctx context.Context, req UnlockConnector) (*ocpi.Response[CommandResponse], error) {
	endpoint, err := c.ocpi.GetEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[CommandResponse]
	if err := c.do(ctx, http.MethodPost, endpoint+"/"+string(CommandTypeUnlockConnector), req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
