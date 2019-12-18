package days

import (
	"fmt"

	"aoc/utils"
)

// Day10Part1 solves Day 10, Part 1
func Day10Part1(input []string) (string, error) {

	asteroids := []utils.Point{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				asteroids = append(asteroids, utils.NewPoint(x, y))
			}
		}
	}

	visibility := map[utils.Point]int{}
	for _, asteroid := range asteroids {
		visibility[asteroid] = getAsteroidVisibility(input, asteroid)
	}

	maxVisibility := 0
	for _, visibilityCount := range visibility {
		if visibilityCount > maxVisibility {
			maxVisibility = visibilityCount
		}
	}

	return fmt.Sprintf("%d", maxVisibility), nil
}

func getAsteroidVisibility(space []string, asteroid utils.Point) int {
	astCount := 0

	// Build all the vectors from the asteroid
	vectors, err := utils.GetAllVectorsFrom(len(space[0]), len(space), asteroid)
	if err != nil {
		return 0
	}

	// For each vector, keep adding multiples until either we find an asteroid or we go outside the boundaries of space
	for _, vector := range vectors {
		multiplier := 0
		for {
			multiplier++
			scaledVector := vector.ScaledBy(multiplier)
			visiblePoint := asteroid.PlusVector(scaledVector)

			if visiblePoint.Y >= len(space) || visiblePoint.Y < 0 || visiblePoint.X >= len(space[visiblePoint.Y]) || visiblePoint.X < 0 {
				break
			}

			if space[visiblePoint.Y][visiblePoint.X] == '#' {
				astCount++
				break
			}
		}
	}

	return astCount
}
