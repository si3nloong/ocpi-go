package ocpi221

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiTariffs(w http.ResponseWriter, r *http.Request) {
	params := GetTariffsParams{}
	response, err := s.tariffsSender.OnGetTariffs(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiTariff(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	tariff, err := s.tariffsReceiver.OnGetClientOwnedTariff(r.Context(), countryCode, partyID, tariffID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewResponse(tariff))
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

	if err := s.tariffsReceiver.OnPutClientOwnedTariff(r.Context(), countryCode, partyID, tariffID, ocpi.RawMessage[Tariff](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) DeleteOcpiTariff(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.tariffsReceiver.OnDeleteClientOwnedTariff(r.Context(), countryCode, partyID, tariffID); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
