package main

import (
	"fmt"
	"os"

	"gogo/command"
)

type File string

func (f File) Writer (p []byte) (n int, err error) {
	return len(f), nil
}

func main() {
	if err := command.RootCmd.Execute() ; err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}
