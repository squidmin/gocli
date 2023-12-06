package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Define a type for a function that returns a *cobra.Command
type commandFunc func() *cobra.Command

func main() {
	rootCmd := newRootCmd()

	// List of functions that return subcommands
	commands := []commandFunc{
		newHelloCmd,
		newGoodbyeCmd,
	}

	// Iterate over the slice and add each command to the root command
	for _, cmdFunc := range commands {
		rootCmd.AddCommand(cmdFunc())
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
