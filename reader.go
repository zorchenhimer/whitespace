package whitespace

import (
	"io"
	"bufio"
	"fmt"
)

// Reader that ignores everything other than space, tab, and newline
type Reader struct {
	base *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{base: bufio.NewReader(r)}
}

func (reader *Reader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, fmt.Errorf("zero length buffer")
	}

	read := 0
	var err error
	var n int
	var r rune

	for read < len(p) && err == nil {
		r, n, err = reader.base.ReadRune()

		if n != 0 {
			switch r {
			case ' ', '\t', '\n':
				p[read] = byte(r)
				read++
			default:
				// ignore everything else
			}
		} else {
			return read, io.EOF
		}

		if err != nil {
			break
		}
	}

	// this read value probably won't be accurate, lol
	return read, err
}

func (r *Reader) ReadRune() (rune, int, error) {
	for {
		r, n, err := r.base.ReadRune()
		if err != nil {
			if r == ' ' || r == '\t' || r == '\n' {
				return r, n, err
			}

			if err == io.EOF {
				return 0x00, 0, err
			}

			// Following the behaviour of bufio.ReadRune(), but "invalid"
			// runes are all the non-whitespace ones.
			return '\uFFFD', 1, err
		}

		if r == ' ' || r == '\t' || r == '\n' {
			return r, n, nil
		}
	}

	//return '\uFFFD', 1, fmt.Errorf("How did you get here?")
}
