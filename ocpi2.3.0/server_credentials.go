package ocpi230

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx := ocpi.GetRequestContext(ctx)
	credential, err := s.ocpi.OnGetCredential(ctx, requestCtx.Token)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, credential)
}

func (s *Server) PostOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	var body ocpi.RawMessage[Credentials]
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	ctx := r.Context()
	requestCtx := ocpi.GetRequestContext(ctx)
	tokenA := requestCtx.Token
	if s.ocpi.IsClientRegistered(ctx, tokenA) {
		// https://github.com/ocpi/ocpi/blob/release-2.2.1-bugfixes/credentials.asciidoc
		// Refer to sectio 1.2.2
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	credentialData, err := body.Data()
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	// Store token B
	if err := s.ocpi.StoreCredentialsTokenB(ctx, credentialData); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	clientConn := NewClientWithTokenA(credentialData.URL, credentialData.Token, nil)
	versionsResponse, err := clientConn.GetVersions(ctx)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}
	versions, err := versionsResponse.Data()
	if err != nil {
		ocpihttp.Response(w, ocpi.NewOCPIError(ocpi.StatusCodeServerErrorNoMatchingEndpoints, `no version available`))
		return
	}
	if _, ok := versions.MutualVersion(ocpi.VersionNumber230); !ok {
		ocpihttp.Response(w, ocpi.NewOCPIError(ocpi.StatusCodeServerErrorNoMatchingEndpoints, `no version available`))
		return
	}

	versionDetailsResponse, err := clientConn.GetVersionDetails(ctx)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}
	versionDetails, err := versionDetailsResponse.Data()
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	if err := s.ocpi.StoreVersionDetails(ctx, versionDetails); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	credentialsWithTokenC, err := s.ocpi.GenerateCredentialsTokenC(ctx, tokenA)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, credentialsWithTokenC)
}

func (s *Server) PutOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	ctx := r.Context()
	requestCtx := ocpi.GetRequestContext(ctx)
	token := requestCtx.Token
	if !s.ocpi.IsClientRegistered(ctx, token) {
		// https://github.com/ocpi/ocpi/blob/release-2.2.1-bugfixes/credentials.asciidoc
		// Refer to sectio 1.2.3
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// TODO: call sender side to get versions and version_details
	credential, err := s.ocpi.OnPutCredential(ctx, token, ocpi.RawMessage[Credentials](body))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, credential)
}

func (s *Server) DeleteOcpiCredentials(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx := ocpi.GetRequestContext(ctx)
	token := requestCtx.Token
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
