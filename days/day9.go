package days

import (
	"aoc/opcode"
	"fmt"
	"strconv"
	"strings"
)

// Day9Part1 solves Day 9, Part 1
func Day9Part1(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings)*1000) // Multiply to make the memory super large
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	// Send 1 to run the program in "test" mode
	in := make(chan int, 1)
	in <- 1

	out := make(chan int)

	// Run the program
	go opcode.Run(codes, in, out)

	// Get all the outputs
	outputs := []int{}
	for output := range out {
		outputs = append(outputs, output)
	}

	if len(outputs) > 1 {
		return "", fmt.Errorf("More than one output produced: %+v", outputs)
	}
	return fmt.Sprintf("%d", outputs[len(outputs)-1]), nil
}

// Day9Part2 solves Day 9, Part 2
func Day9Part2(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings)*1000) // Multiply to make the memory super large
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	// Send 2 to run the program in "boost" mode
	in := make(chan int, 1)
	in <- 2

	out := make(chan int)

	// Run the program
	go opcode.Run(codes, in, out)

	// Get all the outputs
	outputs := []int{}
	for output := range out {
		outputs = append(outputs, output)
	}

	if len(outputs) > 1 {
		return "", fmt.Errorf("More than one output produced: %+v", outputs)
	}
	return fmt.Sprintf("%d", outputs[len(outputs)-1]), nil
}
