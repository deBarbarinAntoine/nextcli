package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// organismCmd represents the organism command
var organismCmd = &cobra.Command{
	Aliases: []string{"o"},
	Use:     "organism [name]",
	Short:   "Generate an organism component",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the organism component you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Organism)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(organismCmd)
}
