package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "This is the help command",
		Long:  `Simple, demonstrative CLI program.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to bq!")
		},
	}
}
