package ocpi211

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
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
			ocpihttp.Response(w, err)
			return
		}
		params.DateFrom = dt
	}
	if queryString.Has("date_to") {
		dt, err := ParseDateTime(queryString.Get("date_to"))
		if err != nil {
			ocpihttp.Response(w, err)
			return
		}
		params.DateTo = &dt
	}
	if queryString.Has("offset") {
		offset, err := strconv.ParseUint(queryString.Get("offset"), 10, 32)
		if err != nil {
			ocpihttp.Response(w, err)
			return
		}
		params.Offset = &offset
	}
	if queryString.Has("limit") {
		limit, err := strconv.ParseUint(queryString.Get("limit"), 10, 32)
		if err != nil {
			ocpihttp.Response(w, err)
			return
		}
		u16 := uint16(limit)
		params.Limit = &u16
	}

	response, err := s.cpo.OnGetSessions(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiSession(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	session, err := s.emsp.OnGetClientOwnedSession(r.Context(), countryCode, partyID, sessionID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, session)
}

func (s *Server) PutOcpiSession(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	if err := s.emsp.OnPutClientOwnedSession(ctx, countryCode, partyID, sessionID, ocpi.RawMessage[Session](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PatchOcpiSession(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	if err := s.emsp.OnPatchClientOwnedSession(ctx, countryCode, partyID, sessionID, ocpi.RawMessage[PartialSession](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
