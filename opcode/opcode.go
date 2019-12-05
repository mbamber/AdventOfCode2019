package opcode

import (
	"fmt"
	"strconv"

	"aoc/utils"
)

// DetermineCodeInfo gets the code and all the paramter modes for that caode
func DetermineCodeInfo(n int) (code Instruction, paramModes []Mode, err error) {
	digits := utils.GetDigits(n)

	// Get the code
	var intcode int
	if len(digits) == 1 {
		intcode = digits[0]
	} else {
		intcode, err = strconv.Atoi(fmt.Sprintf("%d%d", digits[len(digits)-2], digits[len(digits)-1]))
		if err != nil {
			return 0, nil, err
		}
	}
	code = Instruction(intcode)

	paramModes = make([]Mode, InstructionParameterCount[code])
	for i := 0; i < len(paramModes); i++ {
		j := len(digits) - 3 - i
		if j >= 0 {
			paramModes[i] = Mode(digits[j])
		} else {
			paramModes[i] = ModePosition
		}
	}

	return code, paramModes, nil
}

// GetArgumentValue returns the argument value for the current pointer
func GetArgumentValue(ptr int, codes []int, mode Mode) (int, error) {
	switch mode {
	case ModePosition:
		return codes[codes[ptr]], nil
	case ModeImmediate:
		return codes[ptr], nil
	default:
		return 0, fmt.Errorf("Invalid mode %d", mode)
	}
}

// ProcessInstruction processes a single instruction
func ProcessInstruction(instructionPointer int, codes []int, instruction Instruction, modes []Mode) ([]int, *int, error) {
	switch instruction {
	case InstructionAdd:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, nil, err
		}
		arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
		if err != nil {
			return nil, nil, err
		}

		codes[codes[instructionPointer+3]] = arg1 + arg2
		return codes, nil, nil

	case InstructionMultiply:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, nil, err
		}
		arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
		if err != nil {
			return nil, nil, err
		}

		codes[codes[instructionPointer+3]] = arg1 * arg2
		return codes, nil, nil

	case InstructionInput:
		codes[codes[instructionPointer+1]] = 1 // Just supply 1 as an input for now
		return codes, nil, nil

	case InstructionOutput:
		arg, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, nil, err
		}
		fmt.Println("[OUTPUT] ", arg)
		return codes, &arg, nil
	case InstructionHalt:
		return codes, nil, nil
	default:
		return nil, nil, fmt.Errorf("Unknown instruction %d", instruction)
	}
}

// Run runs an opcode program
func Run(codes []int) (int, error) {
	ip := 0
	var output *int
	lastOutput := 0
	for {
		instruction, modes, err := DetermineCodeInfo(codes[ip])
		if err != nil {
			return 0, err
		}

		if instruction == InstructionHalt {
			return lastOutput, nil
		}

		codes, output, err = ProcessInstruction(ip, codes, instruction, modes)
		if err != nil {
			return 0, err
		}
		if output != nil {
			lastOutput = *output
		}

		// Increment
		ip += len(modes) + 1
	}
}
