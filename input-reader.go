package whitespace

import (
	"io"
	"bufio"
)

type ReadNumberCallback func() (int64, error)
type ReadCharCallback func() (rune, error)

func ReadInteractiveChar() (rune, error) {
	return ' ', nil
}

func ReadInteractiveNumber() (int64, error) {
	return ' ', nil
}

type InputReader struct {
	r *bufio.Reader
}

func NewInputReader(reader io.Reader) *InputReader {
	return &InputReader{r: bufio.NewReader(reader)}
}

func (ir InputReader) ReadChar() (rune, error) {
	r, n, err := ir.r.ReadRune()
	if err != nil || n == 0 {
		return '\uFFFD', err
	}

	return r, nil
}

func (ir InputReader) ReadNumber() (int64, error) {
	return 0, nil
}
