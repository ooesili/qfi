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

	// figure out name and destination
	var name, destination string
	if len(args) == 1 {
		destination = args[0]
		name = filepath.Base(destination)
	} else {
		name = args[0]
		destination = args[1]
	}

	return a.Driver.Add(name, destination)
}
