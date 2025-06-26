package ocpi

import (
	"time"
)

type Response[T any] struct {
	Data          T          `json:"data,omitempty"`
	StatusCode    StatusCode `json:"status_code"`
	StatusMessage string     `json:"status_message,omitempty,omitzero"`
	Timestamp     time.Time  `json:"timestamp"`
}

func NewResponse[T any](value T) *Response[T] {
	return &Response[T]{
		Data:          value,
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     time.Now().UTC(),
	}
}

func NewEmptyResponse() *Response[any] {
	return &Response[any]{
		StatusCode: StatusCodeSuccess,
		Timestamp:  time.Now().UTC(),
	}
}

type PaginationResponse[T any] struct {
	Header struct {
		TotalCount int64
		Limit      int64
	}
	Response[[]T]
}
