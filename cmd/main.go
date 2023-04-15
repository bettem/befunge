package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage befunge [file]\n")
		os.Exit(64)
	}

	// Read the program
	file, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("unable to read source file: '%s'", err)
		os.Exit(74)
	}

	// Scan the program into a series of tokens
	program := ScanInput(string(file))
	program.Debug()
	Interpret(program)
}
