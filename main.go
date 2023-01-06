package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: chutoro [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runRepl()
	}
}

func runFile(filename string) {
	src, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Cannot find %s\n", filename)
		os.Exit(66)
	}
	run(string(src))
}

func runRepl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Unexpected error while reading from stdin: %v\n", err)
			os.Exit(70)
		}
		run(line)
	}
}

func run(src string) {
	scanner := newScanner(src)
	if err := scanner.scan(); err != nil {
        panic(err)
    }

    for _, t := range scanner.tokens {
        fmt.Println(t.Text)
    }

	// TODO: uncomment when ready
	// parser := NewParser(tch)
	// stmts := parser.Parse()
}
