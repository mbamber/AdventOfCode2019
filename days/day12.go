package days

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Day12Part1 solves Day 12, Part 1
func Day12Part1(input []string) (string, error) {
	moons := []Moon{}
	r, err := regexp.Compile(`x=(-?\d+).*y=(-?\d+).*z=(-?\d+)`)
	if err != nil {
		return "", err
	}

	for _, inputLine := range input {
		match := r.FindStringSubmatch(inputLine)
		m, err := NewMoonFromStrings(match[1], match[2], match[3])
		if err != nil {
			return "", err
		}

		moons = append(moons, m)
	}

	for i := 0; i < 1000; i++ {
		DoMoonTimestep(moons)
	}

	totalEngery := 0
	for _, moon := range moons {
		totalEngery += moon.Energy()
	}

	return fmt.Sprintf("%d", totalEngery), nil
}

// DoMoonTimestep simulates a single timestep for a slice of Moons
func DoMoonTimestep(moons []Moon) {
	// Adjust for Gravity
	for i := 0; i < len(moons); i++ {
		for j := i + 1; j < len(moons); j++ {
			moons[i].adjustForGravityWith(moons[j])
			moons[j].adjustForGravityWith(moons[i])
		}
	}

	// Adjust for Velocity
	for i := 0; i < len(moons); i++ {
		moons[i].applyVelocity()
	}
}

// Moon represents a moon
type Moon struct {
	X, Y, Z    int
	Dx, Dy, Dz int
}

// NewMoon creates and returns a new Moon at the given coordinates
func NewMoon(x, y, z int) Moon {
	return Moon{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewMoonFromStrings is the same as NewMoon, but with strings
func NewMoonFromStrings(x, y, z string) (Moon, error) {
	xint, err := strconv.Atoi(x)
	if err != nil {
		return Moon{}, err
	}

	yint, err := strconv.Atoi(y)
	if err != nil {
		return Moon{}, err
	}

	zint, err := strconv.Atoi(z)
	if err != nil {
		return Moon{}, err
	}

	return NewMoon(xint, yint, zint), nil
}

func (m *Moon) adjustForGravityWith(m2 Moon) {
	if m.X > m2.X {
		m.Dx--
	} else if m.X < m2.X {
		m.Dx++
	}

	if m.Y > m2.Y {
		m.Dy--
	} else if m.Y < m2.Y {
		m.Dy++
	}

	if m.Z > m2.Z {
		m.Dz--
	} else if m.Z < m2.Z {
		m.Dz++
	}
}

func (m *Moon) applyVelocity() {
	m.X += m.Dx
	m.Y += m.Dy
	m.Z += m.Dz
}

// PotentialEnergy returns the potential energy for a moon
func (m *Moon) PotentialEnergy() int {
	return int(math.Abs(float64(m.X)) + math.Abs(float64(m.Y)) + math.Abs(float64(m.Z)))
}

// KineticEnergy returns the kinetic energy for a moon
func (m *Moon) KineticEnergy() int {
	return int(math.Abs(float64(m.Dx)) + math.Abs(float64(m.Dy)) + math.Abs(float64(m.Dz)))
}

// Energy returns the total energy for a moon (potential * kinetic)
func (m *Moon) Energy() int {
	return m.PotentialEnergy() * m.KineticEnergy()
}
