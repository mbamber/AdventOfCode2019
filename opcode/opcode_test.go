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
		ptr          int
		codes        []int
		mode         opcode.Mode
		relativeBase int
		expectedVal  int
	}{
		"Position": {
			ptr:          1,
			codes:        []int{1002, 4, 3, 4, 33},
			mode:         opcode.ModePosition,
			relativeBase: 0,
			expectedVal:  33,
		},
		"Immediate": {
			ptr:          2,
			codes:        []int{1002, 4, 3, 4, 33},
			mode:         opcode.ModeImmediate,
			relativeBase: 0,
			expectedVal:  3,
		},
		"Relative": {
			ptr:          1,
			codes:        []int{202, -7, 2, 0, 2},
			mode:         opcode.ModeRelative,
			relativeBase: 10,
			expectedVal:  0, // codes[codes[1] + 10] = codes[-7 + 10] = codes[3] = 0
		},
	}

	for name, data := range cases {
		val, err := opcode.GetArgumentValue(data.ptr, data.codes, data.mode, data.relativeBase)
		require.NoErrorf(t, err, "Case %s", name)
		require.Equalf(t, data.expectedVal, val, "Case %s", name)
	}
}

func TestProcessInstruction(t *testing.T) {
	cases := map[string]struct {
		ptr                  int
		codes                []int
		instruction          opcode.Instruction
		modes                []opcode.Mode
		input                int
		relativeBase         int
		expectedCodes        []int
		expectedOutputs      []int
		expectedNewPtr       int
		expectedRelativeBase int
	}{
		"Example 1": {
			ptr:                  0,
			codes:                []int{1002, 4, 3, 4, 33},
			instruction:          opcode.InstructionMultiply,
			modes:                []opcode.Mode{opcode.ModePosition, opcode.ModeImmediate, opcode.ModePosition},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{1002, 4, 3, 4, 99},
			expectedOutputs:      []int{},
			expectedNewPtr:       4,
			expectedRelativeBase: 0,
		},
		"Output": {
			ptr:                  2,
			codes:                []int{3, 0, 4, 0, 99},
			instruction:          opcode.InstructionOutput,
			modes:                []opcode.Mode{opcode.ModePosition},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{3, 0, 4, 0, 99},
			expectedOutputs:      []int{3},
			expectedNewPtr:       4,
			expectedRelativeBase: 0,
		},
		"Jump True": {
			ptr:                  0,
			codes:                []int{1105, 1, 10, 99},
			instruction:          opcode.InstructionJumpTrue,
			modes:                []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{1105, 1, 10, 99},
			expectedOutputs:      []int{},
			expectedNewPtr:       10,
			expectedRelativeBase: 0,
		},
		"Jump False": {
			ptr:                  0,
			codes:                []int{1106, 0, 10, 99},
			instruction:          opcode.InstructionJumpFalse,
			modes:                []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{1106, 0, 10, 99},
			expectedOutputs:      []int{},
			expectedNewPtr:       10,
			expectedRelativeBase: 0,
		},
		"Less Than": {
			ptr:                  0,
			codes:                []int{11107, 1, 2, 0},
			instruction:          opcode.InstructionLessThan,
			modes:                []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{1, 1, 2, 0},
			expectedOutputs:      []int{},
			expectedNewPtr:       4,
			expectedRelativeBase: 0,
		},
		"Not Less Than": {
			ptr:                  0,
			codes:                []int{11107, 2, 1, 0},
			instruction:          opcode.InstructionLessThan,
			modes:                []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{0, 2, 1, 0},
			expectedOutputs:      []int{},
			expectedNewPtr:       4,
			expectedRelativeBase: 0,
		},
		"Equals": {
			ptr:                  0,
			codes:                []int{11108, 1, 1, 0},
			instruction:          opcode.InstructionEquals,
			modes:                []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{1, 1, 1, 0},
			expectedOutputs:      []int{},
			expectedNewPtr:       4,
			expectedRelativeBase: 0,
		},
		"Not Equals": {
			ptr:                  0,
			codes:                []int{11108, 1, 2, 0},
			instruction:          opcode.InstructionEquals,
			modes:                []opcode.Mode{opcode.ModeImmediate, opcode.ModeImmediate, opcode.ModeImmediate},
			input:                1,
			relativeBase:         0,
			expectedCodes:        []int{0, 1, 2, 0},
			expectedOutputs:      []int{},
			expectedNewPtr:       4,
			expectedRelativeBase: 0,
		},
		"Relative Base Offset": {
			ptr:                  0,
			codes:                []int{109, 19},
			instruction:          opcode.InstructionRelativeBaseOffset,
			modes:                []opcode.Mode{opcode.ModeImmediate},
			input:                1,
			relativeBase:         2000,
			expectedCodes:        []int{109, 19},
			expectedOutputs:      []int{},
			expectedNewPtr:       2,
			expectedRelativeBase: 2019,
		},
	}

	for name, data := range cases {
		in := make(chan int)
		out := make(chan int, 20)
		done := make(chan bool, 1)

		err := opcode.ProcessInstruction(&data.ptr, &data.relativeBase, data.codes, data.instruction, data.modes, in, out, done)
		require.NoErrorf(t, err, "Case %s", name)

		outputs := []int{}
		isDone := false
		for !isDone {
			select {
			case isDone = <-done:
				continue
			case output := <-out:
				outputs = append(outputs, output)
			}
		}

		require.Equalf(t, data.expectedCodes, data.codes, "Case %s", name)
		require.Equalf(t, data.expectedNewPtr, data.ptr, "Case %s", name)
		require.Equalf(t, data.expectedRelativeBase, data.relativeBase, "Case %s", name)
		require.Equalf(t, data.expectedOutputs, outputs, "Case %s", name)
	}
}

