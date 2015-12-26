package commands

import "errors"

type DeleteDriver interface {
	Delete(names ...string) error
}

type Delete struct {
	Driver DeleteDriver
}

func (d Delete) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no targets specified")
	}

	return d.Driver.Delete(args...)
}
