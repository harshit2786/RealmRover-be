package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server-cli",
	Short: "Server CLI is a tool to manage your application",
	Long:  "Server CLI provides commands to manage migrations, seed data, and more.",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
