package commands

import (
	"fmt"
	"io"
)

type ListDriver interface {
	List() []string
	Resolve(name string) (string, error)
}

type List struct {
	Driver ListDriver
	Logger io.Writer
}

func (l List) Run(args []string) error {
	// print all targets
	if len(args) == 0 {
		for _, target := range l.Driver.List() {
			fmt.Fprintln(l.Logger, target)
		}
		return nil
	}

	// resolve a single target
	if len(args) == 1 {
		destination, err := l.Driver.Resolve(args[0])
		if err != nil {
			return err
		}

		fmt.Fprintln(l.Logger, destination)
		return nil
	}

	return ErrTooManyArgs
}
