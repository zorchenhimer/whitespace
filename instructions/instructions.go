package instructions

import (
	"fmt"
)

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
	// NTSL
	CmdJumpZero

	// Jump to a label if the top of the stack is Negitave
	// NTTL
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
	Wsp() string
	Asm() string
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

func (c Push)  Wsp() string { return "  "+EncodeNumber(c.Value) }
func (c Copy)  Wsp() string { return " \t "+EncodeNumber(c.Value) }
func (c Slide) Wsp() string { return " \t\n"+EncodeNumber(c.Value)  }

func (c Push)  Asm() string { return fmt.Sprintf("push %d", c.Value) }
func (c Copy)  Asm() string { return fmt.Sprintf("copy %d", c.Value) }
func (c Slide) Asm() string { return fmt.Sprintf("slide %d", c.Value)  }

type Duplicate struct {}
type Swap struct {}
type Discard struct {}

func (c Duplicate) Type() Command { return CmdDuplicate }
func (c Swap)      Type() Command { return CmdSwap }
func (c Discard)   Type() Command { return CmdDiscard }

func (c Duplicate) Wsp() string { return " \n " }
func (c Swap)      Wsp() string { return " \n\t" }
func (c Discard)   Wsp() string { return " \n\n" }

func (c Duplicate) Asm() string { return "duplicate" }
func (c Swap)      Asm() string { return "swap" }
func (c Discard)   Asm() string { return "discard" }

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

func (c Add)      Wsp() string { return "\t   " }
func (c Subtract) Wsp() string { return "\t  \t" }
func (c Multiply) Wsp() string { return "\t  \n" }
func (c Divide)   Wsp() string { return "\t \t " }
func (c Modulo)   Wsp() string { return "\t \t\t" }

func (c Add)      Asm() string { return "add" }
func (c Subtract) Asm() string { return "subtract" }
func (c Multiply) Asm() string { return "multiply" }
func (c Divide)   Asm() string { return "divide" }
func (c Modulo)   Asm() string { return "modulo" }

// Heap
type Store struct {}
type Load struct {}

func (c Store) Type() Command { return CmdStore }
func (c Load)  Type() Command { return CmdLoad }

func (c Store) Wsp() string { return "\t\t " }
func (c Load)  Wsp() string { return "\t\t\t" }

func (c Store) Asm() string { return "store" }
func (c Load)  Asm() string { return "load" }

// Flow control
type FlowControl interface {
	Label() string
}

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

func (c Label)     Label() string     { return c.Value }
func (c Call)      Call() string      { return c.Value }
func (c Jump)      Jump() string      { return c.Value }
func (c JumpZero)  JumpZero() string  { return c.Value }
func (c JumpMinus) JumpMinus() string { return c.Value }

type Return struct {}
type Stop struct {}

//func (c Label)     Label() string { return c.Value }
func (c Call)      Label() string { return c.Value }
func (c Jump)      Label() string { return c.Value }
func (c JumpZero)  Label() string { return c.Value }
func (c JumpMinus) Label() string { return c.Value }

func (c Label)     Type() Command { return CmdLabel }
func (c Call)      Type() Command { return CmdCall }
func (c Jump)      Type() Command { return CmdJump }
func (c JumpZero)  Type() Command { return CmdJumpZero }
func (c JumpMinus) Type() Command { return CmdJumpMinus }
func (c Return)    Type() Command { return CmdReturn }
func (c Stop)      Type() Command { return CmdStop }

func (c Label)     Wsp() string { return "\n  "+c.Value }
func (c Call)      Wsp() string { return "\n \t"+c.Value }
func (c Jump)      Wsp() string { return "\n \n"+c.Value }
func (c JumpZero)  Wsp() string { return "\n\t "+c.Value }
func (c JumpMinus) Wsp() string { return "\n\t\t"+c.Value }
func (c Return)    Wsp() string { return "\n\t\n" }
func (c Stop)      Wsp() string { return "\n\n\n" }

func (c Label)     Asm() string { return "label "+DecodeLabel(c.Value) }
func (c Call)      Asm() string { return "call "+DecodeLabel(c.Value) }
func (c Jump)      Asm() string { return "jump "+DecodeLabel(c.Value) }
func (c JumpZero)  Asm() string { return "jumpzero "+DecodeLabel(c.Value) }
func (c JumpMinus) Asm() string { return "jumpminus "+DecodeLabel(c.Value) }
func (c Return)    Asm() string { return "return" }
func (c Stop)      Asm() string { return "stop" }

// I/O
type PrintChar struct {}
type PrintNumber struct {}
type ReadChar struct {}
type ReadNumber struct {}

func (c PrintChar)   Type() Command { return CmdPrintChar }
func (c PrintNumber) Type() Command { return CmdPrintNumber }
func (c ReadChar)    Type() Command { return CmdReadChar }
func (c ReadNumber)  Type() Command { return CmdReadNumber }

func (c PrintChar)   Wsp() string { return "\t\n  " }
func (c PrintNumber) Wsp() string { return "\t\n \t" }
func (c ReadChar)    Wsp() string { return "\t\n\t " }
func (c ReadNumber)  Wsp() string { return "\t\n\t\t" }

func (c PrintChar)   Asm() string { return "printchar" }
func (c PrintNumber) Asm() string { return "printnumber" }
func (c ReadChar)    Asm() string { return "readchar" }
func (c ReadNumber)  Asm() string { return "readnumber" }

func CmdString(c Command) string {
	switch c {
	case CmdInvalid:
		return "CmdInvalid"
	case CmdPush:
		return "CmdPush"
	case CmdDuplicate:
		return "CmdDuplicate"
	case CmdCopy:
		return "CmdCopy"
	case CmdSwap:
		return "CmdSwap"
	case CmdDiscard:
		return "CmdDiscard"
	case CmdSlide:
		return "CmdSlide"
	case CmdAdd:
		return "CmdAdd"
	case CmdSubtract:
		return "CmdSubtract"
	case CmdMultiply:
		return "CmdMultiply"
	case CmdDivide:
		return "CmdDivide"
	case CmdModulo:
		return "CmdModulo"
	case CmdStore:
		return "CmdStore"
	case CmdLoad:
		return "CmdLoad"
	case CmdLabel:
		return "CmdLabel"
	case CmdCall:
		return "CmdCall"
	case CmdJump:
		return "CmdJump"
	case CmdJumpZero:
		return "CmdJumpZero"
	case CmdJumpMinus:
		return "CmdJumpMinus"
	case CmdReturn:
		return "CmdReturn"
	case CmdStop:
		return "CmdStop"
	case CmdPrintChar:
		return "CmdPrintChar"
	case CmdPrintNumber:
		return "CmdPrintNumber"
	case CmdReadChar:
		return "CmdReadChar"
	case CmdReadNumber:
		return "CmdReadNumber"
	}

	return "Unknown Command"
}
