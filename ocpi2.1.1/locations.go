package ocpi211

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

// GetOcpiLocations handles the /locations endpoint.
func (s *Server) GetOcpiLocations(w http.ResponseWriter, r *http.Request) {

	params := GetLocationsParams{}
	queryString := r.URL.Query()
	if queryString.Has("date_from") {
		dt, err := ParseDateTime(queryString.Get("date_from"))
		if err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
		params.DateFrom = &dt
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
	response, err := s.cpo.OnGetLocations(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	locationID := r.PathValue("location_id")
	evseUID := strings.TrimSpace(r.PathValue("evse_uid"))
	connectorID := strings.TrimSpace(r.PathValue("connector_id"))

	var (
		resp any
		err  error
	)
	if evseUID != "" && connectorID != "" {
		resp, err = s.emsp.OnGetClientOwnedLocationConnector(ctx, countryCode, partyID, locationID, evseUID, connectorID)
	} else if evseUID != "" {
		resp, err = s.emsp.OnGetClientOwnedLocationEVSE(ctx, countryCode, partyID, locationID, evseUID)
	} else {
		resp, err = s.emsp.OnGetClientOwnedLocation(ctx, countryCode, partyID, locationID)
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

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	locationID := r.PathValue("location_id")
	evseUID := strings.TrimSpace(r.PathValue("evse_uid"))
	connectorID := strings.TrimSpace(r.PathValue("connector_id"))

	s.logger.DebugContext(ctx, "PutOcpiLocation",
		"country_code", countryCode,
		"party_id", partyID,
		"location_id", locationID,
		"evse_uid", evseUID,
		"connector_id", connectorID)

	if evseUID != "" && connectorID != "" {
		if err := s.emsp.OnPutClientOwnedLocationConnector(
			ctx,
			countryCode,
			partyID,
			locationID,
			evseUID,
			connectorID,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else if evseUID != "" {
		if err := s.emsp.OnPutClientOwnedLocationEVSE(
			ctx,
			countryCode,
			partyID,
			locationID,
			evseUID,
			ocpi.RawMessage[Location](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else {
		if err := s.emsp.OnPutClientOwnedLocation(
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

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	locationID := r.PathValue("location_id")
	evseUID := strings.TrimSpace(r.PathValue("evse_uid"))
	connectorID := strings.TrimSpace(r.PathValue("connector_id"))

	if evseUID != "" && connectorID != "" {
		if err := s.emsp.OnPatchClientOwnedLocationConnector(
			ctx,
			countryCode,
			partyID,
			locationID,
			evseUID,
			connectorID,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else if evseUID != "" {
		if err := s.emsp.OnPatchClientOwnedLocationEVSE(
			ctx,
			countryCode,
			partyID,
			locationID,
			evseUID,
			ocpi.RawMessage[PatchedLocation](body),
		); err != nil {
			httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
			return
		}
	} else {
		if err := s.emsp.OnPatchClientOwnedLocation(
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

func (c *ClientConn) GetLocation(
	ctx context.Context,
	locationID string,
) (*LocationResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations)
	if err != nil {
		return nil, err
	}

	var o LocationResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+locationID, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
