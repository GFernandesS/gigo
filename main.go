package main

import (
	"log/slog"
	"os"

	"github.com/GFernandesS/gigo/internal/cmd/gigo"
	"github.com/GFernandesS/gigo/internal/cmd/list"
)

func main() {
	mainCommand := gigo.GetCommand()

	mainCommand.CompletionOptions.HiddenDefaultCmd = true

	mainCommand.AddCommand(list.GetCommand())

	if err := mainCommand.Execute(); err != nil {
		slog.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}
