package infra

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type Clipboard interface {
	Copy(text string) error
}

type clipboard struct{}

func NewClipboard() Clipboard {
	return &clipboard{}
}

func (c *clipboard) Copy(text string) error {
	switch runtime.GOOS {
	case "darwin":
		return copyOSX(text)
	case "linux":
		return copyLinux(text)
	default:
		return fmt.Errorf("clipboard copy not supported on %s", runtime.GOOS)
	}
}

func copyOSX(text string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func copyLinux(text string) error {
	// Try xclip first
	if hasCommand("xclip") {
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(text)
		return cmd.Run()
	}

	// Fallback to xsel
	if hasCommand("xsel") {
		cmd := exec.Command("xsel", "--clipboard", "--input")
		cmd.Stdin = strings.NewReader(text)
		return cmd.Run()
	}

	return fmt.Errorf("neither xclip nor xsel found. Please install one of them")
}

func hasCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
