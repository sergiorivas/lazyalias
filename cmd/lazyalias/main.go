package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sergiorivas/lazyalias/internal/config"
	"github.com/sergiorivas/lazyalias/internal/runner"
	"github.com/sergiorivas/lazyalias/internal/ui"
)

var (
	version     = "v0.1.5"
	showVersion = flag.Bool("version", false, "show version information")
)

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
	flag.Parse()

	if *showVersion {
		fmt.Printf("lazyalias version %s\n", version)
		return
	}

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

	var projects []config.Project
	for _, p := range cfg {
		projects = append(projects, p)
	}

	if matchedProject, found := findProjectByName(cfg, currentProjectName); found {
		project = matchedProject
	} else {
		project, err = ui.ShowProjectMenu(projects)
		if err != nil {
			log.Fatal(err)
		}
	}

	command, err := ui.ShowCommandMenu(project.Commands)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if command.Command != "back-to-project" {
			break
		}

		project, err = ui.ShowProjectMenu(projects)
		if err != nil {
			log.Fatal(err)
		}

		command, err = ui.ShowCommandMenu(project.Commands)
		if err != nil {
			log.Fatal(err)
		}
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

	if err := r.CopyToClipboard(cmd); err != nil {
		fmt.Printf("Could not copy to clipboard: %v\n", err)
		fmt.Print("Please copy the command manually")
	} else {
		fmt.Print(grayText("Command has been copied to clipboard!"))
	}

	fmt.Print(grayText("\nCommand to execute:"))
	fmt.Printf("\n\033[32m%s\033[0m\n", cmd)
}
