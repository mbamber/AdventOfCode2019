package utils

import (
	"math"
)

// Point represents a point in Cartesian space
type Point struct {
	X, Y int
}

// EqualTo returns true iff the two points have the same X and Y values
func (p *Point) EqualTo(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}

// ManhattanDistance calculates the manhattan distance to the origin (0,0)
func (p *Point) ManhattanDistance() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

// ManhattanDistanceBetween calculates the manhattan distance to the given point
func (p *Point) ManhattanDistanceBetween(point Point) int {
	return int(math.Abs(float64(point.X-p.X)) + math.Abs(float64(point.Y-p.Y)))
}

// NewPoint creates and returns a new Point
func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// PointSliceContains returns true iff the slice of Points contains `p`
func PointSliceContains(s []Point, p Point) bool {
	for _, point := range s {
		if p.EqualTo(point) {
			return true
		}
	}
	return false
}

// PointsAtManhattanDistance returns all the points that are at a given Manhattan distance from the origin (0,0)
func PointsAtManhattanDistance(d int) (points []Point) {
	for x := -d; x <= d; x++ {
		for y := -d; y <= d; y++ {
			if x+y == d || x+y == -d {
				points = append(points, NewPoint(x, y))
			}
		}
	}
	return points
}
