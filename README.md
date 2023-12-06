# gocli

### Build the application

Run the command:

```shell
cd cmd/cli
go build -o gocli
```

This compiles your program and generates an executable file.

### Run the application

```shell
./gocli
```

You should see some kind of text output displayed in the terminal.

### Adding more commands

Here's an example of adding a `version` command:

```go
package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use: "mycli",
        Short: "MyCLI is a simple CLI program in Go",
        Long: "MyCLI is a simple CLI program that demonstrates the use of the cobra library in Go.",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Welcome to MyCLI!")
        },
    }
	
    err := rootCmd.Execute()
    if err != nil {
        return
    }
}
```


