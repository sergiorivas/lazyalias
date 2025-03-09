package config

import (
	"os"
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

func (m *MockFileSystem) ReadDir(path string) ([]os.DirEntry, error) {
	args := m.Called(path)
	return args.Get(0).([]os.DirEntry), args.Error(1)
}

type MockDirEntry struct {
	name  string
	isDir bool
}

func (m MockDirEntry) Name() string {
	return m.name
}

func (m MockDirEntry) IsDir() bool {
	return m.isDir
}

func (m MockDirEntry) Type() os.FileMode {
	return 0
}

func (m MockDirEntry) Info() (os.FileInfo, error) {
	return nil, nil
}

func TestLoadConfig(t *testing.T) {
	mockFS := new(MockFileSystem)

	mockEntries := []os.DirEntry{
		MockDirEntry{name: "config1.yaml", isDir: false},
		MockDirEntry{name: "config2.yaml", isDir: false},
		MockDirEntry{name: "ignoreme.txt", isDir: false},
		MockDirEntry{name: "somedir", isDir: true},
	}

	validYaml1 := []byte(`
project1:
  name: "Test Project 1"
  commands:
    - name: "Test Command 1"
      command: "echo hello"
`)

	validYaml2 := []byte(`
project2:
  name: "Test Project 2"
  commands:
    - name: "Test Command 2"
      command: "echo world"
`)

	mockFS.On("UserHomeDir").Return("/home/user", nil)
	mockFS.On("ReadDir", "/home/user/.config/lazyalias").Return(mockEntries, nil)
	mockFS.On("ReadFile", "/home/user/.config/lazyalias/config1.yaml").Return(validYaml1, nil)
	mockFS.On("ReadFile", "/home/user/.config/lazyalias/config2.yaml").Return(validYaml2, nil)

	loader := NewFileSystemConfigLoader(mockFS)

	config, err := loader.LoadConfig()

	assert.NoError(t, err)
	assert.NotNil(t, config)

	assert.Equal(t, "Test Project 1", config["project1"].Name)
	assert.Equal(t, "Test Command 1", config["project1"].Commands[0].Name)
	assert.Equal(t, "echo hello", config["project1"].Commands[0].Command)

	assert.Equal(t, "Test Project 2", config["project2"].Name)
	assert.Equal(t, "Test Command 2", config["project2"].Commands[0].Name)
	assert.Equal(t, "echo world", config["project2"].Commands[0].Command)

	mockFS.AssertExpectations(t)
}
