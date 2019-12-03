package days_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/days"
)

func TestDay3Part1(t *testing.T) {
	cases := map[string]struct {
		input    []string
		expected string
	}{
		"example 1": {
			input: []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			expected: "159",
		},
		"example 2": {
			input: []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			expected: "135",
		},
	}

	for name, data := range cases {
		out, err := days.Day3Part1(data.input)
		require.NoErrorf(t, err, "Case: %s", name)
		require.Equalf(t, data.expected, out, "Case: %s", name)
	}
}
