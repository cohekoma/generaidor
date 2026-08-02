package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func Run(args []string) error {
	scaffoldCmd := initScaffoldCmd()
	configCmd := initConfigCmd()
	buildCmd := initBuildCmd()

	if len(args) == 0 {
		return errors.New("missing command!")
	}

	switch args[0] {
	case "scaffold":
		scaffoldCmd.Parse(os.Args[2:])
		runScaffoldCmd()
	case "config":
		configCmd.Parse(os.Args[2:])
		runConfigCmd()
	case "build":
		buildCmd.Parse(os.Args[2:])
		runBuildCmd()

	default:
		fmt.Println("Unknown command. Please refer to usage guide below.")
		printDefaultUsage()
		return errors.New("unknown command!")
	}

	return nil
}

func printDefaultUsage() {
	fmt.Println("Usage: generaidor <command> [options]")
	fmt.Println("Available commands:")
	fmt.Println("\tscaffold: Create default structure for your project.")
	fmt.Println("\tconfig: Config custom directories to scan and ouput result.")
	fmt.Println("\tbuild: Generate static content based on the content directory.")
}

func initBuildCmd() *flag.FlagSet {
	cmd := flag.NewFlagSet("build", flag.ExitOnError)
	outArgs := cmd.String("out", "public", "Set a directory to output static content.")
	_ = outArgs

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
