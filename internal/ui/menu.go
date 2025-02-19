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
		Selected: "• Selected Project: {{ .Name | green }}",
		Details: `
--------- Project ----------
{{ "Name:" | faint }}	{{ .Key }}
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
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "• Selected Command: {{ .Name | green }}",
		Details: `
--------- Command ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Command:" | faint }}	{{ .Command }}`,
	}

	prompt := promptui.Select{
		Label:     "Select Command",
		Items:     commands,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return config.Command{}, err
	}

	return commands[i], nil
}
