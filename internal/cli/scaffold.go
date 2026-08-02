package cli

import (
	"flag"
	"fmt"
)

func initScaffoldCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("scaffold", flag.ExitOnError)

	return cmd
}

func runScaffoldCmd() {
	fmt.Println("Scaffolding...")
}
