package ocpi221

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiCDRs(w http.ResponseWriter, r *http.Request) {
	params := GetCdrsParams{}
	response, err := s.cdrsSender.OnGetCDRs(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiCDR(w http.ResponseWriter, r *http.Request) {
	cdrID := r.PathValue("cdr_id")
	cdr, err := s.cdrsReceiver.OnGetCDR(r.Context(), cdrID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(cdr))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCDR(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	resp, err := s.cdrsReceiver.OnPostCDR(r.Context(), ocpi.RawMessage[ChargeDetailRecord](body))
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
