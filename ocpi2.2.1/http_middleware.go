package ocpi221

import (
	"net/http"
	"strings"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) authorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")

		requestID := strings.TrimSpace(r.Header.Get(ocpi.HttpHeaderXRequestID))
		correlationID := strings.TrimSpace(r.Header.Get(ocpi.HttpHeaderXCorrelationID))
		defer func() {
			w.Header().Set(ocpi.HttpHeaderXRequestID, requestID)
			w.Header().Set(ocpi.HttpHeaderXCorrelationID, correlationID)
		}()

		token := strings.TrimSpace(r.Header.Get("Authorization"))
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := s.tokenResolver(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !s.credentials.VerifyCredentialsToken(r.Context(), token) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(WithRequestContext(
			r.Context(),
			&RequestContext{
				token:         token,
				requestID:     requestID,
				requestURI:    r.RequestURI,
				correlationID: correlationID,
			}),
		))
	})
}