func TestRun(t *testing.T) {
	cases := map[string]struct {
		codes    []int
		input    int
		expected []int
	}{
		"Position Mode Equal to 8": {
			codes:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    8,
			expected: []int{1},
		},
		"Position Mode Not Equal to 8": {
			codes:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    7,
			expected: []int{0},
		},
		"Position Mode Less Than 8": {
			codes:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    7,
			expected: []int{1},
		},
		"Position Mode Not Less Than 8": {
			codes:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    8,
			expected: []int{0},
		},
		"Immediate Mode Equal to 8": {
			codes:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    8,
			expected: []int{1},
		},
		"Immediate Mode Not Equal to 8": {
			codes:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    7,
			expected: []int{0},
		},
		"Immediate Mode Less Than 8": {
			codes:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    7,
			expected: []int{1},
		},
		"Immediate Mode Not Less Than 8": {
			codes:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    8,
			expected: []int{0},
		},
		"Position Mode Jump Zero": {
			codes:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    0,
			expected: []int{0},
		},
		"Position Mode Jump NonZero": {
			codes:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    99,
			expected: []int{1},
		},
		"Immediate Mode Jump Zero": {
			codes:    []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    0,
			expected: []int{0},
		},
		"Immediate Mode Jump NonZero": {
			codes:    []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    99,
			expected: []int{1},
		},
		"Less Than 8 Ouputs 999": {
			codes:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    7,
			expected: []int{999},
		},
		"Equal to 8 Ouputs 1000": {
			codes:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    8,
			expected: []int{1000},
		},
		"Greater Than 8 Ouputs 1001": {
			codes:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    9,
			expected: []int{1001},
		},
		"Quine": {
			codes: append([]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, func() []int {
				// Add a bunch of 0s to the end of the program
				s := []int{}
				for i := 0; i < 100; i++ {
					s = append(s, 0)
				}
				return s
			}()...),
			input:    0,
			expected: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
	}

	for name, data := range cases {
		in := make(chan int, 20)
		in <- data.input

		out := make(chan int, 20)

		err := opcode.Run(data.codes, in, out)
		require.NoErrorf(t, err, "Case %s", name)

		outputs := []int{}
		for output := range out {
			outputs = append(outputs, output)
		}
		require.Equalf(t, data.expected, outputs, "Case %s", name)
	}
}
