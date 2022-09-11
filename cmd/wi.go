// Whitespace interpreter
package main

import (
	"os"
	"io"
	//"strings"
	"fmt"

	"github.com/alexflint/go-arg"
	ws "github.com/zorchenhimer/whitespace"
)

type Arguments struct {
	Input string  `arg:"positional" help:"Input file.  Defaults to STDIN."`
	Output string `arg:"positional" help:"Output file.  Defaults to STDOUT."`
	Reader string `arg:"-r,--reader" help:"IO type.  Unimplemented."`
	Debug bool `arg:"-d,--debug"`
}

func main() {
	args := &Arguments{}
	arg.MustParse(args)
	err := run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}

func run(args *Arguments) error {
	var input io.ReadCloser
	var output io.WriteCloser
	var reader io.Reader // user input
	var err error

	if args.Input == "" {
		input = os.Stdin
	} else {
		input, err = os.Open(args.Input)
		if err != nil {
			return fmt.Errorf("Error opening input: %w", err)
		}
		defer input.Close()

		reader = os.Stdin
	}

	if args.Output == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(args.Output)
		if err != nil {
			return fmt.Errorf("Error creating output: %w", err)
		}
		defer output.Close()
	}

	e, err := ws.NewEngine(input)
	if err != nil {
		return fmt.Errorf("Engine error: %w", err)
	}
	e.Debug = args.Debug

	err = e.Run(reader, output)
	if err != nil {
		return fmt.Errorf("Run error: %w", err)
	}

	return nil
}
