package commands

import (
	"errors"
	"fmt"
	"io"
)

type SummaryDriver interface {
	Summary() string
}

type Summary struct {
	Driver SummaryDriver
	Logger io.Writer
}

func (s Summary) Run(args []string) error {
	if len(args) != 0 {
		return errors.New("too many arguments")
	}

	fmt.Fprint(s.Logger, s.Driver.Summary())
	return nil
}
