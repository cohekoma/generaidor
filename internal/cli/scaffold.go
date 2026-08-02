package cli

import (
	"flag"
	"fmt"
)

func initScaffoldCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("scaffold", flag.ExitOnError)

	return cmd
}

func runScaffoldCmd() error {
	fmt.Println("Scaffolding...")
	return nil
}
