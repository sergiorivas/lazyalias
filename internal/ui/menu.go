package ui

import (
	"fmt"
	"strings"

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
		Label:     "Select a project",
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

func (ui *UI) ShowArgMenu(arg config.Arg) (string, error) {
	if arg.Options == "*" || arg.Options == "" {

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ .  }} ",
			Success: "{{ .  }} ",
		}

		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("✏️ Enter value for %s:", arg.Name),
			Templates: templates,
		}
		value, err := prompt.Run()
		if err != nil {
			return "", err
		}

		return value, nil
	}

	options := strings.Split(arg.Options, "|")

	for i, opt := range options {
		options[i] = strings.TrimSpace(opt)
	}

	templates := &promptui.SelectTemplates{
		Label:    fmt.Sprintf("Select an option for %s", arg.Name),
		Active:   "👉 {{ . | cyan }}",
		Inactive: "  {{ . | white }}",
		Selected: fmt.Sprintf("✏️ Selected option for %s: {{ . | green }}", arg.Name),
	}

	prompt := promptui.Select{
		Label:     "Select an option",
		Items:     options,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return options[i], nil
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
		Label:     "Select a command",
		Items:     allCommands,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return config.Command{}, err
	}

	if len(allCommands[i].Args) > 0 {
		for j, arg := range allCommands[i].Args {
			value, err := ui.ShowArgMenu(arg)
			allCommands[i].Args[j].Value = value
			if err != nil {
				return config.Command{}, err
			}
		}
	}

	return allCommands[i], nil
}
