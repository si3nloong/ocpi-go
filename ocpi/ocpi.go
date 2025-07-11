package ocpi

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Version defines model for DetailsData.Version.
type VersionNumber string

// Defines values for Version.
const (
	VersionNumber20  VersionNumber = "2.0"
	VersionNumber21  VersionNumber = "2.1" // deprecated
	VersionNumber211 VersionNumber = "2.1.1"
	VersionNumber22  VersionNumber = "2.2" // deprecated
	VersionNumber221 VersionNumber = "2.2.1"
	VersionNumber230 VersionNumber = "2.3.0"
)

type HeaderScanner interface {
	ScanHeader(httpHeader http.Header) error
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("version", func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().(VersionNumber)
		if !ok {
			return false
		}
		switch v {
		case VersionNumber20, VersionNumber21, VersionNumber211,
			VersionNumber22, VersionNumber221, VersionNumber230:
			return true
		default:
			return false
		}
	}); err != nil {
		panic(err)
	}
}

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
