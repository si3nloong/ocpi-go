package ocpi211

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	dt := DateTime{}
	t.Run("MarshalJSON", func(t *testing.T) {
		b, err := dt.MarshalJSON()
		require.NoError(t, err)
		require.Equal(t, 22, len(b))
	})
}
