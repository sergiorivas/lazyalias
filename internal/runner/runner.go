package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sergiorivas/lazyalias/internal/types"
)

type ExecutionContext struct {
	OriginalDir string
	TargetDir   string
	Command     types.Command
	Project     types.Project
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
	var finalCommand string

	if ctx.TargetDir != "" && ctx.TargetDir != r.currentDir {
		finalCommand += fmt.Sprintf("cd %s && ", shellescape(ctx.TargetDir))
	}

	command := ctx.Command.Command
	if len(ctx.Command.Args) > 0 {
		for i, arg := range ctx.Command.Args {
			if arg.Value == "" {
				continue
			}
			placeholder := fmt.Sprintf("$arg_%d", i+1)
			command = strings.ReplaceAll(command, placeholder, arg.Value)
		}
	}

	finalCommand += command
	return finalCommand
}

// shellescape escapes a string to be used in shell
func shellescape(s string) string {
	return fmt.Sprintf("'%s'", filepath.ToSlash(s))
}
