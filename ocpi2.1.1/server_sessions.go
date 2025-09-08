package ocpi211

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiSessions(w http.ResponseWriter, r *http.Request) {
	params := GetSessionsParams{}
	queryString := r.URL.Query()
	if !queryString.Has("date_from") {
		ocpihttp.BadRequest(w, r, `missing "date_from" parameter`)
		return
	}
	dateFrom, err := ParseDateTime(queryString.Get("date_from"))
	if err != nil {
		ocpihttp.Response(w, err)
		return
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
		offset, err := strconv.Atoi(queryString.Get("offset"))
		if err != nil {
			ocpihttp.Response(w, err)
			return
		}
		params.Offset = &offset
	}
	if queryString.Has("limit") {
		limit, err := strconv.Atoi(queryString.Get("limit"))
		if err != nil {
			ocpihttp.Response(w, err)
			return
		}
		params.Limit = &limit
	}

	response, err := s.cpo.OnGetSessions(r.Context(), dateFrom, params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
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
		ocpihttp.BadRequest(w, r, err.Error())
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
		ocpihttp.BadRequest(w, r, err.Error())
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
