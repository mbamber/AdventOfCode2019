package days

import (
	"aoc/opcode"
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

// Day11Part1 solves Day 11, Part 1
func Day11Part1(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings)*1000)
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	// Make the cans
	in, out := make(chan int, 10), make(chan int, 10)

	// Process the outputs in a subroutine
	coveredPoints := map[utils.Point]int{}
	go func(in, out chan int, coveredPoints map[utils.Point]int) {
		startingPoint := utils.NewPoint(0, 0)
		direction := utils.NewVector(0, 1)

		for {
			// Send the color that we are already over
			col, exists := coveredPoints[startingPoint]
			if exists {
				in <- col
			} else {
				in <- 0
			}
			// Read the new color and save it
			newCol := <-out
			coveredPoints[startingPoint] = newCol

			// Read the new direction
			dir, open := <-out
			if !open {
				// The program has terminated
				return
			}

			// Move
			switch direction.Angle() {
			case 0.0:
				if dir == 0 {
					direction = utils.NewVector(-1, 0)
				} else {
					direction = utils.NewVector(1, 0)
				}
			case 90.0:
				if dir == 0 {
					direction = utils.NewVector(0, 1)
				} else {
					direction = utils.NewVector(0, -1)
				}
			case 180.0:
				if dir == 0 {
					direction = utils.NewVector(1, 0)
				} else {
					direction = utils.NewVector(-1, 0)
				}
			case 270.0:
				if dir == 0 {
					direction = utils.NewVector(0, -1)
				} else {
					direction = utils.NewVector(0, 1)
				}
			}
			startingPoint = startingPoint.PlusVector(direction)
		}
	}(in, out, coveredPoints)

	// Start the program
	opcode.Run(codes, in, out)

	return fmt.Sprintf("%d", len(coveredPoints)), nil
}
