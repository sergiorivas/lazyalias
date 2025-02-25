package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config map[string]Project

type Project struct {
	Name     string    `yaml:"name"`
	Commands []Command `yaml:"commands"`
	Folder   string    `yaml:"folder,omitempty"`
	Key      string    `yaml:"-"` // This will store the map key
}

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Args    []Arg  `yaml:"args"`
}

type Arg struct {
	Name    string `yaml:"name"`
	Options string `yaml:"options"`
	Value   string
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	for key, project := range config {
		project.Key = key
		if project.Name == "" {
			project.Name = key
		}
		if project.Folder != "" {
			absPath, err := filepath.Abs(project.Folder)
			if err != nil {
				return nil, err
			}
			project.Folder = absPath
			config[key] = project
		} else {
			config[key] = project
		}
	}

	return config, nil
}
