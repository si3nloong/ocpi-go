package ocpi221

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const timeLayout = "2006-01-02T15:04:05.999Z"

var (
	yyyymmddthhmmsszRegexp     = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}Z$`)
	yyyymmddthhmmssRegexp      = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}$`)
	yyyymmddthhmmssnanozRegexp = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}\.(\d{0,3})Z$`)
	yyyymmddthhmmssnanoRegexp  = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}T\d{2}\:\d{2}\:\d{2}\.(\d{0,3})$`)
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
	case yyyymmddthhmmssnanozRegexp.MatchString(value):
		submatches := yyyymmddthhmmssnanozRegexp.FindStringSubmatch(value)
		t, err := time.Parse("2006-01-02T15:04:05."+strings.Repeat("9", len(submatches[1]))+"Z", value)
		if err != nil {
			return DateTime{}, err
		}
		return DateTime{Time: t}, nil
	case yyyymmddthhmmssnanoRegexp.MatchString(value):
		submatches := yyyymmddthhmmssnanoRegexp.FindStringSubmatch(value)
		t, err := time.Parse("2006-01-02T15:04:05."+strings.Repeat("9", len(submatches[1])), value)
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

func (dt DateTime) String() string {
	b := make([]byte, 0, len(timeLayout))
	b = dt.Time.AppendFormat(b, timeLayout)
	return string(b)
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = dt.Time.AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	str := unsafe.String(unsafe.SliceData(b), len(b))
	str, err := strconv.Unquote(str)
	if err != nil {
		return fmt.Errorf("ocpi221: unable to parse DateTime due to %w", err)
	}
	*dt, err = ParseDateTime(str)
	return err
}

// DisplayText defines model for cdrBody_tariffs_tariff_alt_text.
type DisplayText struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}
