package opcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/opcode"
)

func TestDetermineCodeInfo(t *testing.T) {
	cases := map[string]struct {
		input               int
		expectedInstruction opcode.Instruction
		expectedModes       []opcode.Mode
	}{
		"Example 1": {
			input:               1002,
			expectedInstruction: opcode.InstructionMultiply,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModeImmediate, opcode.ModePosition},
		},
		"Single Digit": {
			input:               2,
			expectedInstruction: opcode.InstructionMultiply,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModePosition, opcode.ModePosition},
		},
		"Double Digit": {
			input:               02,
			expectedInstruction: opcode.InstructionMultiply,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModePosition, opcode.ModePosition},
		},
		"Halt": {
			input:               99,
			expectedInstruction: opcode.InstructionHalt,
			expectedModes:       []opcode.Mode{},
		},
		"Input": {
			input:               3,
			expectedInstruction: opcode.InstructionInput,
			expectedModes:       []opcode.Mode{opcode.ModePosition},
		},
		"Output": {
			input:               104,
			expectedInstruction: opcode.InstructionOutput,
			expectedModes:       []opcode.Mode{opcode.ModeImmediate},
		},
		"Jump True": {
			input:               5,
			expectedInstruction: opcode.InstructionJumpTrue,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModePosition},
		},
		"Jump False": {
			input:               6,
			expectedInstruction: opcode.InstructionJumpFalse,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModePosition},
		},
		"Less than": {
			input:               7,
			expectedInstruction: opcode.InstructionLessThan,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModePosition, opcode.ModePosition},
		},
		"Equals": {
			input:               8,
			expectedInstruction: opcode.InstructionEquals,
			expectedModes:       []opcode.Mode{opcode.ModePosition, opcode.ModePosition, opcode.ModePosition},
		},
	}

	for name, data := range cases {
		instruction, modes, err := opcode.DetermineCodeInfo(data.input)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equalf(t, data.expectedInstruction, instruction, "Case %s", name)
		require.Equalf(t, data.expectedModes, modes, "Case %s", name)
	}
}

func TestGetArgumentValue(t *testing.T) {
	cases := map[string]struct {
		ptr         int
		codes       []int
		mode        opcode.Mode
		expectedVal int
	}{
		"Position": {
			ptr:         1,
			codes:       []int{1002, 4, 3, 4, 33},
			mode:        opcode.ModePosition,
			expectedVal: 33,
		},
		"Immediate": {
			ptr:         2,
			codes:       []int{1002, 4, 3, 4, 33},
			mode:        opcode.ModeImmediate,
			expectedVal: 3,
		},
	}

	for name, data := range cases {
		val, err := opcode.GetArgumentValue(data.ptr, data.codes, data.mode)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equalf(t, data.expectedVal, val, "Case %s", name)
	}
}

