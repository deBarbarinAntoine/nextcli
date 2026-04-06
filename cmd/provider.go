package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var providerCmd = &cobra.Command{
	Aliases: []string{"pr", "pro", "prov"},
	Use:     "provider [name]",
	Short:   "Generate a provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		
		if len(args) == 0 {
			var err error
			name, err = internal.GetUserInput("Type the name of the provider file you want to generate:")
			if err != nil {
				return err
			}
		} else {
			name = args[0]
		}
		
		nextFile := internal.NewNextFile(name, enum.FileTypes.Provider)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(providerCmd)
}
