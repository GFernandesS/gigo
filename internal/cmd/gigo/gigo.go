package gigo

import (
	"log/slog"
	"os"

	"github.com/GFernandesS/gigo/internal/api"

	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gigo [tech 1]  [tech 2] ...",
		Short: "Generate a .gitignore file for your project",
		Long:  "🗑️  Gigo is a command-line tool that helps you generate .gitignore files for your projects based on the technologies you are using based on https://gitignore.io.",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}
}

func run(cmd *cobra.Command, args []string) {
	if !validateTechsByTemplates(args) {
		return
	}

	gitIgnoreContent, err := api.GenerateGitIgnore(args)

	if err != nil {
		return
	}

	if err := os.WriteFile(".gitignore", gitIgnoreContent, 0644); err != nil {
		slog.Error("Error writing .gitignore file", "error", err)
		return
	}

	slog.Info("✅ .gitignore generated successfully")
}

func validateTechsByTemplates(techs []string) bool {
	availableTemplates, err := api.ListTemplates()

	if err != nil {
		slog.Error("Error fetching available templates", "error", err)
		return false
	}

	for _, tech := range techs {
		if _, exists := availableTemplates[tech]; !exists {
			slog.Warn("Template not found. Use gigo list to see available templates", "template", tech)
			return false
		}
	}

	return true
}
