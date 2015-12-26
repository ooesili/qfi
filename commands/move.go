package commands

import "errors"

type MoveDriver interface {
	Move(name, destination string) error
}

type Move struct {
	Driver MoveDriver
}

func (m Move) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no target specified")
	}
	if len(args) == 1 {
		return errors.New("no file specified")
	}
	if len(args) > 2 {
		return errors.New("too many arguments")
	}

	return m.Driver.Move(args[0], args[1])
}
