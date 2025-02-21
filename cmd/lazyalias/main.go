package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sergiorivas/lazyalias/internal/config"
	"github.com/sergiorivas/lazyalias/internal/runner"
	"github.com/sergiorivas/lazyalias/internal/ui"
)

// getCurrentProjectName returns the name of the current directory
func getCurrentProjectName() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(currentDir), nil
}

// findProjectByName returns a project if it exists in the config
func findProjectByName(cfg config.Config, name string) (config.Project, bool) {
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

	return config.Project{}, false
}

func main() {

	fmt.Printf("Welcome to LAZYALIAS 🎉🎉🎉\n")
	// Load configuration
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := filepath.Join(homeDir, ".config", "lazyalias", "config.yaml")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	ui := ui.NewUI()
	var project config.Project

	// Check if current directory matches a project
	currentProjectName, err := getCurrentProjectName()
	if err != nil {
		log.Fatal(err)
	}

	if matchedProject, found := findProjectByName(cfg, currentProjectName); found {
		project = matchedProject
	} else {
		// If no match, show project selection menu
		var projects []config.Project
		for _, p := range cfg {
			projects = append(projects, p)
		}

		project, err = ui.ShowProjectMenu(projects)
		if err != nil {
			log.Fatal(err)
		}
	}

	command, err := ui.ShowCommandMenu(project.Commands)
	if err != nil {
		log.Fatal(err)
	}

	r, err := runner.NewRunner()
	if err != nil {
		log.Fatal(err)
	}

	ctx := runner.ExecutionContext{
		TargetDir: project.Folder,
		Command:   command,
		Project:   project,
	}

	cmd := r.PrepareCommand(ctx)

	fmt.Printf("\n💻 Command to execute:\n")
	fmt.Printf("------------------------\n")
	fmt.Printf("%s\n\n", cmd)

	if err := r.CopyToClipboard(cmd); err != nil {
		fmt.Printf("Could not copy to clipboard: %v\n", err)
		fmt.Printf("Please copy the command manually\n")
	} else {
		fmt.Printf("📋 Command has been copied to clipboard!\n")
	}
}
