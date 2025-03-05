package infra

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type HistoryFileUpdater struct {
	Type     string
	HistFile string
}

type HistoryEntry struct {
	Command   string
	Timestamp int64
}

func detectShellType() string {
	if shell := os.Getenv("SHELL"); shell != "" {
		return filepath.Base(shell)
	}
	return ""
}

func getHistoryFilePath(shellType string) (string, error) {
	if histFile := os.Getenv("HISTFILE"); histFile != "" {
		return histFile, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("no se pudo obtener el directorio home: %v", err)
	}

	switch shellType {
	case "bash":
		return filepath.Join(homeDir, ".bash_history"), nil
	case "zsh":
		return filepath.Join(homeDir, ".zsh_history"), nil
	case "sh":
		return filepath.Join(homeDir, ".sh_history"), nil
	default:
		return "", fmt.Errorf("shell no soportado: %s", shellType)
	}
}

func NewHistoryFileUpdater() (*HistoryFileUpdater, error) {
	shellType := detectShellType()
	histFile, err := getHistoryFilePath(shellType)
	if err != nil {
		return nil, err
	}

	return &HistoryFileUpdater{
		Type:     shellType,
		HistFile: histFile,
	}, nil
}

func (s *HistoryFileUpdater) formatEntry(entry HistoryEntry) string {
	switch s.Type {
	case "zsh":
		return fmt.Sprintf(": %d:0;%s\n", entry.Timestamp, entry.Command)
	case "bash", "sh":
		return entry.Command + "\n"
	default:
		return entry.Command + "\n"
	}
}

func (s *HistoryFileUpdater) writeToHistoryFile(entry HistoryEntry) error {
	f, err := os.OpenFile(s.HistFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	formattedEntry := s.formatEntry(entry)
	_, err = f.WriteString(formattedEntry)
	return err
}

func (s *HistoryFileUpdater) Add(command string) error {
	entry := HistoryEntry{
		Command:   command,
		Timestamp: time.Now().Unix(),
	}

	if err := s.writeToHistoryFile(entry); err != nil {
		return fmt.Errorf("unable to save the history: %v", err)
	}

	return nil
}
