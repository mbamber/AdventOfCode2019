package days

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/utils"
)

// Day3Part1 solves Day 3, Part 1
func Day3Part1(input []string) (string, error) {
	l1 := strings.Split(input[0], ",")
	l2 := strings.Split(input[1], ",")

	wire1, err := generateWire(l1)
	if err != nil {
		return "", err
	}
	wire2, err := generateWire(l2)
	if err != nil {
		return "", err
	}

	points := []utils.Point{}
	for _, p := range wire1 {
		if utils.PointSliceContains(wire2, p) {
			points = append(points, p)
		}
	}

	nearestPoint := points[1] // Both lists contain (0,0) at index 0
	for _, p := range points {
		if p.EqualTo(utils.NewPoint(0, 0)) {
			continue
		}

		if p.ManhattanDistance() < nearestPoint.ManhattanDistance() {
			nearestPoint = p
		}
	}

	return fmt.Sprintf("%d", nearestPoint.ManhattanDistance()), nil
}

func generateWire(instructions []string) (points []utils.Point, err error) {
	points = append(points, utils.NewPoint(0, 0))

	for _, instruction := range instructions {
		currPoints, err := getPoints(points[len(points)-1], instruction)
		if err != nil {
			return nil, err
		}
		for _, point := range currPoints {
			points = append(points, point)
		}
	}
	return points, nil
}

func getPoints(start utils.Point, instruction string) (points []utils.Point, err error) {
	dir := string(instruction[0])
	count, err := strconv.Atoi(strings.TrimLeft(instruction, "UDLR"))
	if err != nil {
		return nil, err
	}

	xShift := 0
	yShift := 0

	switch dir {
	case "U":
		yShift = 1
	case "D":
		yShift = -1
	case "L":
		xShift = -1
	case "R":
		xShift = 1
	default:
		return nil, fmt.Errorf("Unknown direction %s", dir)
	}

	for i := 0; i < count; i++ {
		newX := start.X + (i * xShift) + xShift
		newY := start.Y + (i * yShift) + yShift
		points = append(points, utils.NewPoint(newX, newY))
	}
	return points, nil
}

// Day3Part2 solves Day 3, Part 2
func Day3Part2(input []string) (string, error) {
	l1 := strings.Split(input[0], ",")
	l2 := strings.Split(input[1], ",")

	wire1, err := generateWire(l1)
	if err != nil {
		return "", err
	}
	wire2, err := generateWire(l2)
	if err != nil {
		return "", err
	}

	// Get the intersections
	intersections := getIntersections(wire1, wire2)

	nearest := 9999999999
	for _, d := range intersections {
		if d < nearest && d > 0 {
			nearest = d
		}
	}
	return fmt.Sprintf("%d", nearest), nil
}

func getIntersections(w1, w2 []utils.Point) (intersections map[utils.Point]int) {
	intersections = make(map[utils.Point]int)
	for _, p1 := range w1 {
		// Loop through w1 again to get the distance (first occurance)
		var d1, d2 int
		var tmpP1, p2 utils.Point
		for d1, tmpP1 = range w1 {
			if tmpP1.EqualTo(p1) {
				break
			}
		}
		// Calculate how far through w2 we are
		for d2, p2 = range w2 {
			if p1.EqualTo(p2) {
				intersections[p1] = d1 + d2
			}
		}
	}
	return intersections
}
