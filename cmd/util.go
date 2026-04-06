package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// utilCmd represents the util command
var utilCmd = &cobra.Command{
	Aliases: []string{"u"},
	Use:     "util [name]",
	Short:   "Generate a util file",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the util file you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Util)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(utilCmd)
}
