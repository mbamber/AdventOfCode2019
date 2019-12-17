package opcode

import (
	"fmt"
	"strconv"

	"aoc/utils"
)

// DetermineCodeInfo gets the code and all the paramter modes for that node
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
func GetArgumentValue(ptr int, codes []int, mode Mode, relativeBase int) (int, error) {
	switch mode {
	case ModePosition:
		return codes[codes[ptr]], nil
	case ModeImmediate:
		return codes[ptr], nil
	case ModeRelative:
		return codes[codes[ptr]+relativeBase], nil
	default:
		return 0, fmt.Errorf("Invalid mode %d", mode)
	}
}

// GetDestinationLocation returns a destination address in which to store a value
func GetDestinationLocation(ptr int, codes []int, mode Mode, relativeBase int) (int, error) {
	switch mode {
	case ModePosition:
		return codes[ptr], nil
	case ModeRelative:
		return codes[ptr] + relativeBase, nil
	default:
		return 0, fmt.Errorf("Invalid mode %d", mode)
	}
}

// ProcessInstruction processes a single instruction
func ProcessInstruction(instructionPointer, relativeBase *int, codes []int, instruction Instruction, modes []Mode, in, out chan int, done chan bool) error {

	// When the instruction has been processed, send `true` to the `done` chan so the caller knows we're done
	defer func() {
		*instructionPointer += (InstructionParameterCount[instruction] + 1)
		done <- true
	}()

	switch instruction {
	case InstructionAdd:
		arg1, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		arg2, err := GetArgumentValue(*instructionPointer+2, codes, modes[1], *relativeBase)
		if err != nil {
			return err
		}

		dst, err := GetDestinationLocation(*instructionPointer+3, codes, modes[2], *relativeBase)
		if err != nil {
			return err
		}

		codes[dst] = arg1 + arg2
		return nil

	case InstructionMultiply:
		arg1, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		arg2, err := GetArgumentValue(*instructionPointer+2, codes, modes[1], *relativeBase)
		if err != nil {
			return err
		}

		dst, err := GetDestinationLocation(*instructionPointer+3, codes, modes[2], *relativeBase)
		if err != nil {
			return err
		}

		codes[dst] = arg1 * arg2
		return nil

	case InstructionInput:
		dst, err := GetDestinationLocation(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		codes[dst] = <-in
		return nil

	case InstructionOutput:
		arg, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		out <- arg
		return nil
	case InstructionJumpTrue:
		arg1, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		if arg1 != 0 {
			arg2, err := GetArgumentValue(*instructionPointer+2, codes, modes[1], *relativeBase)
			if err != nil {
				return err
			}

			*instructionPointer = arg2
			*instructionPointer -= (InstructionParameterCount[instruction] + 1) // because we automatically add this as part of the defer
		}
		return nil
	case InstructionJumpFalse:
		arg1, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		if arg1 == 0 {
			arg2, err := GetArgumentValue(*instructionPointer+2, codes, modes[1], *relativeBase)
			if err != nil {
				return err
			}

			*instructionPointer = arg2
			*instructionPointer -= (InstructionParameterCount[instruction] + 1) // because we automatically add this as part of the defer
		}
		return nil
	case InstructionLessThan:
		arg1, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		arg2, err := GetArgumentValue(*instructionPointer+2, codes, modes[1], *relativeBase)
		if err != nil {
			return err
		}

		dst, err := GetDestinationLocation(*instructionPointer+3, codes, modes[2], *relativeBase)
		if err != nil {
			return err
		}

		if arg1 < arg2 {
			codes[dst] = 1
		} else {
			codes[dst] = 0
		}
		return nil
	case InstructionEquals:
		arg1, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}
		arg2, err := GetArgumentValue(*instructionPointer+2, codes, modes[1], *relativeBase)
		if err != nil {
			return err
		}

		dst, err := GetDestinationLocation(*instructionPointer+3, codes, modes[2], *relativeBase)
		if err != nil {
			return err
		}

		if arg1 == arg2 {
			codes[dst] = 1
			// codes[codes[*instructionPointer+3]] = 1
		} else {
			codes[dst] = 0
			// codes[codes[*instructionPointer+3]] = 0
		}
		return nil
	case InstructionRelativeBaseOffset:
		arg, err := GetArgumentValue(*instructionPointer+1, codes, modes[0], *relativeBase)
		if err != nil {
			return err
		}

		*relativeBase += arg
		return nil
	case InstructionHalt:
		close(out)
		return nil
	default:
		return fmt.Errorf("Unknown instruction %d", instruction)
	}
}

// Run runs an opcode program
func Run(codes []int, in, out chan int) error {
	ip, relativeBase := 0, 0

	for {
		done := make(chan bool, 1)
		instruction, modes, err := DetermineCodeInfo(codes[ip])
		if err != nil {
			close(out)
			return err
		}

		if instruction == InstructionHalt {
			close(out)
			return nil
		}

		err = ProcessInstruction(&ip, &relativeBase, codes, instruction, modes, in, out, done)
		// Check for any error
		if err != nil {
			close(out)
			return err
		}

		// Block until we're done
		<-done
	}
}
