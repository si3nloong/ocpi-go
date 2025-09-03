package ocpi230

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiActiveChargingProfile(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("duration") {
		ocpihttp.BadRequest(w, r, `missing required query string "duration"`)
		return
	}
	if !r.URL.Query().Has("response_url") {
		ocpihttp.BadRequest(w, r, `missing required query string "response_url"`)
		return
	}
	sessionID := r.PathValue("session_id")
	duration, err := strconv.Atoi(r.URL.Query().Get("duration"))
	if err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}
	responseUrl := strings.TrimSpace(r.URL.Query().Get("response_url"))
	response, err := s.chargingProfilesReceiver.OnGetActiveChargingProfile(r.Context(), sessionID, duration, responseUrl)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, response)
}

func (s *Server) PutOcpiChargingProfile(w http.ResponseWriter, r *http.Request) {
	var body ocpi.RawMessage[SetChargingProfile]
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	sessionID := r.PathValue("session_id")
	response, err := s.chargingProfilesReceiver.OnPutChargingProfile(r.Context(), sessionID, body)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, response)
}

func (s *Server) DeleteOcpiChargingProfile(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("response_url") {
		ocpihttp.BadRequest(w, r, `missing required query string "response_url"`)
		return
	}

	sessionID := r.PathValue("session_id")
	responseUrl := strings.TrimSpace(r.URL.Query().Get("response_url"))
	response, err := s.chargingProfilesReceiver.OnDeleteChargingProfile(r.Context(), sessionID, responseUrl)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, response)
}

func (s *Server) PostOcpiActiveChargingProfile(w http.ResponseWriter, r *http.Request) {
	var body ocpi.RawMessage[ActiveChargingProfileResult]
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	sessionID := r.PathValue("session_id")
	if err := s.chargingProfilesSender.OnPostActiveChargingProfile(r.Context(), sessionID, body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PostOcpiChargingProfile(w http.ResponseWriter, r *http.Request) {
	var body ocpi.RawMessage[ChargingProfileResult]
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	sessionID := r.PathValue("session_id")
	if err := s.chargingProfilesSender.OnPostChargingProfile(r.Context(), sessionID, body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PostOcpiClearProfile(w http.ResponseWriter, r *http.Request) {
	var body ocpi.RawMessage[ClearProfileResult]
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	sessionID := r.PathValue("session_id")
	if err := s.chargingProfilesSender.OnPostClearProfile(r.Context(), sessionID, body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PutOcpiActiveChargingProfile(w http.ResponseWriter, r *http.Request) {
	var body ocpi.RawMessage[ActiveChargingProfile]
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	sessionID := r.PathValue("session_id")
	if err := s.chargingProfilesSender.OnPutActiveChargingProfile(r.Context(), sessionID, body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
