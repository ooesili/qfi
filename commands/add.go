package commands

type AddDriver interface {
	Add(name, destination string) error
}

type Add struct {
	Driver AddDriver
}

func (a *Add) Run(args []string) error {
	if len(args) == 0 {
		return ErrNoTarget
	}
	if len(args) == 1 {
		return ErrNoFile
	}
	if len(args) > 2 {
		return ErrTooManyArgs
	}

	return a.Driver.Add(args[0], args[1])
}
