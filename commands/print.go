package commands

import (
	"fmt"
	"io"
)

type Print struct {
	Text   string
	Logger io.Writer
}

func (p Print) Run(args []string) error {
	if len(args) != 0 {
		return ErrTooManyArgs
	}

	fmt.Fprint(p.Logger, p.Text)
	return nil
}
