// Whitespace transpiler
package main

import (
	"os"
	"io"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	ws "github.com/zorchenhimer/whitespace"
	ins "github.com/zorchenhimer/whitespace/instructions"
)

type Args struct {
	Input string  `arg:"positional" help:"Input filename.  Defaults to STDIN."`
	Output string `arg:"positional" help:"Output filename.  Defaults to STDOUT"`

	Assembly bool `arg:"-a,--to-asm" help:"Translate to assembly"`
	Wsp bool `arg:"-w,--to-wsp" help:"Translate to whitespace"`
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type convertFunc func(reader io.Reader, writer io.Writer) error

func run() error {
	args := &Args{}
	arg.MustParse(args)

	if args.Assembly == args.Wsp && args.Assembly {
		return fmt.Errorf("Cannot translate to both asm and wsp")
	}

	var input io.Reader
	var output io.Writer
	var err error

	if args.Input == "" {
		input = os.Stdin
	} else {
		inputfile, err := os.Open(args.Input)
		if err != nil {
			return fmt.Errorf("error opening input file: %w", err)
		}
		defer inputfile.Close()

		input = inputfile
	}

	if args.Output == "" {
		output = os.Stdout
	} else {
		outputbuf := &bytes.Buffer{}
		defer func() {
			err = os.WriteFile(args.Output, outputbuf.Bytes(), 0644)
			if err != nil {
				panic(fmt.Sprintf("error writing output file: %w", err))
			}
		}()

		output = outputbuf
	}

	var cfunc convertFunc

	if args.Assembly {
		cfunc = toAsm
	} else if args.Wsp {
		cfunc = toWhitespace

	} else if strings.HasSuffix(args.Input, ".wsp") {
		// whitespace -> asm
		cfunc = toAsm

	} else if strings.HasSuffix(args.Input, ".wsa") {
		// asm -> whitespace
		cfunc = toWhitespace
	} else {
		cfunc = toWhitespace
	}

	return cfunc(input, output)

	//return err
}

func toAsm(reader io.Reader, writer io.Writer) error {
	parser := ws.NewParser(ws.NewReader(reader))
	//parser.Debug = true
	inst, err := parser.Parse()

	if len(inst) != 0 {
		for _, i := range inst {
			fmt.Fprintln(writer, i.Asm())
		}
	}

	if err != nil {
		return fmt.Errorf("Parse error: %w", err)
	}
	return nil
}

func toWhitespace(reader io.Reader, writer io.Writer) error {
	input, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("Unable to read input: %w", err)
	}

	lines := strings.Split(string(input), "\n")

	for i, l := range lines {
		l = strings.ToLower(strings.TrimSpace(l))
		// empty lines & comments
		if len(l) == 0 || strings.HasPrefix(l, "#") {
			continue
		}

		parts := strings.Split(l, " ")
		switch parts[0] {
		case "duplicate":
			fmt.Fprint(writer, " \n ")
		case "swap":
			fmt.Fprint(writer, " \n\t")
		case "discard":
			fmt.Fprint(writer, " \n\n")
		case "add":
			fmt.Fprint(writer, "\t   ")
		case "subtract":
			fmt.Fprint(writer, "\t  \t")
		case "multiply":
			fmt.Fprint(writer, "\t  \n")
		case "divide":
			fmt.Fprint(writer, "\t \t ")
		case "modulo":
			fmt.Fprint(writer, "\t \t\t")
		case "store":
			fmt.Fprint(writer, "\t\t ")
		case "load":
			fmt.Fprint(writer, "\t\t\t")
		case "return":
			fmt.Fprint(writer, "\n\t\n")
		case "stop":
			fmt.Fprint(writer, "\n\n\n")
		case "printchar":
			fmt.Fprint(writer, "\t\n  ")
		case "printnumber":
			fmt.Fprint(writer, "\t\n \t")
		case "readchar":
			fmt.Fprint(writer, "\t\n\t ")
		case "readnumber":
			fmt.Fprint(writer, "\t\n\t\t")

		default:
			// everything with an argument
			if len(parts) != 2 {
				return fmt.Errorf("missing argument on line %d", i+1)
			}

			if parts[0] == "copy" || parts[0] == "push" {
				n, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return fmt.Errorf("number parse error on line %d: %w", i+1, err)
				}

				if parts[0] == "copy" {
					fmt.Fprint(writer, " \t ")
				} else {
					fmt.Fprint(writer, "  ")
				}
				fmt.Fprint(writer, ins.EncodeNumber(n))
			} else {
				switch parts[0] {
				case "label":
					fmt.Fprint(writer, "\n  ")
				case "call":
					fmt.Fprint(writer, "\n \t")
				case "jump":
					fmt.Fprint(writer, "\n \n")
				case "jumpzero":
					fmt.Fprint(writer, "\n\t ")
				case "jumpminus":
					fmt.Fprint(writer, "\n\t\t")
				}
				fmt.Fprint(writer, ins.EncodeLabel(parts[1]))
			}
		}
	}
	return nil
}
