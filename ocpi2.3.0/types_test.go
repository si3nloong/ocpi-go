package ocpi230

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	dt := DateTime{}
	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2025-07-07T15:28:54Z"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, rawData)
	})

	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2015-06-29T20:39:09"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, []byte(`"2015-06-29T20:39:09Z"`))
	})

	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2016-12-29T17:45:09.2Z"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, rawData)
	})

	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2016-12-29T17:45:09.2"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, []byte(`"2016-12-29T17:45:09.2Z"`))
	})

	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2018-01-01T01:08:01.123Z"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, rawData)
	})

	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2018-01-01T01:08:01.123"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, []byte(`"2018-01-01T01:08:01.123Z"`))
	})
}
