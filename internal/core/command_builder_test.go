package core

import (
	"testing"

	"github.com/sergiorivas/lazyalias/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestPrepareCommand(t *testing.T) {
	r := &CommandBuilder{currentDir: "/current/dir"}
	ctx := ExecutionContext{
		TargetDir: "/target/dir",
		Command: types.Command{
			Command: "echo $arg_1",
			Args: []types.Arg{
				{Value: "hello"},
			},
		},
	}

	result := r.Build(&ctx)

	expected := "cd '/target/dir' && echo hello"
	assert.Equal(t, expected, result)
}
