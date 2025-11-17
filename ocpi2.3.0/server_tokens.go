package ocpi230

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiTokens(w http.ResponseWriter, r *http.Request) {
	params := GetTokensParams{}
	response, err := s.tokensSender.OnGetTokens(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, DateTime{Time: time.Now().UTC()}, response)
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
		case string(TokenTypeAdHocUser):
			tokenType = TokenTypeAdHocUser
		case string(TokenTypeAppUser):
			tokenType = TokenTypeAppUser
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	token, err := s.tokensReceiver.OnGetClientOwnedToken(
		r.Context(),
		countryCode,
		partyID,
		tokenUID,
		tokenType,
	)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, token)
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
		case string(TokenTypeAdHocUser):
			tokenType = TokenTypeAdHocUser
		case string(TokenTypeAppUser):
			tokenType = TokenTypeAppUser
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	authInfo, err := s.tokensSender.OnPostToken(
		r.Context(),
		tokenUID,
		ocpi.RawMessage[LocationReferences](body),
		tokenType,
	)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, authInfo)
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
		case string(TokenTypeAdHocUser):
			tokenType = TokenTypeAdHocUser
		case string(TokenTypeAppUser):
			tokenType = TokenTypeAppUser
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	if err := s.tokensReceiver.OnPutClientOwnedToken(
		r.Context(),
		countryCode,
		partyID,
		tokenUID,
		ocpi.RawMessage[Token](body),
		tokenType,
	); err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.EmptyResponse(w, DateTime{Time: time.Now().UTC()})
}

func (s *Server) PatchOcpiToken(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tokenUID := r.PathValue("token_uid")

	tokenType := TokenTypeRFID
	if r.URL.Query().Has("type") {
		switch r.URL.Query().Get("type") {
		case string(TokenTypeRFID):
			tokenType = TokenTypeRFID
		case string(TokenTypeAdHocUser):
			tokenType = TokenTypeAdHocUser
		case string(TokenTypeAppUser):
			tokenType = TokenTypeAppUser
		case string(TokenTypeOther):
			tokenType = TokenTypeOther
		}
	}

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	if err := s.tokensReceiver.OnPatchClientOwnedToken(
		r.Context(),
		countryCode,
		partyID,
		tokenUID,
		ocpi.RawMessage[PartialToken](body),
		tokenType,
	); err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.EmptyResponse(w, DateTime{Time: time.Now().UTC()})
}
