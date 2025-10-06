package ocpi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	t.Run("Validate with decimal number", func(t *testing.T) {
		var n struct {
			Number json.Number `validate:"number"`
		}
		n.Number = json.Number("12.31")
		require.NoError(t, validate.StructCtx(t.Context(), n))
	})

	t.Run("Validate with integer number", func(t *testing.T) {
		var n struct {
			Number json.Number `validate:"number"`
		}
		n.Number = json.Number("9999")
		require.NoError(t, validate.StructCtx(t.Context(), n))
	})

	t.Run("Validate with decimal number using number tag", func(t *testing.T) {
		var n struct {
			Number json.Number `validate:"number=3 3"`
		}
		n.Number = json.Number("121.313")
		require.NoError(t, validate.StructCtx(t.Context(), n))
	})

	t.Run("Skip validate when using omitempty tag", func(t *testing.T) {
		t.Run("when 0", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"omitempty,number"`
			}
			n.Number = json.Number("0")
			require.NoError(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("when empty string", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"omitempty,number"`
			}
			require.NoError(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("when pointer string is nil", func(t *testing.T) {
			var n struct {
				Number *json.Number `validate:"omitempty,number"`
			}
			require.NoError(t, validate.StructCtx(t.Context(), n))
		})
	})

	t.Run("Validate with errors", func(t *testing.T) {
		t.Run("Empty string", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"number"`
			}
			require.Error(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("Empty pointer string", func(t *testing.T) {
			var n struct {
				Number *json.Number `validate:"number"`
			}
			require.Error(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("Overflow with decimal places", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"number"`
			}
			n.Number = json.Number("0.7388182")
			require.Error(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("Alphabert", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"number"`
			}
			n.Number = json.Number("abc")
			require.Error(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("Overflow selected digit", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"number=3 3"`
			}
			n.Number = json.Number("1121.313")
			require.Error(t, validate.StructCtx(t.Context(), n))
		})

		t.Run("Overflow selected decimal places", func(t *testing.T) {
			var n struct {
				Number json.Number `validate:"number=3 1"`
			}
			n.Number = json.Number("1121.313")
			require.Error(t, validate.StructCtx(t.Context(), n))
		})
	})
}
