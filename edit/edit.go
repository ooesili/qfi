// Package edit runs different editors depending on what file type a target
// points to.
package edit

import (
	"fmt"

	"github.com/ooesili/qfi/detect"
)

// WrapperShouldChdirError is returned when the shell wrapper function should
// perform a chdir.
type WrapperShouldChdirError struct{ Destination string }

// Error implements the error interface.
func (e WrapperShouldChdirError) Error() string {
	return fmt.Sprintf("wrapper should chdir: %s", e.Destination)
}

// Detector detects a file's type.
type Detector interface {
	Detect(path string) detect.Type
}

// Resolver resolves a target's destination.
type Resolver interface {
	Resolve(name string) (string, error)
}

// Executor runs an external command.
type Executor interface {
	Exec(name string, args ...string) error
}

// Editor figures out which command to call to edit the destination file.
type Editor struct {
	NormalEditor string

	Detector Detector
	Executor Executor
	Resolver Resolver
}

// Edit uses the Executor to call the appropriate command, or return a
// WrapperShouldChdirError in case the the destination is a directory.
func (e Editor) Edit(name string) error {
	// resolve target
	destination, err := e.Resolver.Resolve(name)
	if err != nil {
		return err
	}

	// figure out which command to use
	var command string
	fileType := e.Detector.Detect(destination)
	switch fileType {
	// files
	case detect.NormalFile:
		command = e.NormalEditor
	case detect.UnwritableFile:
		command = "sudoedit"
	case detect.InaccessibleFile:
		command = "sudoedit"
	case detect.NonexistentFile:
		command = e.NormalEditor

	// directories
	case detect.NormalDirectory:
		return WrapperShouldChdirError{destination}
	case detect.UnreadableDirectory:
		return WrapperShouldChdirError{destination}

	// UnknownFile
	default:
		return fmt.Errorf("unknown file type for: %s", destination)
	}

	// run the editor
	return e.Executor.Exec(command, destination)
}
