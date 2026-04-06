package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Aliases: []string{"s"},
	Use:     "service [name]",
	Short:   "Generate a service",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the service file you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Service)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(serviceCmd)
}
