package infra

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type CommandRunner interface {
	Run(name string, args ...string) error
	LookPath(name string) (string, error)
	SetText(text string)
}

type RealCommandRunner struct {
	text string
}

type ClipboardOption func(*clipboard)

func (r *RealCommandRunner) Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = strings.NewReader(r.text)
	return cmd.Run()
}

func (r *RealCommandRunner) LookPath(name string) (string, error) {
	return exec.LookPath(name)
}

func (r *RealCommandRunner) SetText(text string) {
	r.text = text
}

type OSDetector interface {
	GetOS() string
}

type RealOSDetector struct{}

func (d *RealOSDetector) GetOS() string {
	return runtime.GOOS
}

type Clipboard interface {
	Copy(text string) error
}

type clipboard struct {
	runner     CommandRunner
	osDetector OSDetector
}

func (c *clipboard) copyOSX() error {
	return c.runner.Run("pbcopy")
}

func (c *clipboard) copyLinux() error {
	// Try xclip first
	if c.hasCommand("xclip") {
		return c.runner.Run("xclip", "-selection", "clipboard")
	}

	// Fallback to xsel
	if c.hasCommand("xsel") {
		return c.runner.Run("xsel", "--clipboard", "--input")
	}

	return fmt.Errorf("neither xclip nor xsel found. Please install one of them")
}

func (c *clipboard) hasCommand(cmd string) bool {
	_, err := c.runner.LookPath(cmd)
	return err == nil
}

func WithCommandRunner(runner CommandRunner) ClipboardOption {
	return func(c *clipboard) {
		if runner != nil {
			c.runner = runner
		}
	}
}

func WithOSDetector(detector OSDetector) ClipboardOption {
	return func(c *clipboard) {
		if detector != nil {
			c.osDetector = detector
		}
	}
}

func NewClipboard(opts ...ClipboardOption) Clipboard {
	c := &clipboard{
		runner:     &RealCommandRunner{},
		osDetector: &RealOSDetector{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
func (c *clipboard) Copy(text string) error {
	c.runner.SetText(text)
	switch c.osDetector.GetOS() {
	case "darwin":
		return c.copyOSX()
	case "linux":
		return c.copyLinux()
	default:
		return fmt.Errorf("clipboard copy not supported on %s", c.osDetector.GetOS())
	}
}
