package ui

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/sergiorivas/lazyalias/internal/types"
)

func (ui *UI) ShowArgMenu(arg types.Arg) (string, error) {
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
