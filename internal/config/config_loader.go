package config

import (
	"path/filepath"
	"strings"

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

func (l *FileSystemConfigLoader) getConfigDir() (string, error) {
	homeDir, err := l.fs.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "lazyalias"), nil
}

func (l *FileSystemConfigLoader) LoadConfig() (Config, error) {
	configDir, err := l.getConfigDir()
	if err != nil {
		return nil, err
	}

	files, err := l.fs.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	config := make(Config)

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		filePath := filepath.Join(configDir, file.Name())
		data, err := l.fs.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var fileConfig Config
		if err := yaml.Unmarshal(data, &fileConfig); err != nil {
			return nil, err
		}

		for key, project := range fileConfig {
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
	}

	return config, nil
}
