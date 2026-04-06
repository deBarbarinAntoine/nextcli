package cmd

import (
	"os"
	
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nextcli",
	Short: "NextCLI is a simple CLI tool for Next.js project to help generate files and components.",
	Long: `NextCLI is a simple CLI tool for Next.js project to help generate files and components.
It can be used to generate pages, atomic design components, models, providers, repositories, services and more.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
