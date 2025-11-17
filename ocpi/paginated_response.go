package ocpi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PaginatedResponse[T Timestamp, D any] struct {
	headers       http.Header
	RawData       json.RawMessage `json:"data,omitempty"`
	StatusCode    StatusCode      `json:"status_code"`
	StatusMessage string          `json:"status_message,omitempty,omitzero"`
	Timestamp     T               `json:"timestamp"`
}

func (r PaginatedResponse[T, D]) Data() ([]D, error) {
	if r.StatusCode >= StatusCodeSuccess && r.StatusCode < StatusCodeClientError {
		var o []D
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

func (r PaginatedResponse[T, D]) StrictData() ([]D, error) {
	if r.StatusCode >= StatusCodeSuccess && r.StatusCode < StatusCodeClientError {
		var o []D
		if len(r.RawData) == 0 {
			return o, nil
		}
		if err := json.Unmarshal(r.RawData, &o); err != nil {
			return nil, err
		}
		var s = struct {
			v []D `validate:"omitempty,dive,required"`
		}{o}
		if err := validate.Struct(s); err != nil {
			return nil, err
		}
		return o, nil
	}
	return nil, NewOCPIError(r.StatusCode, r.StatusMessage)
}

func (r *PaginatedResponse[T, D]) ScanHeaders(headers http.Header) error {
	r.headers = headers.Clone()
	return nil
}

func (r PaginatedResponse[T, D]) Link() string {
	return r.headers.Get(HttpHeaderLink)
}

func (r PaginatedResponse[T, D]) TotalCount() (uint64, error) {
	return strconv.ParseUint(r.headers.Get(HttpHeaderXTotalCount), 10, 64)
}

func (r PaginatedResponse[T, D]) Limit() (int, error) {
	return strconv.Atoi(r.headers.Get(HttpHeaderXLimit))
}

func NewPaginatedResponse[D any](link string, limit, totalCount int, time Timestamp, data []D) *PaginatedResponse[Timestamp, D] {
	b, _ := json.Marshal(data)
	headers := make(http.Header)
	headers.Set(HttpHeaderXLimit, strconv.Itoa(limit))
	headers.Set(HttpHeaderLink, link)
	headers.Set(HttpHeaderXTotalCount, strconv.Itoa(totalCount))
	return &PaginatedResponse[Timestamp, D]{
		headers:       headers,
		RawData:       b,
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     time,
	}
}
