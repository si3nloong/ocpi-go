package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiVersionDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpoints := []Endpoint{}

	b, err := json.Marshal(ocpi.NewResponse(endpoints))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (s *Server) PostOcpiCredentials(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) PutOcpiCredentials(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) DeleteOcpiCredentials(w http.ResponseWriter, r *http.Request) {
}
