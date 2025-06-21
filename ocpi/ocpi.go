package ocpi

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Version defines model for DetailsData.Version.
type VersionNumber string

// Defines values for Version.
const (
	VersionNumber20  VersionNumber = "2.0"
	VersionNumber21  VersionNumber = "2.1" // deprecated
	VersionNumber211 VersionNumber = "2.1.1"
	VersionNumber22  VersionNumber = "2.2"
	VersionNumber221 VersionNumber = "2.2.1"
	VersionNumber230 VersionNumber = "2.3.0"
)

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("version", func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().(VersionNumber)
		if !ok {
			return false
		}
		switch v {
		case VersionNumber20, VersionNumber21, VersionNumber211,
			VersionNumber22, VersionNumber221:
			return true
		default:
			return false
		}
	}); err != nil {
		panic(err)
	}

}

// ModuleID defines the OCPI module identifier.

type RawMessage[T any] json.RawMessage

func (r RawMessage[T]) Data() (T, error) {
	var o T
	if any(o) == nil && r == nil {
		return o, nil
	}
	if err := json.Unmarshal((json.RawMessage)(r), &o); err != nil {
		return o, err
	}
	if err := validate.Struct(o); err != nil {
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
