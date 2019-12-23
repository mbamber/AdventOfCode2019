package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoonMotion(t *testing.T) {

	type moonDetails struct {
		x, y, z    int
		dx, dy, dz int
	}

	cases := map[string]struct {
		input           []days.Moon
		expectedOutputs map[int][]moonDetails
	}{
		"Example 1": {
			input: []days.Moon{
				days.NewMoon(-1, 0, 2),
				days.NewMoon(2, -10, -7),
				days.NewMoon(4, -8, 8),
				days.NewMoon(3, 5, -1),
			},
			expectedOutputs: map[int][]moonDetails{
				1: []moonDetails{
					moonDetails{
						x:  2,
						y:  -1,
						z:  1,
						dx: 3,
						dy: -1,
						dz: -1,
					},
					moonDetails{
						x:  3,
						y:  -7,
						z:  -4,
						dx: 1,
						dy: 3,
						dz: 3,
					},
					moonDetails{
						x:  1,
						y:  -7,
						z:  5,
						dx: -3,
						dy: 1,
						dz: -3,
					},
					moonDetails{
						x:  2,
						y:  2,
						z:  0,
						dx: -1,
						dy: -3,
						dz: 1,
					},
				},
				2: []moonDetails{
					moonDetails{
						x:  5,
						y:  -3,
						z:  -1,
						dx: 3,
						dy: -2,
						dz: -2,
					},
					moonDetails{
						x:  1,
						y:  -2,
						z:  2,
						dx: -2,
						dy: 5,
						dz: 6,
					},
					moonDetails{
						x:  1,
						y:  -4,
						z:  -1,
						dx: 0,
						dy: 3,
						dz: -6,
					},
					moonDetails{
						x:  1,
						y:  -4,
						z:  2,
						dx: -1,
						dy: -6,
						dz: 2,
					},
				},
			},
		},
	}

	for name, data := range cases {
		moons := append(data.input[:0:0], data.input...)

		// Get the maximum number of steps
		maxStep := 0
		for step := range data.expectedOutputs {
			if step > maxStep {
				maxStep = step
			}
		}

		for step := 1; step <= maxStep; step++ {
			days.DoMoonTimestep(moons)
			expected, exists := data.expectedOutputs[step]
			if !exists {
				continue // No tests for this step
			}

			for i, moon := range moons {
				expectedMoonDetails := expected[i]
				require.Equalf(t, expectedMoonDetails.x, moon.X, "Case %s, Moon %+v (%d), Step %d, X", name, moon, i, step)
				require.Equalf(t, expectedMoonDetails.y, moon.Y, "Case %s, Moon %+v (%d), Step %d, Y", name, moon, i, step)
				require.Equalf(t, expectedMoonDetails.z, moon.Z, "Case %s, Moon %+v (%d), Step %d, Z", name, moon, i, step)
				require.Equalf(t, expectedMoonDetails.dx, moon.Dx, "Case %s, Moon %+v (%d), Step %d, dX", name, moon, i, step)
				require.Equalf(t, expectedMoonDetails.dy, moon.Dy, "Case %s, Moon %+v (%d), Step %d, dY", name, moon, i, step)
				require.Equalf(t, expectedMoonDetails.dz, moon.Dz, "Case %s, Moon %+v (%d), Step %d, dZ", name, moon, i, step)
			}
		}
	}
}

func TestMoonEnergy(t *testing.T) {
	type moonEnergy struct {
		pot, kin, tot int
	}

	cases := map[string]struct {
		input           []days.Moon
		expectedOutputs map[int][]moonEnergy
	}{
		"Example 1": {
			input: []days.Moon{
				days.NewMoon(-1, 0, 2),
				days.NewMoon(2, -10, -7),
				days.NewMoon(4, -8, 8),
				days.NewMoon(3, 5, -1),
			},
			expectedOutputs: map[int][]moonEnergy{
				10: []moonEnergy{
					moonEnergy{
						pot: 6,
						kin: 6,
						tot: 36,
					},
					moonEnergy{
						pot: 9,
						kin: 5,
						tot: 45,
					},
					moonEnergy{
						pot: 10,
						kin: 8,
						tot: 80,
					},
					moonEnergy{
						pot: 6,
						kin: 3,
						tot: 18,
					},
				},
			},
		},
	}

	for name, data := range cases {
		moons := append(data.input[:0:0], data.input...)

		// Get the maximum number of steps
		maxStep := 0
		for step := range data.expectedOutputs {
			if step > maxStep {
				maxStep = step
			}
		}

		for step := 1; step <= maxStep; step++ {
			days.DoMoonTimestep(moons)

			expected, exists := data.expectedOutputs[step]
			if !exists {
				continue // No tests for this step
			}

			for i, moon := range moons {
				expectedMoonEnergy := expected[i]
				require.Equalf(t, expectedMoonEnergy.pot, moon.PotentialEnergy(), "Case %s, Moon %+v (%d), Step %d, Potential Energy", name, moon, i, step)
				require.Equalf(t, expectedMoonEnergy.kin, moon.KineticEnergy(), "Case %s, Moon %+v (%d), Step %d, Kinetic Energy", name, moon, i, step)
				require.Equalf(t, expectedMoonEnergy.tot, moon.Energy(), "Case %s, Moon %+v (%d), Step %d, Total Energy", name, moon, i, step)
			}
		}
	}
}
