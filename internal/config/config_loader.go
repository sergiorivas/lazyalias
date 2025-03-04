package config

import (
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/sergiorivas/lazyalias/internal/infra"
	"github.com/sergiorivas/lazyalias/internal/types"
)

type Config map[string]types.Project

type ConfigLoader interface {
	LoadConfig() (Config, error)
}

type FileSystemConfigLoader struct {
	fs infra.FileSystem
}

func NewFileSystemConfigLoader(fs infra.FileSystem) *FileSystemConfigLoader {
	return &FileSystemConfigLoader{fs: fs}
}

func (l *FileSystemConfigLoader) getConfigPath() (string, error) {
	homeDir, err := l.fs.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "lazyalias", "config.yaml"), nil
}

func (l *FileSystemConfigLoader) LoadConfig() (Config, error) {
	configPath, err := l.getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := l.fs.ReadFile(configPath)
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
