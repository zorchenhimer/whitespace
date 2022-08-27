package whitespace

import (
	"fmt"
	"strings"
	"io"

	inst "github.com/zorchenhimer/whitespace/instructions"
)

type Parser struct {
	r *Reader
	labels map[string]int
}

func NewParser(reader *Reader) *Parser {
	return &Parser{r: reader, labels: make(map[string]int)}
}

func (p *Parser) Parse() ([]inst.Instruction, error) {
	var err error
	var r rune
	cmds := []inst.Instruction{}
	for {
		r, _, err = p.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var cmd inst.Instruction
		switch r {
			case ' ':
				// stack
				cmd, err = p.parseStack()
			case '\n':
				// flow control
				cmd, err = p.parseFlow()
				if cmd.Type() == inst.CmdLabel{
					lbl := cmd.(*inst.Label)
					if _, exist := p.labels[lbl.Value]; exist {
						return nil, fmt.Errorf("Duplicate label %q", lbl.Value)
					}
					p.labels[lbl.Value] = len(cmds)
				}
			case '\t':
				// parse next rune to complete IMP
				r, _, err = p.r.ReadRune()
				if err != nil {
					return nil, fmt.Errorf("Broken IMP")
				}

				switch r {
				case ' ':
					// math
					cmd, err = p.parseMath()
				case '\t':
					// heap
					cmd, err = p.parseHeap()
				case '\n':
					// I/O
					cmd, err = p.parseIO()
				default:
					cmd = nil
					err = fmt.Errorf("Bad tab IMP")
				}

			default:
				cmd = nil
				err = fmt.Errorf("Bad IMP")
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		cmds = append(cmds, cmd)
	}

	return cmds, nil
}

func (p *Parser) parseIO() (inst.Instruction, error) {
	runes := []rune{}
	for i := 0; i < 2; i++ {
		r, n, err := p.r.ReadRune()
		if err != nil || n == 0 {
			return nil, fmt.Errorf("Bad io parse")
		}

		runes = append(runes, r)
	}

	if len(runes) != 2 {
		return nil, fmt.Errorf("Bad io parse length")
	}

	switch runes[0] {
	case ' ':
		switch runes[1] {
		case ' ':
			return &inst.PrintChar{}, nil
		case '\t':
			return &inst.PrintNumber{}, nil
		default:
			return nil, fmt.Errorf("Bad io parse space")
		}
	case '\t':
		switch runes[1] {
		case ' ':
			return &inst.ReadChar{}, nil
		case '\t':
			return &inst.ReadNumber{}, nil
		default:
			return nil, fmt.Errorf("Bad IO parse tab")
		}
	default:
		return nil, fmt.Errorf("Bad IO parse")
	}
}

func (p *Parser) parseMath() (inst.Instruction, error) {
	runes := []rune{}
	for i := 0; i < 2; i++ {
		r, n, err := p.r.ReadRune()
		if err != nil || n == 0 {
			return nil, fmt.Errorf("Bad math read")
		}

		runes = append(runes, r)
	}

	if len(runes) != 2 {
		return nil, fmt.Errorf("Bad math parse length")
	}

	switch runes[0] {
	case ' ':
		switch runes[1] {
		case ' ':
			return &inst.Add{}, nil
		case '\t':
			return &inst.Subtract{}, nil
		case '\n':
			return &inst.Multiply{}, nil
		default:
			return nil, fmt.Errorf("Bad space math")
		}
	case '\t':
		switch runes[1] {
		case ' ':
			return &inst.Divide{}, nil
		case '\t':
			return &inst.Modulo{}, nil
		default:
			return nil, fmt.Errorf("Bad tab bath")
		}
	default:
		return nil, fmt.Errorf("Bad math parse")
	}
}

func (p *Parser) parseHeap() (inst.Instruction, error) {
	r, n, err := p.r.ReadRune()
	if err != nil || n == 0 {
		return nil, fmt.Errorf("Bad heap read")
	}

	switch r {
	case ' ':
		return &inst.Store{}, nil
	case '\t':
		return &inst.Load{}, nil
	default:
		return nil, fmt.Errorf("Bad heap parse")
	}
}

func (p *Parser) parseNumber() (int64, error) {
	first := true
	var positive bool
	var val int64

	NumLoop:
	for {
		r, n, err := p.r.ReadRune()
		if err != nil || n == 0 {
			return 0, fmt.Errorf("Bad number parse: %w", err)
		}

		if first {
			first = false
			switch r {
			case ' ':
				positive = true
			case '\t':
				positive = false
			default:
				return 0, fmt.Errorf("Negatively positive")
			}
			continue
		}

		switch r {
		case ' ':
			val = (val << 1) | 0
		case '\t':
			val = (val << 1) | 1
		case '\n':
			break NumLoop
		default:
			return 0, fmt.Errorf("Imaginary number")
		}
	}

	if !positive {
		val = val * -1
	}
	return val, nil
}

func (p *Parser) parseLabel() (string, error) {
	label := strings.Builder{}
	for {
		r, n, err := p.r.ReadRune()
		if err != nil || n == 0 {
			return "", fmt.Errorf("Bad label parse")
		}

		switch r {
		case ' ', '\t':
			label.WriteRune(r)
		case '\n':
			return label.String(), nil
		default:
			return "", fmt.Errorf("Imaginary label")
		}
	}
}

func (p *Parser) parseStack() (inst.Instruction, error) {
	r, n, err := p.r.ReadRune()
	if err != nil || n == 0 {
		return nil, fmt.Errorf("Bad stack read")
	}

	if r == ' ' {
		num, err := p.parseNumber()
		if err != nil {
			return nil, fmt.Errorf("Stack push error: %w", err)
		}

		return &inst.Push{Value: num}, nil
	}

	r2, n, err := p.r.ReadRune()
	if err != nil || n == 0 {
		return nil, fmt.Errorf("Bad second stack read")
	}

	switch r {
	case '\t':
		num, err := p.parseNumber()
		if err != nil {
			return nil, fmt.Errorf("Stack tab error: %w", err)
		}

		switch r2 {
		case ' ':
			return &inst.Copy{Value: num}, nil
		case '\n':
			return &inst.Slide{Value: num}, nil
		}

	case '\n':
		switch r2 {
		case ' ':
			return &inst.Duplicate{}, nil
		case '\t':
			return &inst.Swap{}, nil
		case '\n':
			return &inst.Discard{}, nil
		default:
			return nil, fmt.Errorf("Stack newline error")
		}
	}

	return nil, fmt.Errorf("Stack error")
}

func (p *Parser) parseFlow() (inst.Instruction, error) {
	runes := []rune{}
	for i := 0; i < 2; i++ {
		r, n, err := p.r.ReadRune()
		if err != nil || n == 0 {
			return nil, fmt.Errorf("Bad math read")
		}

		runes = append(runes, r)
	}

	if len(runes) != 2 {
		return nil, fmt.Errorf("Bad flow parse length")
	}

	switch runes[0] {
	case ' ':
		label, err := p.parseLabel()
		if err != nil {
			return nil, fmt.Errorf("Flow space label error: %w", err)
		}

		switch runes[1] {
		case ' ':
			return &inst.Label{Value: label}, nil
		case '\t':
			return &inst.Call{Value: label}, nil
		case '\n':
			return &inst.Jump{Value: label}, nil
		}
	case '\t':
		if runes[1] == '\n' {
			return &inst.Return{}, nil
		}

		label, err := p.parseLabel()
		if err != nil {
			return nil, fmt.Errorf("Flow tab label error: %w", err)
		}

		switch runes[1] {
		case ' ':
			return &inst.JumpZero{Value: label}, nil
		case '\t':
			return &inst.JumpMinus{Value: label}, nil
		default:
			return nil, fmt.Errorf("Flow tab error")
		}
	case '\n':
		if runes[1] == '\n' {
			return &inst.Stop{}, nil
		}
		return nil, fmt.Errorf("Flow newline error ")
	}
	return nil, fmt.Errorf("Flow error")
}
