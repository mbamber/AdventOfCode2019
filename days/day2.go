package days

import (
	"errors"
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

// Day2Part2 solves Day 2, Part 2
func Day2Part2(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings))
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			// Copy the original input so we don't use a stale version
			currCodes := append(codes[:0:0], codes...)

			// Set the inputs
			currCodes[1] = noun
			currCodes[2] = verb

			res, err := process(currCodes)
			if err != nil {
				return "", err
			}
			if res == 19690720 {
				return fmt.Sprintf("%d", 100*noun+verb), nil
			}
		}
	}
	return "", errors.New("Could not solve")
}
