// Command hello prints a greeting.
//
// Usage:
//
//	hello [-name NAME]
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run parses args (without the program name), writes the greeting to stdout,
// and returns the process exit code.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("hello", flag.ContinueOnError)
	fs.SetOutput(stderr)
	name := fs.String("name", "world", "name to greet")

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	fmt.Fprintf(stdout, "Hello, %s!\n", *name)
	return 0
}
