package ui

import (
	"github.com/manifoldco/promptui"
	"github.com/sergiorivas/lazyalias/internal/config"
)

type UI struct{}

func NewUI() *UI {
	return &UI{}
}

func (ui *UI) ShowProjectMenu(projects []config.Project) (config.Project, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "🗂️ Selected Project: {{ .Name | green }}",
		Details: `
--------- Project ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Commands:" | faint }}	{{ len .Commands }} available
{{ if .Folder }}{{ "Folder:" | faint }}	{{ .Folder }}{{ end }}`,
	}

	prompt := promptui.Select{
		Label:     "Select Project",
		Items:     projects,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return config.Project{}, err
	}

	return projects[i], nil
}

func (ui *UI) ShowCommandMenu(commands []config.Command) (config.Command, error) {
	backCommand := config.Command{
		Name:    "⬅️ Back to Projects",
		Command: "back-to-project",
	}
	allCommands := append(commands, backCommand)

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "{{ if ne .Command \"back-to-project\" }}📟 Selected Command: {{ .Name | green }} {{ else }}{{ .Name }} {{ end }}",
		Details: `
--------- Command ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Command:" | faint }}	{{ .Command }}`,
	}

	prompt := promptui.Select{
		Label:     "Select Command",
		Items:     allCommands,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return config.Command{}, err
	}

	return allCommands[i], nil
}
