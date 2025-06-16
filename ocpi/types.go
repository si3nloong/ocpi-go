package ocpi

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

func (dt *DateTime) IsZero() bool {
	return dt.Time.IsZero()
}

func (dt *DateTime) UTC() DateTime {
	return DateTime{dt.Time.UTC()}
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
