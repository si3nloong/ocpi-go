package ocpi211

import (
	"net/http"
)

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (s *Server) PostOcpiCredentials(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) PutOcpiCredentials(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) DeleteOcpiCredentials(w http.ResponseWriter, r *http.Request) {
}
