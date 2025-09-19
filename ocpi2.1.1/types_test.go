package ocpi211

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	dt := DateTime{}
	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2015-06-29T20:39:09Z"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))

		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, []byte(`"2015-06-29T20:39:09Z"`))

		require.Equal(t, dt.String(), `2015-06-29T20:39:09Z`)
	})

	t.Run("DateTime", func(t *testing.T) {
		rawData := []byte(`"2015-06-29T20:39:09"`)
		require.NoError(t, dt.UnmarshalJSON(rawData))

		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, b, []byte(`"2015-06-29T20:39:09Z"`))

		require.Equal(t, dt.String(), `2015-06-29T20:39:09Z`)
	})
}
