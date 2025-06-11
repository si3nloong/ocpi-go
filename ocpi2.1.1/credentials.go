package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpoints := []Endpoint{
		{Identifier: ModuleIDCredentials, URL: httputil.GetHostname(r) + s.baseUrl + "/credentials"},
		{Identifier: ModuleIDLocations, URL: httputil.GetHostname(r) + s.baseUrl + "/locations"},
		{Identifier: ModuleIDSessions, URL: httputil.GetHostname(r) + s.baseUrl + "/sessions"},
		{Identifier: ModuleIDTokens, URL: httputil.GetHostname(r) + s.baseUrl + "/tokens"},
		{Identifier: ModuleIDTariffs, URL: httputil.GetHostname(r) + s.baseUrl + "/tariffs"},
		{Identifier: ModuleIDCdrs, URL: httputil.GetHostname(r) + s.baseUrl + "/cdrs"},
	}

	b, err := json.Marshal(ocpi.NewResponse(endpoints))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
