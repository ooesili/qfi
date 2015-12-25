package commands

import "errors"

type Config interface {
	Add(name, destination string) error
}

func New(config Config) Commands {
	return Commands{config}
}

type Commands struct {
	config Config
}

func (c Commands) Add(args []string) error {
	// check number of arguments
	if len(args) == 0 {
		return errors.New("no target specified")
	}
	if len(args) == 1 {
		return errors.New("no file specified")
	}
	if len(args) > 2 {
		return errors.New("too many arguments")
	}

	return c.config.Add(args[0], args[1])
}
