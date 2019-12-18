package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/utils"
)

func TestGCD(t *testing.T) {
	cases := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"GCD of 18 and 12 is 6": {
			a:        18,
			b:        12,
			expected: 6,
		},
		"GCD of 12 and 18 is 6": {
			a:        12,
			b:        18,
			expected: 6,
		},
		"GCD of 13 and 14 is 1": {
			a:        13,
			b:        14,
			expected: 1,
		},
		"GCD of 2 and -4 is 2": {
			a:        2,
			b:        -4,
			expected: 2,
		},
		"GCD of -2 and -4 is -2": {
			a:        -2,
			b:        -4,
			expected: -2,
		},
	}

	for name, data := range cases {
		out := utils.GCD(data.a, data.b)
		require.Equalf(t, data.expected, out, "Case %s", name)
	}
}
