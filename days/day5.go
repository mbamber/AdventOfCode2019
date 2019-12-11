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

	in := make(chan int)
	out := make(chan int)

	var err error
	go func() {
		err = opcode.Run(codes, in, out)
	}()

	go func() {
		in <- 1
	}()

	output, open := <-out
	var prevOut int
	for open {
		prevOut = output
		output, open = <-out
	}

	// Now check the error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", prevOut), nil
}

// Day5Part2 solves Day 5, Part 2
func Day5Part2(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings))
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	in := make(chan int)
	out := make(chan int)

	var err error
	go func() {
		err = opcode.Run(codes, in, out)
	}()

	go func() {
		in <- 5
	}()

	output, open := <-out
	var prevOut int
	for open {
		prevOut = output
		output, open = <-out
	}

	// Now check the error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", prevOut), nil
}
