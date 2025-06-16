package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Receiver interface {
	// (GET /ocpi/locations/{country_code}/{party_id}/{location_id})
	GetLocation(ctx context.Context, countryCode string, partyId string, locationId string) (*Location, error)
	// (GET /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	GetLocationEVSE(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string) (*EVSE, error)
	// (GET /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	GetLocationConnector(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, connectorId string) (*Connector, error)
	// (PUT /ocpi/locations/{country_code}/{party_id}/{location_id})
	PutLocation(ctx context.Context, countryCode string, partyId string, locationId string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PutLocationEVSE(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PutLocationConnector(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, connectorId string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/locations/{country_code}/{party_id}/{location_id})
	PatchLocation(ctx context.Context, countryCode string, partyId string, locationId string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PatchLocationEVSE(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PatchLocationConnector(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, connectorId string, body ocpi.RawMessage[PatchedLocation]) error

	// (GET /ocpi/tariffs/{country_code}/{party_id}/{tariff_id})
	GetTariff(ctx context.Context, countryCode string, partyId string, sessionId string) (*Tariff, error)
	// (PUT /ocpi/tariffs/{country_code}/{party_id}/{tariff_id})
	PutTariff(ctx context.Context, countryCode string, partyId string, tariffId string, body ocpi.RawMessage[Tariff]) error
	// (PATCH /ocpi/tariffs/{country_code}/{party_id}/{tariff_id})
	PatchTariff(ctx context.Context, countryCode string, partyId string, tariffId string, body ocpi.RawMessage[Tariff]) error

	// (GET /ocpi/sessions/{country_code}/{party_id}/{session_id})
	GetSession(ctx context.Context, countryCode string, partyId string, sessionId string) (*Session, error)
	// (PUT /ocpi/sessions/{country_code}/{party_id}/{session_id})
	PutSession(ctx context.Context, countryCode string, partyId string, sessionId string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/sessions/{country_code}/{party_id}/{session_id})
	PatchSession(ctx context.Context, countryCode string, partyId string, sessionId string, body ocpi.RawMessage[PatchedSession]) error

	// (GET /ocpi/tokens/{country_code}/{party_id}/{token_uid})
	GetToken(ctx context.Context, countryCode string, partyId string, tokenUid string, tokenType ...TokenType) (*Token, error)
	// (PUT /ocpi/tokens/{country_code}/{party_id}/{token_uid})
	PutToken(ctx context.Context, countryCode string, partyId string, tokenUid string, body ocpi.RawMessage[Token]) error
	// (PATCH /ocpi/tokens/{country_code}/{party_id}/{token_uid})
	PatchToken(ctx context.Context, countryCode string, partyId string, tokenUid string, body ocpi.RawMessage[PatchedToken]) error

	// (GET /ocpi/cdrs/{cdr_id})
	GetCDR(ctx context.Context, cdrId string) (*ChargeDetailRecord, error)
	// (POST /ocpi/cdrs)
	PostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) error

	// (GET /ocpi/terminals/{terminal_id})
	GetTerminal(ctx context.Context, terminalId string) (*Terminal, error)
	// (GET /ocpi/financialadviceconfirmations/{financial_advice_confirmation_id})
	GetFinancialAdviceConfirmation(ctx context.Context, financialAdviceConfirmationId string) (*FinancialAdviceConfirmation, error)
	// (POST /ocpi/terminals)
	PostTerminal(ctx context.Context, body ocpi.RawMessage[Terminal]) error
	// (POST /ocpi/financialadviceconfirmations)
	PostFinancialAdviceConfirmation(ctx context.Context, body ocpi.RawMessage[FinancialAdviceConfirmation]) error
}

func (s *Server) GetOcpiVersionDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpoints := []Endpoint{
		{Identifier: ModuleIDCredentials, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/credentials"},
		{Identifier: ModuleIDCredentials, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/credentials"},
	}
	if s.cdrsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/cdrs"})
	}
	if s.cdrsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/cdrs"})
	}
	if s.chargingProfilesSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/chargingprofiles"})
	}
	if s.chargingProfilesReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/chargingprofiles"})
	}
	if s.commandsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/commands"})
	}
	if s.commandsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/commands"})
	}
	if s.hubClientInfoSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/clientinfo"})
	}
	if s.hubClientInfoReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/clientinfo"})
	}
	if s.locationsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/locations"})
	}
	if s.locationsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/locations"})
	}
	if s.sessionsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/sessions"})
	}
	if s.sessionsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/sessions"})
	}
	if s.tariffsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/tariffs"})
	}
	if s.tariffsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/tariffs"})
	}
	if s.tokensSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/tokens"})
	}
	if s.tokensReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/tokens"})
	}

	b, err := json.Marshal(ocpi.NewResponse(endpoints))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	locationId := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorId := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	var (
		resp any
		err  error
	)
	if evseUid != "" && connectorId != "" {
		resp, err = s.receiver.GetLocationConnector(ctx, countryCode, partyId, locationId, evseUid, connectorId)
	} else if evseUid != "" {
		resp, err = s.receiver.GetLocationEVSE(ctx, countryCode, partyId, locationId, evseUid)
	} else {
		resp, err = s.receiver.GetLocation(ctx, countryCode, partyId, locationId)
	}
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(resp))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PutOcpiLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	ctx := r.Context()
	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	locationId := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorId := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	s.logger.DebugContext(ctx, "PutOcpiLocation",
		"country_code", countryCode,
		"party_id", partyId,
		"location_id", locationId,
		"evse_uid", evseUid,
		"connector_id", connectorId)

	if evseUid != "" && connectorId != "" {
		if err := s.receiver.PutLocationConnector(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			connectorId,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.GenericServerError)
			return
		}
	} else if evseUid != "" {
		if err := s.receiver.PutLocationEVSE(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.GenericServerError)
			return
		}
	} else {
		if err := s.receiver.PutLocation(
			ctx,
			countryCode,
			partyId,
			locationId,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.GenericServerError)
			return
		}
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PatchOcpiLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	ctx := r.Context()
	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	locationId := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorId := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	if evseUid != "" && connectorId != "" {
		if err := s.receiver.PatchLocationConnector(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			connectorId,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.GenericServerError)
			return
		}
	} else if evseUid != "" {
		if err := s.receiver.PatchLocationEVSE(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.GenericServerError)
			return
		}
	} else {
		if err := s.receiver.PatchLocation(
			ctx,
			countryCode,
			partyId,
			locationId,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.GenericServerError)
			return
		}
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	sessionId := chi.URLParam(r, "session_id")

	session, err := s.receiver.GetSession(
		r.Context(),
		countryCode,
		partyId,
		sessionId,
	)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(session))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PutOcpiSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	ctx := r.Context()
	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	sessionId := chi.URLParam(r, "session_id")

	if err := s.receiver.PutSession(
		ctx,
		countryCode,
		partyId,
		sessionId,
		ocpi.RawMessage[Session](body),
	); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PatchOcpiSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	ctx := r.Context()
	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	sessionId := chi.URLParam(r, "session_id")

	if err := s.receiver.PatchSession(
		ctx,
		countryCode,
		partyId,
		sessionId,
		ocpi.RawMessage[PatchedSession](body),
	); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCdr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	ctx := r.Context()
	if err := s.receiver.PostCDR(ctx, ocpi.RawMessage[ChargeDetailRecord](body)); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	commandType := chi.URLParam(r, "command_type")

	resp, err := s.commandsReceiver.PostCommand(r.Context(), CommandType(commandType))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(resp))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
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
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	if err := s.commandsSender.PostAsyncCommand(r.Context(), CommandType(commandType), uid, ocpi.RawMessage[CommandResult](body)); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiCDR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	cdr, err := s.receiver.GetCDR(r.Context(), id)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(cdr))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCDR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	if err := s.receiver.PostCDR(r.Context(), ocpi.RawMessage[ChargeDetailRecord](body)); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse[any](nil))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
