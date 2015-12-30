package edit

import (
	"os"
	"os/exec"
)

// RealExecutor implements Executor and runs actual external commands.
type RealExecutor struct {
}

// Exec runs an external command.
func (RealExecutor) Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
