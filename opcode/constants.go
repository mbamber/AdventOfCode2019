package opcode

// Mode represents and addressing mode for opcode arguments
type Mode int

const (
	// ModePosition uses the value at the given address
	ModePosition Mode = 0
	// ModeImmediate uses the given value
	ModeImmediate Mode = 1
	// ModeRelative uses the value relative to the relative base
	ModeRelative Mode = 2
)

// Instruction represents an opcode instruction
type Instruction int

const (
	// InstructionAdd represents addition
	InstructionAdd Instruction = 1
	// InstructionMultiply represents multiplication
	InstructionMultiply Instruction = 2
	// InstructionInput represents an instruction that asks for input
	InstructionInput Instruction = 3
	// InstructionOutput represents an instruction that gives an output
	InstructionOutput Instruction = 4
	// InstructionJumpTrue represents the jump-if-true instruction
	InstructionJumpTrue Instruction = 5
	// InstructionJumpFalse represents the jump-if-false instruction
	InstructionJumpFalse Instruction = 6
	// InstructionLessThan represents the less than instruction
	InstructionLessThan Instruction = 7
	// InstructionEquals represents the equal to instruction
	InstructionEquals Instruction = 8
	// InstructionRelativeBaseOffset represents the relative base offset instruction
	InstructionRelativeBaseOffset Instruction = 9
	// InstructionHalt represents the operation that halts the program
	InstructionHalt Instruction = 99
)

// InstructionParameterCount denotes how many arguments each agument accepts
var InstructionParameterCount = map[Instruction]int{
	InstructionAdd:                3,
	InstructionMultiply:           3,
	InstructionInput:              1,
	InstructionOutput:             1,
	InstructionJumpTrue:           2,
	InstructionJumpFalse:          2,
	InstructionLessThan:           3,
	InstructionEquals:             3,
	InstructionRelativeBaseOffset: 1,
	InstructionHalt:               0,
}
