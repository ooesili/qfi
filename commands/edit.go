package commands

type EditDriver interface {
	Edit(name string) error
}

type Edit struct {
	Driver EditDriver
}

func (e Edit) Run(args []string) error {
	if len(args) == 0 {
		return ErrNoTarget
	}
	if len(args) > 1 {
		return ErrTooManyArgs
	}

	return e.Driver.Edit(args[0])
}
