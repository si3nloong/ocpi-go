package ocpi221

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) StartSession(ctx context.Context, req StartSession) (*ocpi.Response[CommandResponse], error) {
	req.Token.LastUpdated = DateTime{Time: req.Token.LastUpdated.UTC()}
	reqCtx := GetRequestContext(ctx)
	reqCtx.FromCountryCode = req.Token.CountryCode
	reqCtx.FromPartyID = req.Token.PartyID

	var res ocpi.Response[CommandResponse]
	if err := c.CallEndpoint(WithRequestContext(ctx, reqCtx), ModuleIDCommands, InterfaceRoleReceiver, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeStartSession)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) StopSession(ctx context.Context, req StopSession) (*ocpi.Response[CommandResponse], error) {
	var res ocpi.Response[CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeStopSession)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) ReserveNow(ctx context.Context, req ReserveNow) (*ocpi.Response[CommandResponse], error) {
	req.Token.LastUpdated = DateTime{Time: req.Token.LastUpdated.UTC()}
	var res ocpi.Response[CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeReserveNow)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) CancelReservation(ctx context.Context, req CancelReservation) (*ocpi.Response[CommandResponse], error) {
	var res ocpi.Response[CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeCancelReservation)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) UnlockConnector(ctx context.Context, req UnlockConnector) (*ocpi.Response[CommandResponse], error) {
	var res ocpi.Response[CommandResponse]
	if err := c.CallEndpoint(ctx, ModuleIDCommands, InterfaceRoleReceiver, http.MethodPost, func(endpoint string) string {
		return endpoint + "/" + string(CommandTypeUnlockConnector)
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
