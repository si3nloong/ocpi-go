package ocpi230

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"unsafe"
)

var (
	yyyymmddthhmmssRegexp     = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}$`)
	yyyymmddthhmmssnanoRegexp = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}\.\d+$`)
)

type DateTime struct {
	Time time.Time
}

func ParseDateTime(value string) (DateTime, error) {
	switch {
	case yyyymmddthhmmssnanoRegexp.MatchString(value):
		t, err := time.Parse("2006-01-02T15:04:05.999999999", value)
		if err != nil {
			return DateTime{}, err
		}
		return DateTime{Time: t}, nil
	case yyyymmddthhmmssRegexp.MatchString(value):
		t, err := time.Parse("2006-01-02T15:04:05", value)
		if err != nil {
			return DateTime{}, err
		}
		return DateTime{Time: t}, nil
	default:
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return DateTime{}, err
		}
		return DateTime{Time: t}, nil
	}
}

func (dt *DateTime) IsZero() bool {
	return dt.Time.IsZero()
}

func (dt *DateTime) UTC() DateTime {
	return DateTime{dt.Time.UTC()}
}

func (dt *DateTime) Format(layout string) string {
	return dt.Time.Format(layout)
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return dt.Time.MarshalJSON()
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	str := unsafe.String(unsafe.SliceData(b), len(b))
	str, err := strconv.Unquote(str)
	if err != nil {
		return fmt.Errorf("ocpi: unable to parse DateTime due to %w", err)
	}
	switch {
	case yyyymmddthhmmssnanoRegexp.MatchString(str):
		t, err := time.Parse("2006-01-02T15:04:05.999999999", str)
		if err != nil {
			return err
		}
		*dt = DateTime{t}
	case yyyymmddthhmmssRegexp.MatchString(str):
		t, err := time.Parse("2006-01-02T15:04:05", str)
		if err != nil {
			return err
		}
		*dt = DateTime{t}
	default:
		t, err := time.Parse(time.RFC3339, str)
		if err != nil {
			return err
		}
		*dt = DateTime{t}
	}
	return nil
}

// DisplayText defines model for cdrBody_tariffs_tariff_alt_text.
type DisplayText struct {
	Language string `json:"language" validate:"required,len=2"`
	Text     string `json:"text" validate:"required,max=512"`
}
