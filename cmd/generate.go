package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Aliases: []string{"g", "gen"},
	Use:     "generate",
	Short:   "Generate a base file for your Next.js application",
	Long: `This command cannot be used alone,
you need to specify the kind of file you want to generate.
Example: nextcli generate page /post/[id]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify the kind of file you want to generate.")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
