# Whitespace

This is an implementation of the [Whitespace esoteric programming language](https://en.wikipedia.org/wiki/Whitespace_\(programming_language\)).

# Usage

## wt

This is a utility to translate between an assembly representation of whitespace
and pure whitespace.

    Usage: wt [--to-asm] [--to-wsp] [INPUT [OUTPUT]]

    Positional arguments:
      INPUT                  Input filename.  Defaults to STDIN.
      OUTPUT                 Output filename.  Defaults to STDOUT

    Options:
      --to-asm, -a           Translate to assembly
      --to-wsp, -w           Translate to whitespace
      --help, -h             display this help and exit

## wi

This is the whitespace interpreter.  It only reads pure whitespace, not the
assembly representation.

    Usage: wi [--debug] [INPUT [OUTPUT]]

    Positional arguments:
      INPUT                  Input file.  Defaults to STDIN.
      OUTPUT                 Output file.  Defaults to STDOUT.

    Options:
      --debug, -d
      --help, -h             display this help and exit

If the input is a file (ie, passed as an argument), user input uses STDIN.
Otherwise, no user input is allowed.

# License

MIT License.  See `LICENSE.md`.
