package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
}

func NewCommander() *Commander {
	configLoader := config.NewFileSystemConfigLoader(infra.NewOSFileSystem())
	ui := ui.NewUI()
	commandBuilder, _ := NewCommandBuilder()
	clipboard := infra.NewClipboard()

	return &Commander{
		configLoader:   configLoader,
		ui:             *ui,
		commandBuilder: *commandBuilder,
		clipboard:      clipboard,
	}
}

func grayText(text string) string {
	return fmt.Sprintf("\033[90m%s\033[0m", text)
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
		if strings.ToLower(key) == name {
			return project, true
		}
	}

	return types.Project{}, false
}

func getConfigPath() (string, error) {
	fs := infra.NewOSFileSystem()
	homeDir, err := fs.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "lazyalias", "config.yaml"), nil
}

func (c *Commander) Run() error {
	configPath, err := getConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := c.configLoader.LoadConfig(configPath)
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

	cmd := c.commandBuilder.Build(ctx)

	if err := c.clipboard.Copy(cmd); err != nil {
		fmt.Printf("Could not copy to clipboard: %v\n", err)
		fmt.Print("Please copy the command manually")
	} else {
		fmt.Print(grayText("Command has been copied to clipboard!"))
	}

	fmt.Print(grayText("\nCommand to execute:"))
	fmt.Printf("\n\033[32m%s\033[0m\n", cmd)

	return nil
}
