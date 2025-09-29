package ocpi230

import (
	"testing"

	"github.com/si3nloong/ocpi-go/ocpi"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	type S struct {
		AuthMethod AuthMethod `json:"authMethod" validate:"authMethod230"`
	}

	t.Run("Validation not passed", func(t *testing.T) {
		rawMessage := ocpi.RawMessage[S](`{"authMethod":"ss"}`)
		data, err := rawMessage.StrictData()
		require.Error(t, err)
		require.Equal(t, data.AuthMethod, AuthMethod("ss"))
	})

	t.Run("Validation passed", func(t *testing.T) {
		rawMessage := ocpi.RawMessage[S](`{"authMethod":"COMMAND"}`)
		data, err := rawMessage.StrictData()
		require.NoError(t, err)
		require.Equal(t, data.AuthMethod, AuthMethodCommand)
	})
}
