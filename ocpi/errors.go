package ocpi

import "fmt"

var (
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
	return fmt.Sprintf("ocpi: error due to %d", e.StatusCode)
}

type HTTPError struct {
	StatusCode int
}
