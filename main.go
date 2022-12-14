package main

import (
	"fmt"
	commands "server/cmd"

	"github.com/spf13/cobra"
)

func main() {
	// Create a new cobra command for the "hello" command
	rootCmd := &cobra.Command{
		Use:   "server [command]",
		Short: "Say hello to the specified person",
		Long:  "This command prints a friendly greeting to the specified person.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", args[0])
		},
	}

	rootCmd.AddCommand(commands.VersionCmd)
	rootCmd.AddCommand(commands.ServerCmd)
	rootCmd.AddCommand(commands.WebCmd)

	// Execute the root command, which will run the "hello" command if it
	// is specified on the command line
	rootCmd.Execute()
}
