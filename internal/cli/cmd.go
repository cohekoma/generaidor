package cli

import (
	"errors"
	"fmt"
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
		if err := scaffoldCmd.Parse(args[1:]); err != nil {
			return err
		}
		return runScaffoldCmd()
	case "config":
		if err := configCmd.Parse(args[1:]); err != nil {
			return err
		}
		return runConfigCmd()
	case "build":
		if err := buildCmd.Parse(args[1:]); err != nil {
			return err
		}
		return runBuildCmd()

	default:
		fmt.Println("Unknown command. Please refer to usage guide below.")
		printDefaultUsage()
		return errors.New("unknown command!")
	}
}

func printDefaultUsage() {
	fmt.Println("Usage: generaidor <command> [options]")
	fmt.Println("Available commands:")
	fmt.Println("\tscaffold: Create default structure for your project.")
	fmt.Println("\tconfig: Config custom directories to scan and ouput result.")
	fmt.Println("\tbuild: Generate static content based on the content directory.")
}
