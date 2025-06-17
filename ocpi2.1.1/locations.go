package ocpi211

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

// GetOcpiLocations handles the /locations endpoint.
func (s *Server) GetOcpiLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: Implement the logic for handling OCPI locations.
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("GetOcpiLocations not implemented"))
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
		resp, err = s.emsp.GetLocationConnector(ctx, countryCode, partyId, locationId, evseUid, connectorId)
	} else if evseUid != "" {
		resp, err = s.emsp.GetLocationEVSE(ctx, countryCode, partyId, locationId, evseUid)
	} else {
		resp, err = s.emsp.GetLocation(ctx, countryCode, partyId, locationId)
	}
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

func (s *Server) PutOcpiLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
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
		if err := s.emsp.PutLocationConnector(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			connectorId,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else if evseUid != "" {
		if err := s.emsp.PutLocationEVSE(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else {
		if err := s.emsp.PutLocation(
			ctx,
			countryCode,
			partyId,
			locationId,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PatchOcpiLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	ctx := r.Context()
	countryCode := chi.URLParam(r, "country_code")
	partyId := chi.URLParam(r, "party_id")
	locationId := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorId := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	if evseUid != "" && connectorId != "" {
		if err := s.emsp.PatchLocationConnector(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			connectorId,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else if evseUid != "" {
		if err := s.emsp.PatchLocationEVSE(
			ctx,
			countryCode,
			partyId,
			locationId,
			evseUid,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else {
		if err := s.emsp.PatchLocation(
			ctx,
			countryCode,
			partyId,
			locationId,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
