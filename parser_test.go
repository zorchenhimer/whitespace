package whitespace

import (
	"testing"
	"strings"

	inst "github.com/zorchenhimer/whitespace/instructions"
)

type TestCase struct {
	Name string
	Input string
	Output []inst.Instruction
}

func TestParseSingle(t *testing.T) {
	tests := []TestCase{
		{"Add",      "\t   ",   []inst.Instruction{inst.Add{}}},
		{"Subtract", "\t  \t",  []inst.Instruction{inst.Subtract{}}},
		{"Multiply", "\t  \n",  []inst.Instruction{inst.Multiply{}}},
		{"Divide",   "\t \t ",  []inst.Instruction{inst.Divide{}}},
		{"Modulo",   "\t \t\t", []inst.Instruction{inst.Modulo{}}},

		{"Push 0",   "   \n",             []inst.Instruction{inst.Push{0}}},
		{"Push 1",   "   \t\n",           []inst.Instruction{inst.Push{1}}},
		{"Push -75", "  \t\t  \t \t\t\n", []inst.Instruction{inst.Push{-75}}},

		{"Copy 1",   " \t  \t\n",          []inst.Instruction{inst.Copy{1}}},
		{"Copy -75", " \t  \t  \t \t\t\n", []inst.Instruction{inst.Copy{-75}}},

		{"Slide 1", " \t\n \t\n",            []inst.Instruction{inst.Slide{1}}},
		{"Slide 75", " \t\n \t  \t \t\t\n",  []inst.Instruction{inst.Slide{75}}},

		{"Discard",   " \n\n", []inst.Instruction{inst.Discard{}}},
		{"Duplicate", " \n ",  []inst.Instruction{inst.Duplicate{}}},
		{"Swap",      " \n\t", []inst.Instruction{inst.Swap{}}},

		{"Store", "\t\t ",  []inst.Instruction{inst.Store{}}},
		{"Load",  "\t\t\t", []inst.Instruction{inst.Load{}}},

		{"Label SSS", "\n      \n",   []inst.Instruction{inst.Label{"   "}}},
		{"Label TST", "\n   \t \t\n", []inst.Instruction{inst.Label{"\t \t"}}},
		{"Call SSS",  "\n \t   \n",    []inst.Instruction{inst.Call{"   "}}},
		{"Call TST",  "\n \t\t \t\n",  []inst.Instruction{inst.Call{"\t \t"}}},

		{"Jump SSS",      "\n \n   \n",     []inst.Instruction{inst.Jump{"   "}}},
		{"Jump TST",      "\n \n\t \t\n",   []inst.Instruction{inst.Jump{"\t \t"}}},
		{"JumpZero SSS",  "\n\t    \n",     []inst.Instruction{inst.JumpZero{"   "}}},
		{"JumpZero TST",  "\n\t \t \t\n",   []inst.Instruction{inst.JumpZero{"\t \t"}}},
		{"JumpMinus SSS", "\n\t\t   \n",    []inst.Instruction{inst.JumpMinus{"   "}}},
		{"JumpMinus TST", "\n\t\t\t \t\n",  []inst.Instruction{inst.JumpMinus{"\t \t"}}},

		{"Return", "\n\t\n", []inst.Instruction{inst.Return{}}},
		{"Stop",   "\n\n\n", []inst.Instruction{inst.Stop{}}},

		{"Print Char",   "\t\n  ",   []inst.Instruction{inst.PrintChar{}}},
		{"Print Number", "\t\n \t",  []inst.Instruction{inst.PrintNumber{}}},
		{"Read Char",    "\t\n\t ",  []inst.Instruction{inst.ReadChar{}}},
		{"Read Number",  "\t\n\t\t", []inst.Instruction{inst.ReadNumber{}}},
	}

	runParseTests(t, tests)
}

func TestParseNumber(t *testing.T) {
	tests := []struct{
		Name string
		Input string
		Output int64
	}{
		{"0",   " \n", 0},
		{"75",  " \t  \t \t\t\n", 75},
		{"-75", "\t\t  \t \t\t\n", -75},
		{"50",  " \t\t  \t \n", 50},
		{"-50", "\t\t\t  \t \n", -50},
	}

	for _, tst := range tests {
		t.Logf("%s: %q", tst.Name, tst.Input)
		p := NewParser(NewReader(strings.NewReader(tst.Input)))
		n, err := p.parseNumber()
		if err != nil {
			t.Logf("Parse error: %s", err)
			t.Fail()
			continue
		}

		if n != tst.Output {
			t.Logf("Value mismatch: %d vs %d", n, tst.Output)
			t.Fail()
		}
	}
}

func runParseTests(t *testing.T, tests []TestCase) {
	t.Helper()

	for _, tst := range tests {
		t.Log(tst.Name)
		p := NewParser(NewReader(strings.NewReader(tst.Input)))
		lst, err := p.Parse()
		if err != nil {
			t.Logf("Parse failed: %s", err)
			t.Fail()
			continue
		}

		if !instEqual(t, lst, tst.Output) {
			t.Logf("Unexpected output\n Rec: %v\n Exp: %v", lst, tst.Output)
			t.Fail()
		}
	}
}

func instEqual(t *testing.T, a, b []inst.Instruction) bool {
	t.Helper()

	if len(a) != len(b) {
		return false
	}
	t.Logf("len: %d", len(a))

	for i := 0; i < len(a); i++ {
		if a[i].Type() != b[i].Type() {
			t.Logf("%v != %v", a[i].Type(), b[i].Type())
			return false
		}
		t.Logf("%v == %v", a[i].Type(), b[i].Type())
	}

	t.Logf("equal\n%v\n%v", a, b)
	return true
}
