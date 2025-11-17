package ocpi230

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiClientInfos(w http.ResponseWriter, r *http.Request) {
	params := GetHubClientInfoParams{}
	response, err := s.hubClientInfoSender.OnGetHubClientInfos(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, DateTime{Time: time.Now().UTC()}, response)
}

func (s *Server) GetOcpiClientInfo(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")

	clientInfo, err := s.hubClientInfoReceiver.OnGetHubClientInfo(r.Context(), countryCode, partyID)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, clientInfo)
}

func (s *Server) PutOcpiClientInfo(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")

	if err := s.hubClientInfoReceiver.OnPutHubClientInfo(r.Context(), countryCode, partyID, ocpi.RawMessage[ClientInfo](body)); err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.EmptyResponse(w, DateTime{Time: time.Now().UTC()})
}
