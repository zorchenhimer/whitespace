package instructions

import (
	"fmt"
	"strings"
)

func EncodeNumber(n int64) string {
	negative := false
	if n < 0 {
		negative = true
		n *= -1
	}

	b := fmt.Sprintf("%b", n)
	enc := strings.ReplaceAll(strings.ReplaceAll(b, "1", "\t"), "0", " ")
	if negative {
		enc = "\t"+enc
	} else {
		enc = " "+enc
	}
	return enc+"\n"
}

func EncodeLabel(l string) string {
	return strings.ReplaceAll(strings.ReplaceAll(l, "t", "\t"), "s", " ") + "\n"
}

func DecodeLabel(l string) string {
	return strings.ReplaceAll(strings.ReplaceAll(l, "\t", "t"), " ", "s")
}
