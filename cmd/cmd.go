package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Init() {
	scaffoldCmd := initScaffoldCmd()
	configCmd := initConfigCmd()
	buildCmd := initBuildCmd()

	if len(os.Args) < 2 {
		fmt.Println("You need to supply at least one command for Generaidor to run. Please refer to usage!")
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "scaffold":
		scaffoldCmd.Parse(os.Args[2:])
		runScaffoldCmd()
	case "config":
		configCmd.Parse(os.Args[2:])
		runConfigCmd()
	case "build":
		buildCmd.Parse(os.Args[2:])
		runBuildCmd()
	}
}

func initBuildCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("build", flag.ExitOnError)
	outArg := cmd.String("out", "public", "Set a directory to output static content.")
	fmt.Println("outing...", outArg)
	return cmd
}

func initConfigCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("config", flag.ExitOnError)

	return cmd
}

func initScaffoldCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("scaffold", flag.ExitOnError)

	return cmd
}

func runBuildCmd() {
	fmt.Println("Building...")
}

func runScaffoldCmd() {
	fmt.Println("Scaffolding...")
}

func runConfigCmd() {
	fmt.Println("Configuring...")
}
