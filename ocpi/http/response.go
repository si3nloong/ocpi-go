package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func Response[T any](w http.ResponseWriter, value T) {
	switch vi := any(value).(type) {
	case error:
		var httpErr ocpi.HTTPError
		if errors.As(vi, &httpErr) {
			w.WriteHeader(httpErr.StatusCode)
			return
		}

		var ocpiErr ocpi.OCPIError
		if errors.As(vi, &ocpiErr) {
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(ocpi.Response[any]{
				StatusCode:    ocpiErr.StatusCode,
				StatusMessage: vi.Error(),
				Timestamp:     time.Now(),
			})
			w.Write(b)
			return
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(ocpi.Response[any]{
			StatusCode:    ocpi.StatusCodeServerError,
			StatusMessage: vi.Error(),
			Timestamp:     time.Now(),
		})
		w.Write(b)

	case *ocpi.Response[T]:
		b, _ := json.Marshal(vi)
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	default:
		b, _ := json.Marshal(ocpi.NewResponse(value))
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
