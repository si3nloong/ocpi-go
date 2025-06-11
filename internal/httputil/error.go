package httputil

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func ResponseError(w http.ResponseWriter, err error, statusCode ocpi.StatusCode) {
	b, _ := json.Marshal(ocpi.Response[any]{
		StatusCode:    statusCode,
		StatusMessage: err.Error(),
		Timestamp:     time.Now(),
	})
	h := w.Header()
	h.Del("Content-Length")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
