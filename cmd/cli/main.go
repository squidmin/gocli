package main

import (
	"fmt"
	"os"
)

func main() {
	rootCmd := newRootCmd()
	rootCmd.AddCommand(newHelloCmd())
	rootCmd.AddCommand(newGoodbyeCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
