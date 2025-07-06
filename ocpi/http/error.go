package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func ResponseError(w http.ResponseWriter, err error) {
	var httpErr ocpi.HTTPError
	if errors.As(err, &httpErr) {
		w.WriteHeader(http.StatusOK)
		return
	}
	b, _ := json.Marshal(ocpi.Response[any]{
		// StatusCode:    statusCode,
		StatusMessage: err.Error(),
		Timestamp:     time.Now(),
	})
	h := w.Header()
	h.Del("Content-Length")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
