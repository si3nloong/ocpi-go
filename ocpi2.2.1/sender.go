package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Sender interface {
	GetLocations(ctx context.Context, params GetOcpiLocationsParams) ([]Location, error)
	GetTariffs(ctx context.Context, params GetOcpiTariffsParams) ([]Tariff, error)
	GetSessions(ctx context.Context, params GetOcpiSessionsParams) ([]Session, error)
	GetTokens(ctx context.Context, params GetOcpiTokensParams) ([]Token, error)
	PostToken(ctx context.Context, tokenUid string, tokenType ...TokenType) (*AuthorizationInfo, error)
}

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetOcpiLocationsParams{}
	locations, err := s.sender.GetLocations(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(locations))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetOcpiLocationsParams{}
	locations, err := s.sender.GetLocations(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(locations))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiTariffs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetOcpiTariffsParams{}
	tariffs, err := s.sender.GetTariffs(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(tariffs))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetOcpiSessionsParams{}
	sessions, err := s.sender.GetSessions(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(sessions))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiTokens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetOcpiTokensParams{}
	tokens, err := s.sender.GetTokens(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(tokens))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PostOcpiToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenUid := chi.URLParam(r, "token_uid")

	token, err := s.sender.PostToken(r.Context(), tokenUid)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(token))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.GenericServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
