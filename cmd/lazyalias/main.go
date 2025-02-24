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
	version     = "v0.1.3"
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

	if err := r.CopyToClipboard(cmd); err != nil {
		fmt.Printf("Could not copy to clipboard: %v\n", err)
		fmt.Printf("Please copy the command manually")
	} else {
		fmt.Printf(grayText("Command has been copied to clipboard!"))
	}

	fmt.Printf(grayText("\nCommand to execute:"))
	fmt.Printf("\n%s\n", cmd)
}
