package ocpi221

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiCDRs(w http.ResponseWriter, r *http.Request) {
	params := GetCDRsParams{}
	response, err := s.cdrsSender.OnGetCDRs(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, DateTime{Time: time.Now().UTC()}, response)
}

func (s *Server) GetOcpiCDR(w http.ResponseWriter, r *http.Request) {
	cdrID := r.PathValue("cdr_id")
	cdr, err := s.cdrsReceiver.OnGetCDR(r.Context(), cdrID)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, cdr)
}

func (s *Server) PostOcpiCDR(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	resp, err := s.cdrsReceiver.OnPostCDR(r.Context(), ocpi.RawMessage[CDR](body))
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse(DateTime{Time: time.Now().UTC()}))
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	w.Header().Set("Location", resp.Location)
	w.Write(b)
	w.WriteHeader(http.StatusCreated)
}
