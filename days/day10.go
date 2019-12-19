package days

import (
	"fmt"
	"sort"

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

// Day10Part2 solves Day 10, Part 2
func Day10Part2(input []string) (string, error) {
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
	var root utils.Point
	for asteroid, visibilityCount := range visibility {
		if visibilityCount > maxVisibility {
			maxVisibility = visibilityCount
			root = asteroid
		}
	}

	// Build the vectors from the root to every other asteroid
	vectors := map[utils.Vector][]utils.Point{}
	for _, asteroid := range asteroids {
		if !root.EqualTo(asteroid) {
			vec := root.VectorTo(asteroid)

			if vectors[vec] == nil {
				vectors[vec] = []utils.Point{asteroid}
			} else {
				vectors[vec] = append(vectors[vec], asteroid)
			}
		}
	}

	// Now build a map of angle to vector (we need this because the vectors may have different magnitudes)
	angles := map[float64][]utils.Vector{}
	for vector := range vectors {
		deg := vector.Angle()
		if angles[deg] == nil {
			angles[deg] = []utils.Vector{vector}
		} else {
			angles[deg] = append(angles[deg], vector)
		}
	}

	// Sort all angles by the magintude of the vector
	for _, vectors := range angles {
		sort.Slice(vectors, func(i, j int) bool {
			return vectors[i].Magnitude() < vectors[j].Magnitude()
		})
	}

	// Order all the angles too
	orderedAngles := []float64{}
	for angle := range angles {
		orderedAngles = append(orderedAngles, angle)
	}
	orderedAngles = sort.Float64Slice(orderedAngles)

	// Now figure out the 200th asteroid to be destroyed
	angleIndex := 0
	var theVector utils.Vector
	for asteroidNum := 0; asteroidNum < 200; asteroidNum++ {
		vectors := []utils.Vector{}
		for len(vectors) == 0 {
			deg := orderedAngles[angleIndex]
			vectors = angles[deg]
			angleIndex++
			if angleIndex >= len(orderedAngles) {
				angleIndex = 0
			}
		}

		theVector = vectors[0]
		vectors = vectors[1:] // Remove the first element from the slice
	}

	points := vectors[theVector]
	return fmt.Sprintf("%d", points[0].X*100+points[0].Y), nil
}
