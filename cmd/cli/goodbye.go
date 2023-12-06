package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newGoodbyeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "subcommand_2",
		Short: "Prints the goodbye message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Goodbye from subcommand_2!")
		},
	}
}
