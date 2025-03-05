package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	"github.com/sergiorivas/lazyalias/internal/config"
	"github.com/sergiorivas/lazyalias/internal/infra"
	"github.com/sergiorivas/lazyalias/internal/types"
	"github.com/sergiorivas/lazyalias/internal/ui"
)

type Commander struct {
	configLoader   config.ConfigLoader
	ui             ui.UI
	commandBuilder CommandBuilder
	clipboard      infra.Clipboard
	history        infra.HistoryFileUpdater
}

func NewCommander() *Commander {
	configLoader := config.NewFileSystemConfigLoader(infra.NewOSFileSystem())
	ui := ui.NewUI()
	commandBuilder, _ := NewCommandBuilder()
	clipboard := infra.NewClipboard()
	history, _ := infra.NewHistoryFileUpdater()

	return &Commander{
		configLoader:   configLoader,
		ui:             *ui,
		commandBuilder: *commandBuilder,
		clipboard:      clipboard,
		history:        *history,
	}
}

// getCurrentProjectName returns the name of the current directory
func getCurrentProjectName() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(currentDir), nil
}

// findProjectByName returns a project if it exists in the config
func findProjectByName(cfg config.Config, name string) (types.Project, bool) {
	// First try exact match
	if project, exists := cfg[name]; exists {
		return project, true
	}

	// Then try case-insensitive match
	name = strings.ToLower(name)
	for key, project := range cfg {
		if strings.EqualFold(key, name) {
			return project, true
		}
	}

	return types.Project{}, false
}

func (c *Commander) Run() error {
	cfg, err := c.configLoader.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var project types.Project

	// Check if current directory matches a project
	currentProjectName, err := getCurrentProjectName()
	if err != nil {
		log.Fatal(err)
	}

	var projects []types.Project
	for _, p := range cfg {
		projects = append(projects, p)
	}

	if matchedProject, found := findProjectByName(cfg, currentProjectName); found {
		project = matchedProject
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("🗂️ Selected Project: %s\n", green(project.Name))
	} else {
		project, err = c.ui.ShowProjectMenu(projects)
		if err != nil {
			log.Fatal(err)
		}
	}

	command, err := c.ui.ShowCommandMenu(project.Commands)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if command.Command != ui.BackToProject {
			break
		}

		project, err = c.ui.ShowProjectMenu(projects)
		if err != nil {
			log.Fatal(err)
		}

		command, err = c.ui.ShowCommandMenu(project.Commands)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx := ExecutionContext{
		TargetDir: project.Folder,
		Command:   command,
		Project:   project,
	}

	cmd := c.commandBuilder.Build(&ctx)

	err = c.clipboard.Copy(cmd)
	if err != nil {
		fmt.Printf("Could not copy to clipboard: %v\n", err)
		fmt.Print("Please copy the command manually")
	} else {
		color.RGB(127, 127, 127).Print("Command has been copied to clipboard!\n")
	}

	err = c.history.Add(cmd)
	if err != nil {
		fmt.Printf("Could not save to history: %v\n", err)
	}

	color.RGB(127, 127, 127).Print("Command to execute:\n")
	color.Green("%s\n", cmd)

	return nil
}
