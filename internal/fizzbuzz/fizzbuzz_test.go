package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		name     string
		input    Params
		expected []string
	}{
		{
			"classic",
			Params{3, 5, 16, "fizz", "buzz"},
			[]string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"},
		},
		{
			name:     "limit 1",
			input:    Params{Int1: 3, Int2: 5, Limit: 1, Str1: "fizz", Str2: "buzz"},
			expected: []string{"1"},
		},
		{
			name:  "both divisors same",
			input: Params{Int1: 2, Int2: 2, Limit: 4, Str1: "a", Str2: "b"},
			expected: []string{
				"1", "ab", "3", "ab",
			},
		},
		{
			name:  "no multiples",
			input: Params{Int1: 50, Int2: 60, Limit: 5, Str1: "a", Str2: "b"},
			expected: []string{
				"1", "2", "3", "4", "5",
			},
		},
		{
			name:  "multiples overlap",
			input: Params{Int1: 2, Int2: 3, Limit: 6, Str1: "foo", Str2: "bar"},
			expected: []string{
				"1", "foo", "bar", "foo", "5", "foobar",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Generate(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
