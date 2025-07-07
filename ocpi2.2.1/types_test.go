package ocpi221

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	dt := DateTime{}
	t.Run("UnmarshalJSON", func(t *testing.T) {
		require.NoError(t, dt.UnmarshalJSON([]byte(`"2025-07-07T15:28:54Z"`)))
		require.Equal(t, dt.String(), `2025-07-07 15:28:54 +0000 UTC`)
	})
}
