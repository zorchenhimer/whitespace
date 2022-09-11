package whitespace

import (
	"testing"
	//"fmt"

	inst "github.com/zorchenhimer/whitespace/instructions"
)

type AstTestCase struct {
	Name string
	Input []inst.Instruction
	Output *node
}

func TestAst(t *testing.T) {
	tst := AstTestCase{
		"Call", []inst.Instruction{
			&inst.Push{1},
			&inst.Call{"S"},
			&inst.PrintNumber{},
			&inst.Stop{},
			&inst.Label{"S"},
			&inst.Push{2},
			&inst.Multiply{},
			&inst.Return{},
		}, nil,
	}

	ast, err := getAst(tst.Input)
	if err != nil {
		t.Fatalf("getAst() returned error: %s", err)
	}

	for ast != nil {
		cmd := ast.Instruction.Type()
		//fmt.Printf("[%d] %s", ast.idx, inst.CmdString(cmd))
		t.Logf("[%d] %s", ast.idx, inst.CmdString(cmd))
		switch ast.Instruction.(type) {
		case inst.FlowControl:
			fc := ast.Instruction.(inst.FlowControl)
			t.Logf(" %s", fc.Label())
		case inst.Push:
			i := ast.Instruction.(*inst.Push)
			t.Logf(" %d", i.Value)
		case inst.Copy:
			i := ast.Instruction.(*inst.Copy)
			t.Logf(" %d", i.Value)
		case inst.Slide:
			i := ast.Instruction.(*inst.Slide)
			t.Logf(" %d", i.Value)
		}

		//fmt.Println("")
		ast = ast.Next
	}
}
