package ocpi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Response[T any] struct {
	RawData       json.RawMessage `json:"data,omitempty"`
	StatusCode    StatusCode      `json:"status_code"`
	StatusMessage string          `json:"status_message,omitempty,omitzero"`
	Timestamp     time.Time       `json:"timestamp"`
}

func (r *Response[T]) Decode(dest any) error {
	if err := json.Unmarshal(r.RawData, dest); err != nil {
		return err
	}
	return nil
}

func (r *Response[T]) Data() (T, error) {
	var o T
	if r.StatusCode >= 1_000 && r.StatusCode < 2_000 {
		if len(r.RawData) == 0 {
			return o, nil
		}
		if err := json.Unmarshal(r.RawData, &o); err != nil {
			return o, err
		}
		return o, nil
	}
	return o, NewOCPIError(r.StatusCode, r.StatusMessage)
}

func NewResponse[T any](value T) *Response[T] {
	b, _ := json.Marshal(value)
	return &Response[T]{
		RawData:       b,
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     time.Now().UTC(),
	}
}

func NewEmptyResponse() *Response[any] {
	return &Response[any]{
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     time.Now().UTC(),
	}
}

type PaginatedResponse[T any] struct {
	headers       http.Header
	RawData       json.RawMessage `json:"data,omitempty"`
	StatusCode    StatusCode      `json:"status_code"`
	StatusMessage string          `json:"status_message,omitempty,omitzero"`
	Timestamp     time.Time       `json:"timestamp"`
}

func (r *PaginatedResponse[T]) Data() ([]T, error) {
	var o []T
	if r.StatusCode >= 1_000 && r.StatusCode < 2_000 {
		if len(r.RawData) == 0 {
			return o, nil
		}
		if err := json.Unmarshal(r.RawData, &o); err != nil {
			return o, err
		}
		return o, nil
	}
	return o, NewOCPIError(r.StatusCode, r.StatusMessage)
}

func (r *PaginatedResponse[T]) ScanHeaders(headers http.Header) error {
	r.headers = headers.Clone()
	return nil
}

func (r *PaginatedResponse[T]) Link() string {
	return r.headers.Get(HttpHeaderLink)
}

func (r *PaginatedResponse[T]) TotalCount() (int, error) {
	return strconv.Atoi(r.headers.Get(HttpHeaderXTotalCount))
}

func (r *PaginatedResponse[T]) Limit() (int, error) {
	return strconv.Atoi(r.headers.Get(HttpHeaderXLimit))
}
