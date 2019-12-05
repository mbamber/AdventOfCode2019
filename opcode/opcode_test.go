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
			expectedInstruction: opcode.InstructionAdd,
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
		expectedCodes  []int
		expectedOutput *int
	}{
		"Example 1": {
			ptr:            0,
			codes:          []int{1002, 4, 3, 4, 33},
			instruction:    opcode.InstructionMultiply,
			modes:          []opcode.Mode{opcode.ModePosition, opcode.ModeImmediate, opcode.ModePosition},
			expectedCodes:  []int{1002, 4, 3, 4, 99},
			expectedOutput: nil,
		},
		"Output": {
			ptr:            2,
			codes:          []int{3, 0, 4, 0, 99},
			instruction:    opcode.InstructionOutput,
			modes:          []opcode.Mode{opcode.ModePosition},
			expectedCodes:  []int{3, 0, 4, 0, 99},
			expectedOutput: &three,
		},
	}

	for name, data := range cases {
		codes, out, err := opcode.ProcessInstruction(data.ptr, data.codes, data.instruction, data.modes)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equalf(t, data.expectedCodes, codes, "Case %s", name)
		require.Equalf(t, data.expectedOutput, out, "Case %s", name)
	}
}
