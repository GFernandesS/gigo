package list

import (
	"github.com/GFernandesS/gigo/internal/api"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all templates available",
		Run:   run,
	}
}

func run(cmd *cobra.Command, args []string) {
	availableTemplates, err := api.ListTemplates()

	if err != nil {
		return
	}

	for template := range availableTemplates {
		cmd.Println(template)
	}
}
