// Package main implements a CLI tool for managing GitHub following relationships.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-follow-cli",
		Short: "A CLI tool for managing GitHub following relationships",
		Long: `go-follow-cli is a command-line tool that allows you to follow GitHub users
programmatically using the GitHub API. It requires a GitHub personal access token
with 'user:follow' permissions.`,
	}

	// Check for GitHub token on root command
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if os.Getenv("GITHUB_TOKEN") == "" {
			fmt.Fprintln(os.Stderr, "Error: GITHUB_TOKEN environment variable is required")
			fmt.Fprintln(os.Stderr, "Please set it with: export GITHUB_TOKEN=your_token_here")
			os.Exit(1)
		}
	}

	// Add all commands
	rootCmd.AddCommand(
		newFollowCommand(),
		newFollowAllCommand(),
		newFollowRandomCommand(),
		newFollowNCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
