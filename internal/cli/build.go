package cli

import (
	"flag"
	"fmt"
)

func initBuildCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("build", flag.ExitOnError)
	outArgs := cmd.String("out", "public", "Set a directory to output static content.")
	_ = outArgs

	return cmd
}

func runBuildCmd() error {
	fmt.Println("Building...")
	return nil
}
