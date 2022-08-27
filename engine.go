package whitespace

import (
	"io"
	"fmt"

	inst "github.com/zorchenhimer/whitespace/instructions"
)

type Engine struct {
	commands []inst.Instruction
	labels map[string]int
	stack []int64
}

func NewEngine(reader io.Reader) (*Engine, error) {
	p := NewParser(NewReader(reader))
	inst, err := p.Parse()
	if err != nil {
		return nil, err
	}
	_ = inst

	return nil, fmt.Errorf("no, lmao")
	//return &Engine{commands: cmds, labels: lbls, stack: []int64{}}, nil
}

func (e *Engine) Run(input io.Reader, output io.Writer) error {
	return nil
}

//type InteractiveReader struct {
//	// nothin
//}
//
//func (ir *InteractiveReader) Read(p []byte) (int, error) {
//	return 0, nil
//}

