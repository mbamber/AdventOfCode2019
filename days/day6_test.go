package days_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/days"
)

func TestDay6Part1(t *testing.T) {
	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	out, err := days.Day6Part1(input)
	require.NoError(t, err)
	require.Equal(t, "42", out)
}
