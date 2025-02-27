package infra

import "os"

type FileSystem interface {
	ReadFile(path string) ([]byte, error)
	Getwd() (string, error)
	UserHomeDir() (string, error)
}

type OSFileSystem struct{}

func (fs *OSFileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fs *OSFileSystem) Getwd() (string, error) {
	return os.Getwd()
}

func (fs *OSFileSystem) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func NewOSFileSystem() *OSFileSystem {
	return &OSFileSystem{}
}
