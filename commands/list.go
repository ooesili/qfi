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
	if len(args) == 0 {
		return l.printAllTargets()
	}
	if len(args) == 1 {
		return l.printTargetDestination(args[0])
	}

	return ErrTooManyArgs
}

func (l List) printAllTargets() error {
	for _, target := range l.Driver.List() {
		fmt.Fprintln(l.Logger, target)
	}
	return nil
}

func (l List) printTargetDestination(name string) error {
	destination, err := l.Driver.Resolve(name)
	if err != nil {
		return err
	}
	fmt.Fprintln(l.Logger, destination)
	return nil
}
