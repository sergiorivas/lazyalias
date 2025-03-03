package ui

import (
	"github.com/manifoldco/promptui"
	"github.com/sergiorivas/lazyalias/internal/types"
)

func (ui *UI) ShowCommandMenu(commands []types.Command) (types.Command, error) {
	backCommand := types.Command{
		Name:    "⬅️ Back to Projects",
		Command: BackToProject,
	}
	allCommands := commands
	allCommands = append(allCommands, backCommand)

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "{{ if ne .Command \"" + BackToProject + "\" }}📟 Selected Command: {{ .Name | green }} {{ else }}{{ .Name }} {{ end }}",
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
		return types.Command{}, err
	}

	if len(allCommands[i].Args) > 0 {
		for j, arg := range allCommands[i].Args {
			value, err := ui.ShowArgMenu(arg)
			allCommands[i].Args[j].Value = value
			if err != nil {
				return types.Command{}, err
			}
		}
	}

	return allCommands[i], nil
}
