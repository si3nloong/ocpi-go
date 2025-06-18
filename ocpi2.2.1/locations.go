package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetLocationsParams{}
	response, err := s.locationsSender.GetLocations(r.Context(), params)
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
	evseUID := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorID := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	var (
		resp any
		err  error
	)
	if evseUID != "" && connectorID != "" {
		resp, err = s.locationsReceiver.GetLocationConnector(ctx, countryCode, partyID, locationID, evseUID, connectorID)
	} else if evseUID != "" {
		resp, err = s.locationsReceiver.GetLocationEVSE(ctx, countryCode, partyID, locationID, evseUID)
	} else {
		resp, err = s.locationsReceiver.GetLocation(ctx, countryCode, partyID, locationID)
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
	evseUID := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorID := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	s.logger.DebugContext(ctx, "PutOcpiLocation",
		"country_code", countryCode,
		"party_id", partyID,
		"location_id", locationID,
		"evse_uid", evseUID,
		"connector_id", connectorID)

	if evseUID != "" && connectorID != "" {
		if err := s.locationsReceiver.PutLocationConnector(
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
		if err := s.locationsReceiver.PutLocationEVSE(
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
		if err := s.locationsReceiver.PutLocation(
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
	evseUID := strings.TrimSpace(chi.URLParam(r, "evse_uid"))
	connectorID := strings.TrimSpace(chi.URLParam(r, "connector_id"))

	if evseUID != "" && connectorID != "" {
		if err := s.locationsReceiver.PatchLocationConnector(
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
		if err := s.locationsReceiver.PatchLocationEVSE(
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
		if err := s.locationsReceiver.PatchLocation(
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

func (c *Client) GetLocations(
	ctx context.Context,
	params ...GetLocationsParams,
) (ocpi.Result[[]Location], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("limit", "100")
	if len(params) > 0 {
		p := params[0]
		if p.DateFrom != nil && p.DateFrom.IsZero() {
			query.Set("date_from", p.DateFrom.Format(time.RFC3339))
		}
		if p.DateTo != nil && p.DateTo.IsZero() {
			query.Set("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Set("offset", strconv.FormatUint(uint64(*p.Offset), 10))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Set("limit", strconv.FormatUint(uint64(*p.Limit), 10))
		}
	}
	u.RawQuery = query.Encode()

	var o LocationsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return ocpi.NewResult(o), nil
}

func (c *Client) GetLocation(
	ctx context.Context,
	locationID string,
) (*LocationResponse, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDLocations, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var o LocationResponse
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+locationID, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
