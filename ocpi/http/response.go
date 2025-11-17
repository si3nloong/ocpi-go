package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func EmptyResponse(w http.ResponseWriter, ts ocpi.Timestamp) {
	b, _ := json.Marshal(ocpi.NewEmptyResponse(ts))
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Response[T ocpi.Timestamp, D any](w http.ResponseWriter, ts T, data D) {
	switch vi := any(data).(type) {
	case error:
		httpErr := &ocpi.HTTPError{}
		if errors.As(vi, &httpErr) {
			w.WriteHeader(httpErr.StatusCode)
			b, _ := json.Marshal(ocpi.Response[T, any]{
				StatusCode:    ocpi.StatusCodeClientError,
				StatusMessage: httpErr.Message,
				Timestamp:     ts,
			})
			w.Write(b)
			return
		}

		ocpiErr := &ocpi.OCPIError{}
		if errors.As(vi, &ocpiErr) {
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(ocpi.Response[T, any]{
				StatusCode:    ocpiErr.StatusCode,
				StatusMessage: vi.Error(),
				Timestamp:     ts,
			})
			w.Write(b)
			return
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(ocpi.Response[T, any]{
			StatusCode:    ocpi.StatusCodeServerError,
			StatusMessage: vi.Error(),
			Timestamp:     ts,
		})
		w.Write(b)

	case *ocpi.Response[T, D]:
		b, _ := json.Marshal(vi)
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	// Empty response
	case *ocpi.Response[T, any]:
		b, _ := json.Marshal(vi)
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	default:
		b, _ := json.Marshal(ocpi.NewResponse(ts, data))
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
