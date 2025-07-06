package ocpi221

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiLocations(w http.ResponseWriter, r *http.Request) {
	params := GetLocationsParams{}
	response, err := s.locationsSender.OnGetLocations(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	locationID := r.PathValue("location_id")
	evseUID := strings.TrimSpace(r.PathValue("evse_uid"))
	connectorID := strings.TrimSpace(r.PathValue("connector_id"))

	var (
		resp any
		err  error
	)
	if evseUID != "" && connectorID != "" {
		resp, err = s.locationsSender.OnGetLocationConnector(ctx, locationID, evseUID, connectorID)
	} else if evseUID != "" {
		resp, err = s.locationsSender.OnGetLocationEVSE(ctx, locationID, evseUID)
	} else {
		resp, err = s.locationsSender.OnGetLocation(ctx, locationID)
	}
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, resp)
}

func (s *Server) GetOcpiClientOwnedLocation(w http.ResponseWriter, r *http.Request) {
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
		resp, err = s.locationsReceiver.OnGetClientOwnedLocationConnector(ctx, countryCode, partyID, locationID, evseUID, connectorID)
	} else if evseUID != "" {
		resp, err = s.locationsReceiver.OnGetClientOwnedLocationEVSE(ctx, countryCode, partyID, locationID, evseUID)
	} else {
		resp, err = s.locationsReceiver.OnGetClientOwnedLocation(ctx, countryCode, partyID, locationID)
	}
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, resp)
}

func (s *Server) PutOcpiLocation(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	locationID := r.PathValue("location_id")
	evseUID := strings.TrimSpace(r.PathValue("evse_uid"))
	connectorID := strings.TrimSpace(r.PathValue("connector_id"))

	if evseUID != "" && connectorID != "" {
		if err := s.locationsReceiver.OnPutClientOwnedLocationConnector(ctx, countryCode, partyID, locationID, evseUID, connectorID, ocpi.RawMessage[Location](body)); err != nil {
			ocpihttp.Response(w, err)
			return
		}
	} else if evseUID != "" {
		if err := s.locationsReceiver.OnPutClientOwnedLocationEVSE(ctx, countryCode, partyID, locationID, evseUID, ocpi.RawMessage[Location](body)); err != nil {
			ocpihttp.Response(w, err)
			return
		}
	} else {
		if err := s.locationsReceiver.OnPutClientOwnedLocation(ctx, countryCode, partyID, locationID, ocpi.RawMessage[Location](body)); err != nil {
			ocpihttp.Response(w, err)
			return
		}
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PatchOcpiLocation(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	ctx := r.Context()
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	locationID := r.PathValue("location_id")
	evseUID := strings.TrimSpace(r.PathValue("evse_uid"))
	connectorID := strings.TrimSpace(r.PathValue("connector_id"))

	if evseUID != "" && connectorID != "" {
		if err := s.locationsReceiver.OnPatchClientOwnedLocationConnector(ctx, countryCode, partyID, locationID, evseUID, connectorID, ocpi.RawMessage[PartialLocation](body)); err != nil {
			ocpihttp.Response(w, err)
			return
		}
	} else if evseUID != "" {
		if err := s.locationsReceiver.OnPatchClientOwnedLocationEVSE(ctx, countryCode, partyID, locationID, evseUID, ocpi.RawMessage[PartialLocation](body)); err != nil {
			ocpihttp.Response(w, err)
			return
		}
	} else {
		if err := s.locationsReceiver.OnPatchClientOwnedLocation(ctx, countryCode, partyID, locationID, ocpi.RawMessage[PartialLocation](body)); err != nil {
			ocpihttp.Response(w, err)
			return
		}
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
