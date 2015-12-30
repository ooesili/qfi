package commands

import "github.com/ooesili/qfi/dispatch"

var (
	ErrNoFile      = dispatch.UsageError{"no file specified"}
	ErrNoNewName   = dispatch.UsageError{"new name not specified"}
	ErrNoTarget    = dispatch.UsageError{"no target specified"}
	ErrNoTargets   = dispatch.UsageError{"no targets specified"}
	ErrTooManyArgs = dispatch.UsageError{"too many arguments"}
)
