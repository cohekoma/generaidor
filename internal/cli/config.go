package cli

import (
	"flag"
	"fmt"
)

func initConfigCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("config", flag.ExitOnError)

	return cmd
}

func runConfigCmd() error {
	fmt.Println("Configuring...")
	return nil
}
