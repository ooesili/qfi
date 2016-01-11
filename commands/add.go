package commands

import (
	"path/filepath"

	"github.com/ooesili/qfi/dispatch"
)

var ErrNoTargetOrFile = dispatch.UsageError{"no target or file specified"}

type AddDriver interface {
	Add(name, destination string) error
}

type Add struct {
	Driver AddDriver
}

func (a *Add) Run(args []string) error {
	if len(args) == 0 {
		return ErrNoTargetOrFile
	}
	if len(args) > 2 {
		return ErrTooManyArgs
	}

	return a.Driver.Add(a.parseArgs(args))
}

func (Add) parseArgs(args []string) (name string, destination string) {
	if len(args) == 1 {
		return filepath.Base(args[0]), args[0]
	}
	return args[0], args[1]
}
