package commands

type RenameDriver interface {
	Rename(name, newName string) error
}

type Rename struct {
	Driver RenameDriver
}

func (m Rename) Run(args []string) error {
	if len(args) == 0 {
		return ErrNoTarget
	}
	if len(args) == 1 {
		return ErrNoNewName
	}
	if len(args) > 2 {
		return ErrTooManyArgs
	}

	return m.Driver.Rename(args[0], args[1])
}
