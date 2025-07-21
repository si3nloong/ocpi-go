package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiTariffs(w http.ResponseWriter, r *http.Request) {
	params := GetTariffsParams{}
	response, err := s.cpo.OnGetTariffs(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiTariff(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	tariff, err := s.emsp.OnGetClientOwnedTariff(r.Context(), countryCode, partyID, tariffID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, tariff)
}

func (s *Server) PutOcpiTariff(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.emsp.OnPutClientOwnedTariff(r.Context(), countryCode, partyID, tariffID, ocpi.RawMessage[Tariff](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PatchOcpiTariff(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.emsp.OnPatchClientOwnedTariff(r.Context(), countryCode, partyID, tariffID, ocpi.RawMessage[PartialTariff](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) DeleteOcpiTariff(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.emsp.OnDeleteClientOwnedTariff(r.Context(), countryCode, partyID, tariffID); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
