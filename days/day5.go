package days

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/opcode"
)

// Day5Part1 solves Day 5, Part 1
func Day5Part1(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings))
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	out, err := opcode.Run(codes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", out), nil
}
