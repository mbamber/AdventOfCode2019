package days

import (
	"fmt"
	"strconv"
	"strings"
)

// Day2Part1 solves Day 2, Part 1
func Day2Part1(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings))
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	// Set the initial values
	codes[1] = 12
	codes[2] = 2

	res, err := process(codes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", res), nil
}

func process(codes []int) (int, error) {
	i := 0

	for {
		op := codes[i]
		i1 := codes[i+1]
		i2 := codes[i+2]
		dest := codes[i+3]
		switch op {
		case 1:
			codes[dest] = codes[i1] + codes[i2]
		case 2:
			codes[dest] = codes[i1] * codes[i2]
		case 99:
			return codes[0], nil
		default:
			return 0, fmt.Errorf("Invalid opcode %d", op)
		}

		i += 4
	}
}
