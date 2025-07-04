package ocpi221

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) authorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := strings.TrimSpace(r.Header.Get("Authorization"))
		if !s.credentials.VerifyToken(r.Context(), token) {
			httputil.ResponseError(w, fmt.Errorf(`invalid token`), ocpi.StatusCodeClientError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
