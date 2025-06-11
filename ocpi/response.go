package ocpi

import (
	"time"
)

type Response[T any] struct {
	Data          T          `json:"data,omitempty"`
	StatusCode    StatusCode `json:"status_code"`
	StatusMessage string     `json:"status_message"`
	Timestamp     time.Time  `json:"timestamp"`
}

func NewResponse[T any](value T) *Response[T] {
	return &Response[T]{
		Data:          value,
		StatusCode:    GenericSuccessCode,
		StatusMessage: GenericSuccessCode.String(),
		Timestamp:     time.Now().UTC(),
	}
}
