package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Aliases: []string{"r", "rep"},
	Use:     "repository [name]",
	Short:   "Generate a repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the repository file you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Repository)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(repositoryCmd)
}
