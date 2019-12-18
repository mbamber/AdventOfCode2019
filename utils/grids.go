package utils

import (
	"fmt"
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

// VectorTo returns the vector to the given point
func (p *Point) VectorTo(point Point) Vector {
	return NewVector(point.X-p.X, point.Y-p.Y)
}

// PlusVector adds a vector to a point, and returns the new point
func (p *Point) PlusVector(v Vector) Point {
	return NewPoint(p.X+v.XShift, p.Y+v.YShift)
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

// Vector represents a directional vector
type Vector struct {
	XShift int
	YShift int
}

// EqualTo returns true iff the two vectors are the same
func (v *Vector) EqualTo(vector Vector) bool {
	simpleV := v.Simplified()
	simpleVector := vector.Simplified()
	return simpleV.XShift == simpleVector.XShift && simpleV.YShift == simpleVector.YShift
}

// Simplified returns the given vector, but simplified
func (v *Vector) Simplified() Vector {
	gcd := GCD(v.XShift, v.YShift)
	if gcd == 0 {
		return Vector{
			XShift: 0,
			YShift: 0,
		}
	} else if gcd < 0 {
		// We actually want the absolute value of the GCD, so make sure it's positive
		gcd = -1 * gcd
	}
	return Vector{
		XShift: v.XShift / gcd,
		YShift: v.YShift / gcd,
	}
}

// ScaledBy returns a vector that has been scaled by the given factor
func (v *Vector) ScaledBy(factor int) Vector {
	return Vector{
		XShift: v.XShift * factor,
		YShift: v.YShift * factor,
	}
}

// NewVector creates and returns a new, simplified Vector
func NewVector(x, y int) Vector {
	v := Vector{
		XShift: x,
		YShift: y,
	}
	return v.Simplified()
}

// VectorSliceContains returns true iff the slice of Vectors contains `v`
func VectorSliceContains(s []Vector, v Vector) bool {
	for _, vector := range s {
		if v.EqualTo(vector) {
			return true
		}
	}
	return false
}

// GetAllVectorsFrom returns a slice of all the simplified vectors from the given point to every point within a grid of
// the specified width and height. Returns an error if the given point falls outside of the specified grid dimensions.
func GetAllVectorsFrom(width, height int, p Point) ([]Vector, error) {
	if width < 0 || height < 0 {
		return nil, fmt.Errorf("Width (%d) and Height (%d) must both be greater than 0", width, height)
	}

	if width <= p.X || height <= p.Y {
		return nil, fmt.Errorf("Given point %v does not fall inside the specified width (%d) and height (%d)", p, width, height)
	}

	// Build all the vectors from the asteroid
	vectors := []Vector{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			to := NewPoint(x, y)

			// Skip this point if it's where we're starting from
			if p.EqualTo(to) {
				continue
			}

			// Calculate the vector, and add it to the list if it's not there already
			v := p.VectorTo(to)
			if !VectorSliceContains(vectors, v) {
				vectors = append(vectors, v)
			}

		}
	}

	return vectors, nil
}
