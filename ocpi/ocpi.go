package ocpi

import (
	"encoding/json"
	"time"
)

// Version defines model for DetailsData.Version.
type Version string

// Defines values for Version.
const (
	VersionN20  Version = "2.0"
	VersionN21  Version = "2.1"
	VersionN211 Version = "2.1.1"
	VersionN22  Version = "2.2"
	VersionN221 Version = "2.2.1"
)

type RawMessage[T any] json.RawMessage

func (r RawMessage[T]) Data() (T, error) {
	var o T
	if err := json.Unmarshal((json.RawMessage)(r), &o); err != nil {
		return o, err
	}
	return o, nil
}

type Result[T any] interface {
	Data() (T, error)
}

type ocpiResult[T any] struct {
	resp *Response[T]
	err  error
}

func (r *ocpiResult[T]) Data() (T, error) {
	if r.err != nil {
		var o T
		return o, r.err
	} else if r.resp == nil {
		var o T
		return o, nil
	}
	return r.resp.Data, nil
}

func NewResult[T any](resp Response[T]) Result[T] {
	return &ocpiResult[T]{resp: &resp}
}

type Response[T any] struct {
	Data          T          `json:"data,omitempty"`
	StatusCode    StatusCode `json:"status_code"`
	StatusMessage string     `json:"status_message"`
	Timestamp     time.Time  `json:"timestamp"`
}
