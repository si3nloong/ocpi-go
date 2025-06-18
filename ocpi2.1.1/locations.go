package ocpi211

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

// GetOcpiLocations handles the /locations endpoint.
func (s *Server) GetOcpiLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetLocationsParams{}
	queryString := r.URL.Query()
	if queryString.Has("date_from") {
		dt, err := ocpi.ParseDateTime(queryString.Get("date_from"))
		if err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
		params.DateFrom = &dt
	}
	if queryString.Has("date_to") {
		dt, err := ocpi.ParseDateTime(queryString.Get("date_to"))
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
	response, err := s.cpo.GetLocations(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	countryCode := chi.URLParam(r, "country_code")
	partyID := chi.URLParam(r, "party_id")
	locationID := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorID := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	var (
		resp any
		err  error
	)
	if evseUid != "" && connectorID != "" {
		resp, err = s.emsp.GetLocationConnector(ctx, countryCode, partyID, locationID, evseUid, connectorID)
	} else if evseUid != "" {
		resp, err = s.emsp.GetLocationEVSE(ctx, countryCode, partyID, locationID, evseUid)
	} else {
		resp, err = s.emsp.GetLocation(ctx, countryCode, partyID, locationID)
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
	partyID := chi.URLParam(r, "party_id")
	locationID := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorID := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	s.logger.DebugContext(ctx, "PutOcpiLocation",
		"country_code", countryCode,
		"party_id", partyID,
		"location_id", locationID,
		"evse_uid", evseUid,
		"connector_id", connectorID)

	if evseUid != "" && connectorID != "" {
		if err := s.emsp.PutLocationConnector(
			ctx,
			countryCode,
			partyID,
			locationID,
			evseUid,
			connectorID,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else if evseUid != "" {
		if err := s.emsp.PutLocationEVSE(
			ctx,
			countryCode,
			partyID,
			locationID,
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
			partyID,
			locationID,
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
	partyID := chi.URLParam(r, "party_id")
	locationID := chi.URLParam(r, "location_id")
	evseUid := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorID := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	if evseUid != "" && connectorID != "" {
		if err := s.emsp.PatchLocationConnector(
			ctx,
			countryCode,
			partyID,
			locationID,
			evseUid,
			connectorID,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else if evseUid != "" {
		if err := s.emsp.PatchLocationEVSE(
			ctx,
			countryCode,
			partyID,
			locationID,
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
			partyID,
			locationID,
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