func TestProcessInstruction(t *testing.T) {
	three := 3

	cases := map[string]struct {
		ptr            int
		codes          []int
		instruction    opcode.Instruction
		modes          []opcode.Mode
		input          int
		expectedCodes  []int
		expectedOutput *int
		expectedNewPtr int
	}{
		"Example 1": {
			ptr:            0,
			codes:          []int{1002, 4, 3, 4, 33},
			instruction:    opcode.InstructionMultiply,
			modes:          []opcode.Mode{opcode.ModePosition, opcode.ModeImmediate, opcode.ModePosition},
			input:          1,
			expectedCodes:  []int{1002, 4, 3, 4, 99},
			expectedOutput: nil,
			expectedNewPtr: 0,
		},
		"Output": {
			ptr:            2,
			codes:          []int{3, 0, 4, 0, 99},
			instruction:    opcode.InstructionOutput,
			modes:          []opcode.Mode{opcode.ModePosition},
			input:          1,
			expectedCodes:  []int{3, 0, 4, 0, 99},
			expectedOutput: &three,
			expectedNewPtr: 0,
		},
		"Jump True": {
			ptr:            0,
			codes:          []int{1105, 0, 1, 99},
			instruction:    opcode.InstructionJumpTrue,
			modes:          []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate},
			input:          1,
			expectedCodes:  []int{1105, 0, 1, 99},
			expectedOutput: nil,
			expectedNewPtr: 1,
		},
		"Jump False": {
			ptr:            0,
			codes:          []int{1106, 1, 1, 99},
			instruction:    opcode.InstructionJumpFalse,
			modes:          []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate},
			input:          1,
			expectedCodes:  []int{1106, 1, 1, 99},
			expectedOutput: nil,
			expectedNewPtr: 1,
		},
		"Less Than": {
			ptr:            0,
			codes:          []int{11107, 1, 2, 0},
			instruction:    opcode.InstructionLessThan,
			modes:          []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:          1,
			expectedCodes:  []int{1, 1, 2, 0},
			expectedOutput: nil,
			expectedNewPtr: 0,
		},
		"Not Less Than": {
			ptr:            0,
			codes:          []int{11107, 2, 1, 0},
			instruction:    opcode.InstructionLessThan,
			modes:          []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:          1,
			expectedCodes:  []int{0, 2, 1, 0},
			expectedOutput: nil,
			expectedNewPtr: 0,
		},
		"Equals": {
			ptr:            0,
			codes:          []int{11108, 1, 1, 0},
			instruction:    opcode.InstructionEquals,
			modes:          []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:          1,
			expectedCodes:  []int{1, 1, 1, 0},
			expectedOutput: nil,
			expectedNewPtr: 0,
		},
		"Not Equals": {
			ptr:            0,
			codes:          []int{11108, 1, 2, 0},
			instruction:    opcode.InstructionEquals,
			modes:          []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:          1,
			expectedCodes:  []int{0, 1, 2, 0},
			expectedOutput: nil,
			expectedNewPtr: 0,
		},
	}

	for name, data := range cases {
		codes, out, ptr, err := opcode.ProcessInstruction(data.ptr, data.codes, data.instruction, data.modes, data.input)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equalf(t, data.expectedCodes, codes, "Case %s", name)
		require.Equalf(t, data.expectedOutput, out, "Case %s", name)
		require.Equalf(t, data.expectedNewPtr, ptr, "Case %s", name)
	}
}

func TestRun(t *testing.T) {
	cases := map[string]struct {
		codes    []int
		input    int
		expected int
	}{
		"Position Mode Equal to 8": {
			codes:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    8,
			expected: 1,
		},
		"Position Mode Not Equal to 8": {
			codes:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    7,
			expected: 0,
		},
		"Position Mode Less Than 8": {
			codes:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    7,
			expected: 1,
		},
		"Position Mode Not Less Than 8": {
			codes:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    8,
			expected: 0,
		},
		"Immediate Mode Equal to 8": {
			codes:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    8,
			expected: 1,
		},
		"Immediate Mode Not Equal to 8": {
			codes:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    7,
			expected: 0,
		},
		"Immediate Mode Less Than 8": {
			codes:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    7,
			expected: 1,
		},
		"Immediate Mode Not Less Than 8": {
			codes:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    8,
			expected: 0,
		},
		"Position Mode Jump Zero": {
			codes:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    0,
			expected: 0,
		},
		"Position Mode Jump NonZero": {
			codes:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    99,
			expected: 1,
		},
		"Immediate Mode Jump Zero": {
			codes:    []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    0,
			expected: 0,
		},
		"Immediate Mode Jump NonZero": {
			codes:    []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    99,
			expected: 1,
		},
		"Less Than 8 Ouputs 999": {
			codes:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    7,
			expected: 999,
		},
		"Equal to 8 Ouputs 1000": {
			codes:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    8,
			expected: 1000,
		},
		"Greater Than 8 Ouputs 1001": {
			codes:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    9,
			expected: 1001,
		},
	}

	for name, data := range cases {
		out, err := opcode.Run(data.codes, data.input)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equalf(t, data.expected, out, "Case %s", name)
	}
}
