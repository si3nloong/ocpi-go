package ocpi211

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiTokens(w http.ResponseWriter, r *http.Request) {
	params := GetTokensParams{}
	response, err := s.emsp.OnGetTokens(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
}

func (s *Server) PostOcpiToken(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	tokenUID := r.PathValue("token_uid")
	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	authInfo, err := s.emsp.OnPostToken(r.Context(), tokenUID, tokenType, ocpi.RawMessage[*LocationReferences](body))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	// When the eMSP receives a 'real-time' authorization request from a CPO that contains too little information (no LocationReferences provided) to determine if the Token might be used, the eMSP SHOULD respond with the OCPI status: 2002
	ocpihttp.Response(w, authInfo)
}

func (s *Server) GetOcpiToken(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tokenUID := r.PathValue("token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	token, err := s.cpo.OnGetClientOwnedToken(r.Context(), countryCode, partyID, tokenUID, tokenType)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, token)
}

func (s *Server) PutOcpiToken(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tokenUID := r.PathValue("token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	if err := s.cpo.OnPutClientOwnedToken(r.Context(), countryCode, partyID, tokenUID, ocpi.RawMessage[Token](body), tokenType); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PatchOcpiToken(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tokenUID := r.PathValue("token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	if err := s.cpo.OnPatchClientOwnedToken(r.Context(), countryCode, partyID, tokenUID, ocpi.RawMessage[PartialToken](body), tokenType); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
