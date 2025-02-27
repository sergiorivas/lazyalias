package config

import (
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/sergiorivas/lazyalias/internal/infra"
	"github.com/sergiorivas/lazyalias/internal/types"
)

type Config map[string]types.Project

type ConfigLoader interface {
	LoadConfig(path string) (Config, error)
}

type FileSystemConfigLoader struct {
	fs infra.FileSystem
}

func NewFileSystemConfigLoader(fs infra.FileSystem) *FileSystemConfigLoader {
	return &FileSystemConfigLoader{fs: fs}
}

func (l *FileSystemConfigLoader) LoadConfig(path string) (Config, error) {
	data, err := l.fs.ReadFile(path)
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
