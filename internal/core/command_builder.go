package core

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

type CommandBuilder struct {
	currentDir string
}

func shellescape(s string) string {
	return fmt.Sprintf("'%s'", filepath.ToSlash(s))
}

func NewCommandBuilder() (*CommandBuilder, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &CommandBuilder{currentDir: currentDir}, nil
}

func (r *CommandBuilder) Build(ctx *ExecutionContext) string {
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
