package commands

type MoveDriver interface {
	Move(name, destination string) error
}

type Move struct {
	Driver MoveDriver
}

func (m Move) Run(args []string) error {
	if len(args) == 0 {
		return ErrNoTarget
	}
	if len(args) == 1 {
		return ErrNoFile
	}
	if len(args) > 2 {
		return ErrTooManyArgs
	}

	return m.Driver.Move(args[0], args[1])
}
