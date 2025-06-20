package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiTokens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := GetTokensParams{}
	response, err := s.emsp.OnGetTokens(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) PostOcpiToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenUID := chi.URLParam(r, "token_uid")

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	authInfo, err := s.emsp.OnPostToken(
		r.Context(),
		tokenUID,
		tokenType,
		ocpi.RawMessage[*LocationReferences](body),
	)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(authInfo))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) GetOcpiToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	countryCode := chi.URLParam(r, "country_code")
	partyID := chi.URLParam(r, "party_id")
	tokenUID := chi.URLParam(r, "token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	token, err := s.cpo.OnGetClientOwnedToken(
		r.Context(),
		countryCode,
		partyID,
		tokenUID,
		tokenType,
	)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(token))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PutOcpiToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	countryCode := chi.URLParam(r, "country_code")
	partyID := chi.URLParam(r, "party_id")
	tokenUID := chi.URLParam(r, "token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	if err := s.cpo.OnPutClientOwnedToken(
		r.Context(),
		countryCode,
		partyID,
		tokenUID,
		ocpi.RawMessage[Token](body),
		tokenType,
	); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PatchOcpiToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	countryCode := chi.URLParam(r, "country_code")
	partyID := chi.URLParam(r, "party_id")
	tokenUID := chi.URLParam(r, "token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	if err := s.cpo.OnPatchClientOwnedToken(
		r.Context(),
		countryCode,
		partyID,
		tokenUID,
		ocpi.RawMessage[PatchedToken](body),
		tokenType,
	); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
