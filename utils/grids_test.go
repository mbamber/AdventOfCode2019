package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/utils"
)

func TestGetAlLVectorsFromErrors(t *testing.T) {
	cases := map[string]struct {
		width  int
		height int
		point  utils.Point
	}{
		"Width less than 0": {
			width:  -1,
			height: 1,
			point:  utils.NewPoint(1, 1),
		},
		"Height less than 0": {
			width:  1,
			height: -1,
			point:  utils.NewPoint(1, 1),
		},
		"X coordinate more than width": {
			width:  1,
			height: 1,
			point:  utils.NewPoint(1, 0),
		},
		"Y coordinate more than width": {
			width:  1,
			height: 1,
			point:  utils.NewPoint(0, 1),
		},
	}

	for name, data := range cases {
		_, err := utils.GetAllVectorsFrom(data.width, data.height, data.point)
		require.Errorf(t, err, "Case %s", name)
	}
}

func TestGetAllVectorsFrom(t *testing.T) {
	cases := map[string]struct {
		width    int
		height   int
		point    utils.Point
		expected []utils.Vector
	}{
		"1 x 1 Grid": {
			width:    1,
			height:   1,
			point:    utils.NewPoint(0, 0),
			expected: []utils.Vector{},
		},
		"2 x 2 Grid": {
			width:    2,
			height:   2,
			point:    utils.NewPoint(0, 0),
			expected: []utils.Vector{utils.NewVector(1, 0), utils.NewVector(1, 1), utils.NewVector(0, 1)},
		},
		"3 x 3 Grid": {
			width:    3,
			height:   3,
			point:    utils.NewPoint(0, 0),
			expected: []utils.Vector{utils.NewVector(1, 0), utils.NewVector(1, 1), utils.NewVector(0, 1), utils.NewVector(1, 2), utils.NewVector(2, 1)},
		},
		"3 x 3 Grid Centred": {
			width:    3,
			height:   3,
			point:    utils.NewPoint(1, 1),
			expected: []utils.Vector{utils.NewVector(0, -1), utils.NewVector(1, -1), utils.NewVector(1, 0), utils.NewVector(1, 1), utils.NewVector(0, 1), utils.NewVector(-1, 1), utils.NewVector(-1, 0), utils.NewVector(-1, -1)},
		},
	}

	for name, data := range cases {
		out, err := utils.GetAllVectorsFrom(data.width, data.height, data.point)
		require.NoErrorf(t, err, "Case %s", name)
		require.Len(t, out, len(data.expected), "Case %s", name)

		for _, expected := range data.expected {
			require.Truef(t, utils.VectorSliceContains(out, expected), "Case %s", name)
		}
	}
}

func TestVectorTo(t *testing.T) {
	cases := map[string]struct {
		from     utils.Point
		to       utils.Point
		expected utils.Vector
	}{
		"(0, -1)": {
			from:     utils.NewPoint(0, 1),
			to:       utils.NewPoint(0, 0),
			expected: utils.NewVector(0, -1),
		},
		"(1, -1)": {
			from:     utils.NewPoint(0, 1),
			to:       utils.NewPoint(1, 0),
			expected: utils.NewVector(1, -1),
		},
		"(1, 0)": {
			from:     utils.NewPoint(0, 0),
			to:       utils.NewPoint(1, 0),
			expected: utils.NewVector(1, 0),
		},
		"(1, 1)": {
			from:     utils.NewPoint(0, 0),
			to:       utils.NewPoint(1, 1),
			expected: utils.NewVector(1, 1),
		},
		"(0, 1)": {
			from:     utils.NewPoint(0, 0),
			to:       utils.NewPoint(0, 1),
			expected: utils.NewVector(0, 1),
		},
		"(-1, 1)": {
			from:     utils.NewPoint(1, 0),
			to:       utils.NewPoint(0, 1),
			expected: utils.NewVector(-1, 1),
		},
		"(-1, 0)": {
			from:     utils.NewPoint(1, 0),
			to:       utils.NewPoint(0, 0),
			expected: utils.NewVector(-1, 0),
		},
		"(-1, -1)": {
			from:     utils.NewPoint(1, 1),
			to:       utils.NewPoint(0, 0),
			expected: utils.NewVector(-1, -1),
		},
		"(1, 2)": {
			from:     utils.NewPoint(0, 0),
			to:       utils.NewPoint(1, 2),
			expected: utils.NewVector(1, 2),
		},
		"(0, 0)": {
			from:     utils.NewPoint(1, 1),
			to:       utils.NewPoint(1, 1),
			expected: utils.NewVector(0, 0),
		},
		"(2, 2) simplifies to (1, 1)": {
			from:     utils.NewPoint(0, 0),
			to:       utils.NewPoint(2, 2),
			expected: utils.NewVector(1, 1),
		},
	}

	for name, data := range cases {
		out := data.from.VectorTo(data.to)
		require.Equalf(t, data.expected, out, "Case %s", name)
	}
}

func TestVectorEqualTo(t *testing.T) {
	cases := map[string]struct {
		v1      utils.Vector
		v2      utils.Vector
		isEqual bool
	}{
		"Identical Vectors": {
			v1: utils.Vector{
				XShift: 1,
				YShift: 1,
			},
			v2: utils.Vector{
				XShift: 1,
				YShift: 1,
			},
			isEqual: true,
		},
		"Simplified Vectors": {
			v1: utils.Vector{
				XShift: 2,
				YShift: 2,
			},
			v2: utils.Vector{
				XShift: 1,
				YShift: 1,
			},
			isEqual: true,
		},
		"Not Equal Vectors": {
			v1: utils.Vector{
				XShift: 2,
				YShift: 1,
			},
			v2: utils.Vector{
				XShift: 1,
				YShift: 1,
			},
			isEqual: false,
		},
		"Not Equal Because Sign": {
			v1: utils.Vector{
				XShift: -1,
				YShift: -1,
			},
			v2: utils.Vector{
				XShift: 1,
				YShift: 1,
			},
			isEqual: false,
		},
	}

	for name, data := range cases {
		require.Equalf(t, data.isEqual, data.v1.EqualTo(data.v2), "Case %s", name)
		require.Equalf(t, data.isEqual, data.v2.EqualTo(data.v1), "Case %s", name)
	}
}
