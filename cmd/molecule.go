package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// moleculeCmd represents the molecule command
var moleculeCmd = &cobra.Command{
	Aliases: []string{"m", "mol"},
	Use:     "molecule [name]",
	Short:   "Generate a molecule component",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the molecule component you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Molecule)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(moleculeCmd)
}
