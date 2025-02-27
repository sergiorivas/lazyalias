// internal/runner/runner_test.go
package runner

import (
	"testing"

	"github.com/sergiorivas/lazyalias/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestPrepareCommand(t *testing.T) {
	r := &Runner{currentDir: "/current/dir"}
	ctx := ExecutionContext{
		TargetDir: "/target/dir",
		Command: types.Command{
			Command: "echo $arg_1",
			Args: []types.Arg{
				{Value: "hello"},
			},
		},
	}

	result := r.PrepareCommand(ctx)

	expected := "cd '/target/dir' && echo hello"
	assert.Equal(t, expected, result)
}
