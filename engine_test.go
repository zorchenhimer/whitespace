package whitespace

import (
	"testing"
	"strings"
	//"io"
)


func TestEngineBasic(t *testing.T) {
	source := "   \t\n   \t \n\t   \t\n \t\n\n\n"
	expected := `3`

	reader := strings.NewReader(source)
	e, err := NewEngine(reader)
	if err != nil {
		t.Fatalf("Engine creation fail: %s", err)
	}

	out := &strings.Builder{}
	err = e.Run(nil, out)
	if err != nil {
		t.Fatalf("Run fail: %s", err)
	}

	if out.String() != expected {
		t.Fatalf("Unexpected output.\n Rec: %q\n Exp: %q", out.String(), expected)
	}
	t.Logf("output: %q", out.String())
}
