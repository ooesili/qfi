package commands

import "errors"

type EditDriver interface {
	Edit(name string) error
}

type Edit struct {
	Driver EditDriver
}

func (e Edit) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no target specified")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	return e.Driver.Edit(args[0])
}
