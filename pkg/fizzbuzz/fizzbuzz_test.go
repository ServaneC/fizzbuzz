package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		name     string
		input    Params
		expected []string
	}{
		{
			"basic",
			Params{3, 5, 16, "fizz", "buzz"},
			[]string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"},
		},
	}
	for _, tc := range testCases {
		result := Generate(tc.input)
		require.Equal(t, tc.expected, result)
	}
}
