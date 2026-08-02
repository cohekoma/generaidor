package main

import (
	"fmt"
	"os"

	"github.com/cohekoma/generaidor/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error occured. Err: %v\n", err)
		os.Exit(1)
	}
}
