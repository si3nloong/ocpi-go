package ocpi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

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

var typeOfJsonNumber = reflect.TypeOf(json.Number(""))

func init() {
	validate = validator.New()
	noError(validate.RegisterValidation("number", func(fl validator.FieldLevel) bool {
		val := fl.Field()

		var precision, scale int
		// Parse params
		params := strings.Split(fl.Param(), " ")
		if len(params) != 2 {
			precision = 12
			scale = 4
		} else {
			var err1, err2 error
			precision, err1 = strconv.Atoi(strings.TrimSpace(params[0]))
			scale, err2 = strconv.Atoi(strings.TrimSpace(params[1]))
			if err1 != nil || err2 != nil {
				return false
			}
		}

		var s string
		if fl.Field().Type() == typeOfJsonNumber {
			s = fl.Field().Interface().(fmt.Stringer).String()
			if _, err := strconv.ParseFloat(s, 64); err != nil {
				return false
			}
		} else if val.Kind() == reflect.Float32 || val.Kind() == reflect.Float64 {
			s = fmt.Sprintf("%.20f", val.Float())
		} else {
			return false
		}

		// Format with enough decimals
		s = strings.TrimRight(s, "0")
		parts := strings.Split(s, ".")

		intDigits := len(parts[0])
		decDigits := 0
		if len(parts) == 2 {
			decDigits = len(parts[1])
		}
		// total digits must fit precision, decimals must fit scale
		return intDigits <= precision && decDigits <= scale
	}))
	noError(validate.RegisterValidation("version", func(fl validator.FieldLevel) bool {
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
	}))
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
	return o, nil
}

func (r RawMessage[T]) StrictData() (T, error) {
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

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
