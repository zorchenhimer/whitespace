package whitespace

import (
	"io"
	"fmt"

	inst "github.com/zorchenhimer/whitespace/instructions"
)

type Engine struct {
	ast *node
	stack *Stack[int64]
	calls *Stack[*node]
}

func NewEngine(reader io.Reader) (*Engine, error) {
	p := NewParser(NewReader(reader))
	instlst, err := p.Parse()
	if err != nil {
		return nil, err
	}

	ast, err := getAst(instlst)
	if err != nil {
		return nil, err
	}

	return &Engine{ast: ast, stack: NewStack[int64](), calls: NewStack[*node]()}, nil
}

func (e *Engine) Run(input io.Reader, output io.Writer) error {
	if e.ast.Instruction == nil {
		return fmt.Errorf("nil start node")
	}

	for {
		i := e.ast.Instruction
		branched := false
		switch i.Type() {
		case inst.CmdPush:
			c := i.(*inst.Push)
			e.stack.Push(c.Value)

		case inst.CmdDuplicate:
			v := e.stack.Pop()
			e.stack.Push(v)
			e.stack.Push(v)

		case inst.CmdCopy:
			c := i.(*inst.Copy)
			v := e.stack.Get(c.Value)
			e.stack.Push(v)

		case inst.CmdSwap:
			a := e.stack.Pop()
			b := e.stack.Pop()
			e.stack.Push(a)
			e.stack.Push(b)

		case inst.CmdDiscard:
			e.stack.Pop()

		case inst.CmdSlide:
			t := e.stack.Pop()
			c := i.(*inst.Slide)
			for x := int64(0); x < c.Value; x++ {
				e.stack.Pop()
			}
			e.stack.Push(t)

		case inst.CmdAdd:
			a := e.stack.Pop()
			b := e.stack.Pop()
			e.stack.Push(a+b)

		case inst.CmdSubtract:
			a := e.stack.Pop()
			b := e.stack.Pop()
			e.stack.Push(a-b)

		case inst.CmdMultiply:
			a := e.stack.Pop()
			b := e.stack.Pop()
			e.stack.Push(a*b)

		case inst.CmdDivide:
			a := e.stack.Pop()
			b := e.stack.Pop()
			e.stack.Push(a/b)

		case inst.CmdModulo:
			a := e.stack.Pop()
			b := e.stack.Pop()
			e.stack.Push(a%b)

		case inst.CmdStore, inst.CmdLoad:
			return fmt.Errorf("Store & Load unimplemented")

		case inst.CmdLabel:
			// do nothing

		case inst.CmdCall:
			e.calls.Push(e.ast)

			e.ast = e.ast.Branch
			branched = true

		case inst.CmdJump:
			e.ast = e.ast.Branch
			branched = true

		case inst.CmdJumpZero:
			v := e.stack.Get(0)
			if v == 0 {
				e.ast = e.ast.Branch
				branched = true
			}

		case inst.CmdJumpMinus:
			v := e.stack.Get(0)
			if v < 0 {
				e.ast = e.ast.Branch
				branched = true
			}

		case inst.CmdReturn:
			e.ast = e.calls.Pop()
			branched = true

		case inst.CmdStop:
			return nil

		case inst.CmdPrintChar:
			if output == nil {
				return fmt.Errorf("attempt to print to nil")
			}
			fmt.Fprintf(output, "%c", e.stack.Pop())

		case inst.CmdPrintNumber:
			if output == nil {
				return fmt.Errorf("attempt to print to nil")
			}
			fmt.Fprintf(output, "%d", e.stack.Pop())

		case inst.CmdReadChar:
			if input == nil {
				return fmt.Errorf("attempt to read from nil")
			}

			var c int64
			_, err := fmt.Fscanf(input, "%c", &c)
			if err != nil {
				return fmt.Errorf("char read error: %w", err)
			}
			e.stack.Push(c)

		case inst.CmdReadNumber:
			if input == nil {
				return fmt.Errorf("attempt to read from nil")
			}

			var c int64
			_, err := fmt.Fscanf(input, "%d", &c)
			if err != nil {
				return fmt.Errorf("char read error: %w", err)
			}
			e.stack.Push(c)
		}

		if !branched {
			e.ast = e.ast.Next
		}
	}
	return nil
}

//type InteractiveReader struct {
//	// nothin
//}
//
//func (ir *InteractiveReader) Read(p []byte) (int, error) {
//	return 0, nil
//}

