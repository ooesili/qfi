package commands

import "errors"

type RenameDriver interface {
	Rename(name, newName string) error
}

type Rename struct {
	Driver RenameDriver
}

func (m Rename) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no target specified")
	}
	if len(args) == 1 {
		return errors.New("new name not specified")
	}
	if len(args) > 2 {
		return errors.New("too many arguments")
	}

	return m.Driver.Rename(args[0], args[1])
}
