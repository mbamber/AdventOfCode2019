package days

import (
	"aoc/opcode"
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
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

	// Make the chans
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

// Day11Part2 solves Day 11, Part 2
func Day11Part2(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings)*1000)
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	// Make the chans
	in, out := make(chan int, 10), make(chan int, 10)

	// Process the outputs in a subroutine
	coveredPoints := map[utils.Point]int{}
	locker := sync.Mutex{}
	go func(in, out chan int) {
		startingPoint := utils.NewPoint(0, 0)
		direction := utils.NewVector(0, 1)

		defaultCol := 1
		for {
			// Send the color that we are already over
			col, exists := coveredPoints[startingPoint]
			if exists {
				in <- col
			} else {
				in <- defaultCol
				defaultCol = 0
			}
			// Read the new color and save it
			newCol := <-out
			locker.Lock()
			coveredPoints[startingPoint] = newCol
			locker.Unlock()

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
	}(in, out)

	// Start the program
	opcode.Run(codes, in, out)

	// Now we have all the points that have been written, so just draw it
	minX, maxX, minY, maxY := 0, 0, 0, 0
	locker.Lock()
	for point := range coveredPoints {
		if point.X < minX {
			minX = point.X
		}

		if point.Y < minY {
			minY = point.Y
		}

		if point.X > maxX {
			maxX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}
	locker.Unlock()

	width := int(math.Abs(float64(minX))+math.Abs(float64(maxX))) + 1
	height := int(math.Abs(float64(minY))+math.Abs(float64(maxY))) + 1
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	locker.Lock()
	for point, col := range coveredPoints {
		// Need to shift the coordinates here to make them inside the range
		absX, absY := int(math.Abs(float64(point.X))), int(math.Abs(float64(point.Y)))
		grid[absY][absX] = col
	}
	locker.Unlock()

	hull := ""
	for _, row := range grid {
		for _, col := range row {
			if col == 0 {
				hull += "."
			} else {
				hull += "#"
			}
		}
		hull += "\n"
	}

	return hull, nil
}
