package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) PostOcpiCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	commandType := chi.URLParam(r, "command_type")

	resp, err := s.commandsReceiver.PostCommand(r.Context(), CommandType(commandType))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(resp))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCommandResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	commandType := chi.URLParam(r, "command_type")
	uid := chi.URLParam(r, "uid")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	if err := s.commandsSender.PostAsyncCommand(r.Context(), CommandType(commandType), uid, ocpi.RawMessage[CommandResult](body)); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *Client) StartSession(ctx context.Context, req StartSession) (*CommandResponse, error) {
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

func (c *Client) StopSession(ctx context.Context, req StopSession) (*CommandResponse, error) {
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

func (c *Client) ReserveNow(ctx context.Context, req ReserveNow) (*CommandResponse, error) {
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
