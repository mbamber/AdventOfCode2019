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
	return Vector{
		XShift: point.X - p.X,
		YShift: point.Y - p.Y,
	}
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

// Vector represents a directional vector. Note the positive directions are North and East
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

// Angle returns the number of degrees from (0, 1) (Due North) represented by this Vector. This is a _really_ bad
// way of working this out
func (v *Vector) Angle() float64 {
	// Handle the edge cases where the angle is exactly a compass direction
	if v.XShift == 0 {
		if v.YShift > 0 {
			return 0.0
		} else if v.YShift < 0 {
			return 180.0
		}
		return math.NaN()
	}

	if v.YShift == 0 {
		if v.XShift > 0 {
			return 90.0
		} else if v.XShift < 0 {
			return 270.0
		}
		return math.NaN()
	}

	fy, fx := float64(v.YShift), float64(v.XShift)
	opp, adj := math.Abs(fy), math.Abs(fx)
	rad := math.Atan(opp / adj)
	deg := (360.0 / (2 * math.Pi) * rad)

	// Because we've shifted this round a little (i.e. angle measured clockwise from North) we need to do some adjustments
	if fy > 0 {
		if fx > 0 {
			return 90 - deg
		}
		return 270 + deg
	}
	if fx > 0 {
		return 90 + deg
	}
	return 270 - deg
}

// Magnitude returns the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(float64(v.XShift), 2) + math.Pow(float64(v.YShift), 2))
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
			v = v.Simplified()
			if !VectorSliceContains(vectors, v) {
				vectors = append(vectors, v)
			}

		}
	}

	return vectors, nil
}
