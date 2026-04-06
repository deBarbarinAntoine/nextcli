package cmd

import (
	"nextcli/internal"
	"nextcli/internal/enum"
	
	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Aliases: []string{"p"},
	Use:     "page [url]",
	Short:   "Generate a page",
	RunE: func(cmd *cobra.Command, args []string) error {
		var url string
		
		if len(args) == 0 {
			var err error
			url, err = internal.GetUserInput("Type the URL of the page you want to generate:")
			if err != nil {
				return err
			}
		} else {
			url = args[0]
		}
		
		nextFile := internal.NewNextFile(url, enum.FileTypes.Page)
		
		return nextFile.Generate()
	},
}

func init() {
	generateCmd.AddCommand(pageCmd)
}
