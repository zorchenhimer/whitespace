package whitespace

import (
	"fmt"

	inst "github.com/zorchenhimer/whitespace/instructions"
)

type node struct {
	idx int
	Instruction inst.Instruction
	Next *node
	Branch *node
}

func getAst(instructions []inst.Instruction) (*node, error) {
	labels := make(map[string]*node)   // destinations
	branches := make(map[string][]*node) // sources

	var ast *node
	var start *node

	if len(instructions) == 0 {
		return nil, fmt.Errorf("no instructions given")
	}

	// first pass.  Find labels, add all instructions.
	for idx, i := range instructions {
		a := &node{idx: idx, Instruction: i}

		switch i.Type() {
		case inst.CmdLabel:
			lbl := i.(*inst.Label)
			labels[lbl.Value] = a
		case inst.CmdCall, inst.CmdJump, inst.CmdJumpZero, inst.CmdJumpMinus:
			fc := i.(inst.FlowControl)
			branches[fc.Label()] = append(branches[fc.Label()], a)
		}

		if start == nil {
			start = a
		}

		if ast == nil {
			ast = a
		} else {
			ast.Next = a
			ast = a
		}
	}

	// set branch destinations
	for _, ln := range labels {
		lbl := ln.Instruction.(*inst.Label)
		if ln.Next == nil {
			return nil, fmt.Errorf("Label to nothing")
		}
		n := ln.Next

		if bl, ok := branches[lbl.Value]; ok {
			for _, b := range bl {
				b.Branch = n
			}
		}
	}

	// remove label instructions
	var curr *node
	curr = ast
	for {
		if curr == nil || curr.Next == nil {
			break
		}

		if curr.Next.Instruction.Type() == inst.CmdLabel {
			curr.Next = curr.Next.Next
		}
	}

	// Remove first node if it's a label
	if start.Instruction.Type() == inst.CmdLabel {
		start = start.Next
	}

	return start, nil

	//var prev *node
	//var curr *node
	//curr = ast
	//for {
	//	if curr.Instruction.Type() == inst.CmdLabel {
	//		// first node
	//		if prev == nil {
	//			ast = curr.Next
	//		}

	//		l := curr.(*inst.Label)
	//		if curr.Next == nil {
	//			return nil, fmt.Errorf("Label to nothing")
	//		}
	//		prev = curr
	//		//labels[l.Value] = curr.Next

	//		for v, b := range branches {
	//			if v == l.Value {
	//				b.Branch = curr.Next
	//			}
	//		}
	//	}

	//	if curr.Next == nil {
	//		break
	//	}
	//	curr = curr.Next
	//}

	// second pass, remove label instructions.  FIXME: breaks if first instruction is a label
	//for _, n := range ast {
	//	if n.Next == nil {
	//		break
	//	}

	//	if n.Next.Instruction.Type() == inst.CmdLabel {
	//		l := n.Next.(*inst.Label)
	//		if n.Next.Next == nil {
	//			return nil, fmt.Errorf("Label to nothing")
	//		}
	//		n.Next = n.Next.Next
	//		labels[l.Value] = n.Next

	//		for v, b := range branches {
	//			if v == l.Value {
	//				b.Branch = n.Next
	//			}
	//		}
	//	}
	//}
}
