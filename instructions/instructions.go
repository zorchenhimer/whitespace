package instructions

type Command int
const (
	/*
	Command character notation:
		S Space
		T Tab
		N Newline
		# Number
		L Label

		Both numbers and labels are terminated by a newline.
	*/

	CmdInvalid = Command(iota)

	// SS#
	CmdPush

	// Duplicate the top item on the stack
	// SNS
	CmdDuplicate

	// Copy the Nth item on the stack onto the top of the stack
	// STS#
	CmdCopy

	// Swap the top two items on the stack
	// SNT
	CmdSwap

	// Discard the top item in the stack
	// SNN
	CmdDiscard

	// Slide N items off the stack, keeping the top item
	// STN
	CmdSlide

	// TSSS
	CmdAdd

	// TSST
	CmdSubtract

	// TSSN
	CmdMultiply

	// TSTS
	CmdDivide

	// TSTT
	CmdModulo

	// Store in heap
	// TTS
	CmdStore

	// Load from heap
	// TTT
	CmdLoad

	// Define a label
	// NSSL
	CmdLabel

	// Call a subroutine, given a label
	// NSTL
	CmdCall

	// Jump to a label
	// NSNL
	CmdJump

	// Jump to a label if the top of the stack is Zero
	// NTS
	CmdJumpZero

	// Jump to a label if the top of the stack is Negitave
	// NTT
	CmdJumpMinus

	// Return from subroutine
	// NTN
	CmdReturn

	// Stop program
	// NNN
	CmdStop

	// Output the character at the top of the stack
	// TNSS
	CmdPrintChar

	// Output the Number at the top of the stack
	// TNST
	CmdPrintNumber

	// Read a character and place it in the location given by the top of the stack
	// TNTS
	CmdReadChar

	// Read a number and place it in the location given by the top of the stack
	// TNTT
	CmdReadNumber
)

type Instruction interface {
	Type() Command
}

type Push struct {
	Value int64
}

type Copy struct {
	Value int64
}

type Slide struct {
	Value int64
}

func (c Push)  Type() Command { return CmdPush }
func (c Copy)  Type() Command { return CmdCopy }
func (c Slide) Type() Command { return CmdSlide }

type Duplicate struct {}
type Swap struct {}
type Discard struct {}

func (c Duplicate) Type() Command { return CmdDuplicate }
func (c Swap)      Type() Command { return CmdSwap }
func (c Discard)   Type() Command { return CmdDiscard }

// Math

type Add struct {}
type Subtract struct {}
type Multiply struct {}
type Divide struct {}
type Modulo struct {}

func (c Add)      Type() Command { return CmdAdd }
func (c Subtract) Type() Command { return CmdSubtract }
func (c Multiply) Type() Command { return CmdMultiply }
func (c Divide)   Type() Command { return CmdDivide }
func (c Modulo)   Type() Command { return CmdModulo }

// Heap
type Store struct {}
type Load struct {}

func (c Store) Type() Command { return CmdStore }
func (c Load)  Type() Command { return CmdLoad }

// Flow control
type Label struct {
	Value string
}

type Call struct {
	Value string
}

type Jump struct {
	Value string
}

type JumpZero struct {
	Value string
}

type JumpMinus struct {
	Value string
}

type Return struct {}
type Stop struct {}

func (c Label)     Type() Command { return CmdLabel }
func (c Call)      Type() Command { return CmdCall }
func (c Jump)      Type() Command { return CmdJump }
func (c JumpZero)  Type() Command { return CmdJumpZero }
func (c JumpMinus) Type() Command { return CmdJumpMinus }
func (c Return)    Type() Command { return CmdReturn }
func (c Stop)      Type() Command { return CmdStop }

// I/O
type PrintChar struct {}
type PrintNumber struct {}
type ReadChar struct {}
type ReadNumber struct {}

func (c PrintChar)   Type() Command { return CmdPrintChar }
func (c PrintNumber) Type() Command { return CmdPrintNumber }
func (c ReadChar)    Type() Command { return CmdReadChar }
func (c ReadNumber)  Type() Command { return CmdReadNumber }
