package days_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/days"
)

func TestDay7Part1(t *testing.T) {
	cases := map[string]struct {
		program  []string
		maxValue string
	}{
		"Example 1": {
			program:  []string{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"},
			maxValue: "43210",
		},
		"Example 2": {
			program:  []string{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"},
			maxValue: "54321",
		},
		"Example 3": {
			program:  []string{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"},
			maxValue: "65210",
		},
	}

	for name, data := range cases {
		out, err := days.Day7Part1(data.program)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equal(t, data.maxValue, out, "Case %s", name)
	}
}
