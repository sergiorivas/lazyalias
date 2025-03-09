package ui

import (
	"github.com/manifoldco/promptui"
	"github.com/sergiorivas/lazyalias/internal/types"
)

func (ui *UI) ShowProjectMenu(projects []types.Project) (types.Project, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "🗂️ Selected Project: {{ .Name | green }}",
	}

	prompt := promptui.Select{
		Label:     "Select a project",
		Items:     projects,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return types.Project{}, err
	}

	return projects[i], nil
}
