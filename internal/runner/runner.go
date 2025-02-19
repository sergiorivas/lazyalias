package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sergiorivas/lazyalias/internal/config"
)

type ExecutionContext struct {
	OriginalDir string
	TargetDir   string
	Command     config.Command
	Project     config.Project
}

type Runner struct {
	currentDir string
}

func NewRunner() (*Runner, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &Runner{currentDir: currentDir}, nil
}

func (r *Runner) PrepareCommand(ctx ExecutionContext) string {
	var command string

	if ctx.TargetDir != "" && ctx.TargetDir != r.currentDir {
		command += fmt.Sprintf("cd %s && ", shellescape(ctx.TargetDir))
	}

	command += ctx.Command.Command
	return command
}

func (r *Runner) CopyToClipboard(command string) error {
	switch runtime.GOOS {
	case "darwin":
		return copyOSX(command)
	case "linux":
		return copyLinux(command)
	default:
		return fmt.Errorf("clipboard copy not supported on %s", runtime.GOOS)
	}
}

func copyOSX(command string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(command)
	return cmd.Run()
}

func copyLinux(command string) error {
	// Try xclip first
	if hasCommand("xclip") {
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(command)
		return cmd.Run()
	}

	// Fallback to xsel
	if hasCommand("xsel") {
		cmd := exec.Command("xsel", "--clipboard", "--input")
		cmd.Stdin = strings.NewReader(command)
		return cmd.Run()
	}

	return fmt.Errorf("neither xclip nor xsel found. Please install one of them")
}

func hasCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// shellescape escapes a string to be used in shell
func shellescape(s string) string {
	return fmt.Sprintf("'%s'", filepath.ToSlash(s))
}
