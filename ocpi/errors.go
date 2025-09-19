package ocpi

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNoMutualVersion                  = fmt.Errorf(`ocpi: no mutual version`)
	ErrNotImplemented                   = errors.New("ocpi: not implemented")
	ErrClient                           = &OCPIError{StatusCode: StatusCodeClientError}
	ErrClientInvalidOrMissingParameters = &OCPIError{StatusCode: StatusCodeClientErrorInvalidOrMissingParameters}
)

type OCPIError struct {
	StatusCode StatusCode
	message    *string
}

func NewOCPIError(statusCode StatusCode, message ...string) *OCPIError {
	err := &OCPIError{
		StatusCode: statusCode,
	}
	if len(message) > 0 {
		err.message = &message[0]
	}
	return err
}

func (e *OCPIError) Error() string {
	if e.StatusCode >= 2_000 && e.StatusCode < 3_000 {
		return fmt.Sprintf("ocpi: %d client error", e.StatusCode)
	}
	if e.StatusCode >= 3_000 && e.StatusCode < 4_000 {
		return fmt.Sprintf("ocpi: %d server error", e.StatusCode)
	}
	if e.StatusCode >= 4_000 && e.StatusCode < 5_000 {
		return fmt.Sprintf("ocpi: %d hub error", e.StatusCode)
	}
	return fmt.Sprintf("ocpi: error due to %d", e.StatusCode)
}

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	switch e.StatusCode {
	case http.StatusBadRequest:
		return "bad request"
	case http.StatusMethodNotAllowed:
		return "method not allowed"
	}
	return fmt.Sprintf("http error: %d", e.StatusCode)
}
