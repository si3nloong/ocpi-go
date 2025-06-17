package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetOcpiSessionsParams{}
	response, err := s.sessionsSender.GetSessions(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	writePaginationResponse(w, r, response)
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
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(session))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PutOcpiSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
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

func (s *Server) PatchOcpiSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
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

type GetSessionsParams struct {
	DateTo time.Time
	Offset uint32
	Limit  uint8
}

func (c *Client) GetSessions(
	ctx context.Context,
	dateFrom time.Time,
	params ...GetSessionsParams,
) (*SessionsResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("date_from", dateFrom.Format(time.RFC3339))
	if len(params) > 0 {
		p := params[0]
		if !p.DateTo.IsZero() {
			query.Add("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset > 0 {
			query.Add("offset", strconv.FormatUint(uint64(p.Offset), 10))
		}
		if p.Limit > 0 {
			query.Add("limit", strconv.FormatUint(uint64(p.Limit), 10))
		}
	}
	u.RawQuery = query.Encode()

	var o SessionsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *Client) GetSession(
	ctx context.Context,
	countryCode string,
	partyId string,
	sessionId string,
) (*SessionResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var o SessionResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyId+"/"+sessionId, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *Client) SetSessionChargingPreferences(ctx context.Context, sessionId string) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions, InterfaceRoleReceiver)
	if err != nil {
		return nil, err
	}

	var o SessionResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/sessions/"+sessionId+"/charging_preferences", nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
