package ui

import (
	"github.com/sergiorivas/lazyalias/internal/types"
)

const BackToProject = "back-to-project"

type CliUI interface {
	ShowProjectMenu(projects []types.Project) (types.Project, error)
	ShowCommandMenu(commands []types.Command) (types.Command, error)
	ShowArgMenu(arg types.Arg) (string, error)
}

type UI struct{}

func NewUI() *UI {
	return &UI{}
}
