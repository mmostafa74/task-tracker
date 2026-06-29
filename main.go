package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-i" || os.Args[1] == "--interactive" {
		runREPL()
		return
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printUsage(os.Stdout)
		return
	}

	if err := runCommand(os.Args[1], os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
