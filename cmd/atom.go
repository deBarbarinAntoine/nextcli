package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// atomCmd represents the atom command
var atomCmd = &cobra.Command{
	Aliases: []string{"a"},
	Use:     "atom [name]",
	Short:   "Generate an atom component",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the atom component you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Atom)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(atomCmd)
}
