package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiCDRs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetCdrsParams{}
	response, err := s.cpo.OnGetCDRs(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiCDR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cdrID := r.PathValue("cdr_id")

	cdr, err := s.emsp.OnGetCDR(r.Context(), cdrID)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(cdr))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCDR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	resp, err := s.emsp.OnPostCDR(
		r.Context(),
		ocpi.RawMessage[ChargeDetailRecord](body),
	)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Location", resp.Location.String())
	w.Write(b)
}
