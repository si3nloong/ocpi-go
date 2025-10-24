package ocpi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type PaginatedResponse[T any] struct {
	headers       http.Header
	RawData       json.RawMessage `json:"data,omitempty"`
	StatusCode    StatusCode      `json:"status_code"`
	StatusMessage string          `json:"status_message,omitempty,omitzero"`
	Timestamp     time.Time       `json:"timestamp"`
}

func (r *PaginatedResponse[T]) Data() ([]T, error) {
	if r.StatusCode >= StatusCodeSuccess && r.StatusCode < 2_000 {
		var o []T
		if len(r.RawData) == 0 {
			return o, nil
		}
		if err := json.Unmarshal(r.RawData, &o); err != nil {
			return o, err
		}
		return o, nil
	}
	return nil, NewOCPIError(r.StatusCode, r.StatusMessage)
}

func (r *PaginatedResponse[T]) StrictData() ([]T, error) {
	if r.StatusCode >= StatusCodeSuccess && r.StatusCode < 2_000 {
		var o []T
		if len(r.RawData) == 0 {
			return o, nil
		}
		if err := json.Unmarshal(r.RawData, &o); err != nil {
			return nil, err
		}
		var s = struct {
			v []T `validate:"omitempty,dive,required"`
		}{o}
		if err := validate.Struct(s); err != nil {
			return nil, err
		}
		return o, nil
	}
	return nil, NewOCPIError(r.StatusCode, r.StatusMessage)
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

func NewPaginatedResponse[T any](link string, limit, totalCount int, v []T) *PaginatedResponse[T] {
	b, _ := json.Marshal(v)
	headers := make(http.Header)
	headers.Add(HttpHeaderXLimit, strconv.Itoa(limit))
	headers.Add(HttpHeaderLink, link)
	headers.Add(HttpHeaderXTotalCount, strconv.Itoa(totalCount))
	return &PaginatedResponse[T]{
		headers:       headers,
		RawData:       b,
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     time.Now().UTC(),
	}
}
