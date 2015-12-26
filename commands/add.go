package commands

import "errors"

type AddDriver interface {
	Add(name, destination string) error
}

type Add struct {
	Driver AddDriver
}

func (a *Add) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no target specified")
	}
	if len(args) == 1 {
		return errors.New("no file specified")
	}
	if len(args) > 2 {
		return errors.New("too many arguments")
	}

	return a.Driver.Add(args[0], args[1])
}
