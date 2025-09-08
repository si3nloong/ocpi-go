package ocpi221

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"unsafe"
)

var (
	yyyymmddthhmmsszRegexp    = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}Z$`)
	yyyymmddthhmmssRegexp     = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}$`)
	yyyymmddthhmmssnanoRegexp = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}\.\d+$`)
)

func ParseDateTime(value string) (DateTime, error) {
	switch {
	case yyyymmddthhmmsszRegexp.MatchString(value):
		t, err := time.Parse("2006-01-02T15:04:05Z", value)
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

type DateTime struct {
	Time time.Time
}

func (dt DateTime) String() string {
	return dt.Time.String()
}

func (dt DateTime) IsZero() bool {
	return dt.Time.IsZero()
}

func (dt DateTime) UTC() time.Time {
	return dt.Time.UTC()
}

func (dt DateTime) In(loc *time.Location) time.Time {
	return dt.Time.In(loc)
}

func (dt DateTime) Format(layout string) string {
	return dt.Time.Format(layout)
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return dt.Time.MarshalJSON()
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	str := unsafe.String(unsafe.SliceData(b), len(b))
	str, err := strconv.Unquote(str)
	if err != nil {
		return fmt.Errorf("ocpi221: unable to parse DateTime due to %w", err)
	}
	switch {
	case yyyymmddthhmmsszRegexp.MatchString(str):
		t, err := time.ParseInLocation("2006-01-02T15:04:05Z", str, time.UTC)
		if err != nil {
			return err
		}
		*dt = DateTime{t}
	case yyyymmddthhmmssnanoRegexp.MatchString(str):
		t, err := time.ParseInLocation("2006-01-02T15:04:05.999999999", str, time.UTC)
		if err != nil {
			return err
		}
		*dt = DateTime{t}
	case yyyymmddthhmmssRegexp.MatchString(str):
		t, err := time.ParseInLocation("2006-01-02T15:04:05", str, time.UTC)
		if err != nil {
			return err
		}
		*dt = DateTime{t}

	default:
		t, err := time.Parse(time.RFC3339Nano, str)
		if err != nil {
			return err
		}
		*dt = DateTime{t}
	}
	return nil
}

// DisplayText defines model for cdrBody_tariffs_tariff_alt_text.
type DisplayText struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}
