package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiCDRs(w http.ResponseWriter, r *http.Request) {
	params := GetCdrsParams{}
	response, err := s.cpo.OnGetCDRs(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiCDR(w http.ResponseWriter, r *http.Request) {
	cdrID := r.PathValue("cdr_id")

	cdr, err := s.emsp.OnGetCDR(r.Context(), cdrID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, cdr)
}

func (s *Server) PostOcpiCDR(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	resp, err := s.emsp.OnPostCDR(r.Context(), ocpi.RawMessage[ChargeDetailRecord](body))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Location", resp.Location)
	w.Write(b)
}
