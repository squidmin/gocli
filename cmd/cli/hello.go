package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newHelloCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "subcommand_1",
		Short: "Prints the hello message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from subcommand_1!")
		},
	}
}
