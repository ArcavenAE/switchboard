package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var version = "dev"

func run(stdout io.Writer, args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(stdout)
	showVersion := fs.Bool("version", false, "print version and exit")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	if *showVersion || fs.NArg() == 0 {
		if _, err := fmt.Fprintf(stdout, "switchboard %s\n", version); err != nil {
			return fmt.Errorf("write version: %w", err)
		}
	}

	return nil
}

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		os.Exit(1)
	}
}
