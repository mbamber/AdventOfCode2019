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
func ProcessInstruction(instructionPointer int, codes []int, instruction Instruction, modes []Mode, in, out chan int, done chan bool) (newCodes []int, newInstructionPointer int, err error) {
	// When the instruction has been processed, send `true` to the `done` chan so the caller knows we're done
	defer func() {
		done <- true
	}()

	switch instruction {
	case InstructionAdd:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
		if err != nil {
			return nil, 0, err
		}

		codes[codes[instructionPointer+3]] = arg1 + arg2
		return codes, 0, nil

	case InstructionMultiply:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
		if err != nil {
			return nil, 0, err
		}

		codes[codes[instructionPointer+3]] = arg1 * arg2
		return codes, 0, nil

	case InstructionInput:
		codes[codes[instructionPointer+1]] = <-in
		return codes, 0, nil

	case InstructionOutput:
		arg, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		out <- arg
		return codes, 0, nil
	case InstructionJumpTrue:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		if arg1 != 0 {
			arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
			if err != nil {
				return nil, 0, err
			}
			return codes, arg2, nil
		}
		return codes, 0, nil
	case InstructionJumpFalse:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		if arg1 == 0 {
			arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
			if err != nil {
				return nil, 0, err
			}
			return codes, arg2, nil
		}
		return codes, 0, nil
	case InstructionLessThan:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
		if err != nil {
			return nil, 0, err
		}

		if arg1 < arg2 {
			codes[codes[instructionPointer+3]] = 1
		} else {
			codes[codes[instructionPointer+3]] = 0
		}
		return codes, 0, nil
	case InstructionEquals:
		arg1, err := GetArgumentValue(instructionPointer+1, codes, modes[0])
		if err != nil {
			return nil, 0, err
		}
		arg2, err := GetArgumentValue(instructionPointer+2, codes, modes[1])
		if err != nil {
			return nil, 0, err
		}

		if arg1 == arg2 {
			codes[codes[instructionPointer+3]] = 1
		} else {
			codes[codes[instructionPointer+3]] = 0
		}
		return codes, 0, nil
	case InstructionHalt:
		return codes, 0, nil
	default:
		return nil, 0, fmt.Errorf("Unknown instruction %d", instruction)
	}
}

// Run runs an opcode program
func Run(codes []int, in, out chan int) error {
	ip := 0
	var newPtr int

	done := make(chan bool)
	for {
		instruction, modes, err := DetermineCodeInfo(codes[ip])
		if err != nil {
			close(out)
			return err
		}

		if instruction == InstructionHalt {
			close(out)
			return nil
		}

		go func() {
			codes, newPtr, err = ProcessInstruction(ip, codes, instruction, modes, in, out, done)
		}()

		// Block until we're done
		<-done

		// Check for any error
		if err != nil {
			close(out)
			return err
		}

		// Increment
		if newPtr == 0 {
			ip += len(modes) + 1
		} else {
			ip = newPtr
		}
	}
}
