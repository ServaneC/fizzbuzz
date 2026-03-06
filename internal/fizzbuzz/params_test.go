package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamsValidation(t *testing.T) {
	tests := []struct {
		name  string
		input Params
		valid bool
	}{
		{
			"valid params",
			Params{3, 5, 10, "fizz", "buzz"},
			true,
		},
		{
			"zero int1",
			Params{0, 5, 10, "fizz", "buzz"},
			false,
		},
		{
			"zero limit",
			Params{3, 5, 0, "fizz", "buzz"},
			false,
		},
		{
			"empty str1",
			Params{3, 5, 10, "", "buzz"},
			false,
		},
		{
			"empty str2",
			Params{3, 5, 10, "fizz", ""},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.valid, tt.input.IsValid() == nil)
		})
	}
}
