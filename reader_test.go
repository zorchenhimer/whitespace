package whitespace

import (
	"testing"
	"strings"
	"io"
)

// TODO
//func TestRead(t *testing.T) {
//}

func TestReadRune(t *testing.T) {
	tests := []struct{
		input string
		expected []rune
	}{
		{"    ", []rune{' ',' ',' ',' '}},
		{"one two three four ", []rune{' ',' ',' ',' '}},
		{"one two\tthree\nfour ", []rune{' ','\t','\n',' '}},
		{"\t\t\t\t", []rune{'\t', '\t', '\t', '\t'}},
		{"\n\n\n\n", []rune{'\n', '\n', '\n', '\n'}},
		{" \t\n", []rune{' ', '\t', '\n'}},
	}

	MainLoop:
	for id, tst := range tests {
		sr := strings.NewReader(tst.input)
		reader := NewReader(sr)
		out := []rune{}

		for {
			r, n, err := reader.ReadRune()
			if err != nil && err != io.EOF {
				t.Logf("[%d] ReadRune() returned error: %s", id, err)
				t.Fail()
				continue MainLoop
			}

			if n > 0 {
				out = append(out, r)
			}

			if err == io.EOF {
				break
			}
		}

		if len(tst.expected) != len(out) {
			t.Logf("[%d] Unexpected size: %d; expected %d; %v; %v", id, len(out), len(tst.expected), out, tst.expected)
			t.Fail()
			continue
		}

		for i := 0; i < len(tst.expected); i++ {
			if tst.expected[i] != out[i] {
				t.Logf("[%d] Received %v; expected %v", id, out, tst.expected)
				t.Fail()
				break
			}
		}
	}
}
