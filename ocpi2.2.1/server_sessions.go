package ocpi221

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiSessions(w http.ResponseWriter, r *http.Request) {
	params := GetSessionsParams{}
	response, err := s.sessionsSender.OnGetSessions(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
}

func (s *Server) PutOcpiSesionChargingPreferences(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	ctx := r.Context()
	sessionID := r.PathValue("session_id")
	chargingPreferences, err := s.sessionsSender.OnPutSessionChargingPreferences(
		ctx,
		sessionID,
		ocpi.RawMessage[ChargingPreferences](body),
	)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(chargingPreferences))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiSession(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	sessionID := r.PathValue("session_id")

	session, err := s.sessionsReceiver.OnGetClientOwnedSession(r.Context(), countryCode, partyID, sessionID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(session))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
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

	if err := s.sessionsReceiver.OnPutClientOwnedSession(ctx, countryCode, partyID, sessionID, ocpi.RawMessage[Session](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
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

	if err := s.sessionsReceiver.OnPatchClientOwnedSession(ctx, countryCode, partyID, sessionID, ocpi.RawMessage[PartialSession](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
