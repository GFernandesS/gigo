package api

import (
	"io"
	"log/slog"
	"net/http"
	"strings"
)

const gitIgnoreBaseURL = "https://www.toptal.com/developers/gitignore/api/"

func ListTemplates() (map[string]struct{}, error) {
	resp, err := http.Get(gitIgnoreBaseURL + "list")

	if err != nil {
		slog.Error("Error fetching templates list", "error", err)
		return nil, err
	}

	defer resp.Body.Close()

	templates, err := io.ReadAll(resp.Body)

	if err != nil {
		slog.Error("Error reading templates list", "error", err)
		return nil, err
	}

	templatesList := strings.Split(string(templates), ",")

	templatesMap := make(map[string]struct{})

	for _, template := range templatesList {
		templatesMap[template] = struct{}{}
	}

	return templatesMap, nil
}

func GenerateGitIgnore(techs []string) ([]byte, error) {
	resp, err := http.Get(gitIgnoreBaseURL + strings.Join(techs, ","))

	if err != nil {
		slog.Error("Error getting .gitignore templates", "error", err)
		return nil, err
	}

	defer resp.Body.Close()

	gitIgnoreContent, err := io.ReadAll(resp.Body)

	if err != nil {
		slog.Error("Error reading .gitignore content", "error", err)
		return nil, err
	}

	return gitIgnoreContent, nil
}
