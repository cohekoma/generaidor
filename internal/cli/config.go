package cli

import (
	"flag"
	"fmt"
)

func initConfigCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("config", flag.ExitOnError)

	return cmd
}

func runConfigCmd() {
	fmt.Println("Configuring...")
}
