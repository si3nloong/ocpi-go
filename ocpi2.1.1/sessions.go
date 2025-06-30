package ocpi211

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type GetSessionsParams struct {
	DateFrom DateTime
	DateTo   *DateTime
	Offset   *uint64
	Limit    *uint16
}

func (s *Server) GetOcpiSessions(w http.ResponseWriter, r *http.Request) {

	params := GetSessionsParams{}
	queryString := r.URL.Query()
	if queryString.Has("date_from") {
		dt, err := ParseDateTime(queryString.Get("date_from"))
		if err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
		params.DateFrom = dt
	}
	if queryString.Has("date_to") {
		dt, err := ParseDateTime(queryString.Get("date_to"))
		if err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
		params.DateTo = &dt
	}
	if queryString.Has("offset") {
		offset, err := strconv.ParseUint(queryString.Get("offset"), 10, 32)
		if err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
		params.Offset = &offset
	}
	if queryString.Has("limit") {
		limit, err := strconv.ParseUint(queryString.Get("limit"), 10, 32)
		if err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
		u16 := uint16(limit)
		params.Limit = &u16
	}

	response, err := s.cpo.OnGetSessions(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiSession(w http.ResponseWriter, r *http.Request) {

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	session, err := s.emsp.OnGetClientOwnedSession(
		r.Context(),
		countryCode,
		partyID,
		sessionID,
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

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	if err := s.emsp.OnPutClientOwnedSession(
		ctx,
		countryCode,
		partyID,
		sessionID,
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

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	if err := s.emsp.OnPatchClientOwnedSession(
		ctx,
		countryCode,
		partyID,
		sessionID,
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

func (c *ClientConn) GetSessions(
	ctx context.Context,
	dateFrom time.Time,
	params ...GetSessionsParams,
) (*SessionsResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions)
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
		if p.DateTo != nil && !p.DateTo.IsZero() {
			query.Add("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Add("offset", strconv.FormatUint(uint64(*p.Offset), 10))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Add("limit", strconv.FormatUint(uint64(*p.Limit), 10))
		}
	}
	u.RawQuery = query.Encode()

	var o SessionsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (c *ClientConn) GetSession(
	ctx context.Context,
	countryCode string,
	partyID string,
	sessionID string,
) (*SessionResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDSessions)
	if err != nil {
		return nil, err
	}

	var o SessionResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+sessionID, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
