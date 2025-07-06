package ocpi221

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx := GetRequestContext(ctx)
	credential, err := s.ocpi.OnGetCredential(ctx, requestCtx.Token())
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, credential)
}

func (s *Server) PostOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	ctx := r.Context()
	requestCtx := GetRequestContext(ctx)
	token := requestCtx.Token()
	if s.ocpi.IsClientRegistered(ctx, token) {
		// https://github.com/ocpi/ocpi/blob/release-2.2.1-bugfixes/credentials.asciidoc
		// Refer to sectio 1.2.2
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	credential, err := s.ocpi.OnPostCredential(ctx, token, ocpi.RawMessage[Credential](body))
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ocpihttp.Response(w, credential)
}

func (s *Server) PutOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r)
		return
	}

	ctx := r.Context()
	requestCtx := GetRequestContext(ctx)
	token := requestCtx.Token()
	if !s.ocpi.IsClientRegistered(ctx, token) {
		// https://github.com/ocpi/ocpi/blob/release-2.2.1-bugfixes/credentials.asciidoc
		// Refer to sectio 1.2.3
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	credential, err := s.ocpi.OnPutCredential(ctx, token, ocpi.RawMessage[Credential](body))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, credential)
}

func (s *Server) DeleteOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx := GetRequestContext(ctx)
	token := requestCtx.Token()
	if !s.ocpi.IsClientRegistered(ctx, token) {
		// https://github.com/ocpi/ocpi/blob/release-2.2.1-bugfixes/credentials.asciidoc
		// Refer to sectio 1.2.4
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := s.ocpi.OnDeleteCredential(r.Context(), token)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
