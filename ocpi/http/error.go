package http

import (
	"net/http"
)

func BadRequest(w http.ResponseWriter, r *http.Request, errMsg string) {
	h := w.Header()
	h.Del("Content-Length")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(errMsg))
}

// func ResponseError(w http.ResponseWriter, err error) {
// 	httpErr := &ocpi.HTTPError{}
// 	if errors.As(err, &httpErr) {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}
// 	b, _ := json.Marshal(ocpi.Response[any]{
// 		// StatusCode:    statusCode,
// 		StatusMessage: err.Error(),
// 		Timestamp:     time.Now(),
// 	})
// 	h := w.Header()
// 	h.Del("Content-Length")
// 	h.Set("X-Content-Type-Options", "nosniff")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(b)
// }
