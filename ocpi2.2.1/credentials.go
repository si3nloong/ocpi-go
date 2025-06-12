package ocpi221

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	credential, err := s.credentials.GetCredential(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(credential))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	credential, err := s.credentials.PostCredential(r.Context(), r.Header.Get("Authorization"), ocpi.RawMessage[Credential](body))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(credential))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PutOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	credential, err := s.credentials.PutCredential(r.Context(), r.Header.Get("Authorization"), ocpi.RawMessage[Credential](body))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(credential))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) DeleteOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	credential, err := s.credentials.DeleteCredential(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(credential))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
