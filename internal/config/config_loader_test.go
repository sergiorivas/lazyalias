package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileSystem struct {
	mock.Mock
}

func (m *MockFileSystem) Getwd() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockFileSystem) UserHomeDir() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockFileSystem) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func TestLoadConfig(t *testing.T) {
	mockFS := new(MockFileSystem)
	validYaml := []byte(`
project1:
  name: "Test Project"
  commands:
    - name: "Test Command"
      command: "echo hello"
`)
	mockFS.On("UserHomeDir").Return("/home/user", nil)
	mockFS.On("ReadFile", "/home/user/.config/lazyalias/config.yaml").Return(validYaml, nil)

	loader := NewFileSystemConfigLoader(mockFS)

	config, err := loader.LoadConfig()

	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "Test Project", config["project1"].Name)
	assert.Equal(t, "Test Command", config["project1"].Commands[0].Name)
	assert.Equal(t, "echo hello", config["project1"].Commands[0].Command)
}
