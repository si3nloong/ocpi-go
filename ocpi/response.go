package ocpi

import (
	"encoding/json"
	"reflect"
)

type Response[T Timestamp, D any] struct {
	RawData       json.RawMessage `json:"data,omitempty"`
	StatusCode    StatusCode      `json:"status_code"`
	StatusMessage string          `json:"status_message,omitempty,omitzero"`
	Timestamp     T               `json:"timestamp"`
}

func (r *Response[T, D]) Decode(dest any) error {
	if err := json.Unmarshal(r.RawData, dest); err != nil {
		return err
	}
	return nil
}

func (r *Response[T, D]) Data() (D, error) {
	var o D
	if r.StatusCode >= StatusCodeSuccess && r.StatusCode < StatusCodeClientError {
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

func (r *Response[T, D]) StrictData() (D, error) {
	var o D
	if r.StatusCode >= StatusCodeSuccess && r.StatusCode < StatusCodeClientError {
		if len(r.RawData) == 0 {
			return o, nil
		}
		if err := json.Unmarshal(r.RawData, &o); err != nil {
			return o, err
		}
		if reflect.TypeOf(o).Kind() == reflect.Struct {
			if err := validate.Struct(o); err != nil {
				return o, err
			}
		}
		return o, nil
	}
	return o, NewOCPIError(r.StatusCode, r.StatusMessage)
}

func NewResponse[D any](ts Timestamp, data D) *Response[Timestamp, D] {
	b, _ := json.Marshal(data)
	return &Response[Timestamp, D]{
		RawData:       b,
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     ts,
	}
}

func NewEmptyResponse(ts Timestamp) *Response[Timestamp, any] {
	return &Response[Timestamp, any]{
		StatusCode:    StatusCodeSuccess,
		StatusMessage: StatusCodeSuccess.String(),
		Timestamp:     ts,
	}
}
