package opcode

// Mode represents and addressing mode for opcode arguments
type Mode int

const (
	// ModePosition uses the value at the given address
	ModePosition Mode = iota
	// ModeImmediate uses the given value
	ModeImmediate
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
	// InstructionHalt represents the operation that halts the program
	InstructionHalt Instruction = 99
)

// InstructionParameterCount denotes how many arguments each agument accepts
var InstructionParameterCount = map[Instruction]int{
	InstructionAdd:      3,
	InstructionMultiply: 3,
	InstructionInput:    1,
	InstructionOutput:   1,
	InstructionHalt:     0,
}
