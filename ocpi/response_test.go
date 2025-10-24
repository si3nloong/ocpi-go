package ocpi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestResponse(t *testing.T) {
	t.Run("Response", func(t *testing.T) {
		type A struct {
			Name string `validate:"required"`
		}
		var o Response[A]
		o.RawData = []byte(`[{}]`)
		o.StatusCode = StatusCodeSuccess
		o.Timestamp = time.Now()
		result, err := o.StrictData()
		require.NoError(t, err)
		_ = result
		// require.ElementsMatch(t, result, []A{})
	})
}
