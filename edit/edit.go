// Package edit runs different editors depending on what file type a target
// points to.
package edit

import (
	"errors"
	"fmt"

	"github.com/ooesili/qfi/detect"
)

// ErrWrapperShouldChdir is returned when the shell wrapper function should
// perform a chdir.
var ErrWrapperShouldChdir = errors.New("wrapper chould chdir")

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
	destination, err := e.Resolver.Resolve(name)
	if err != nil {
		return err
	}
	command, args, err := e.getEditorCommand(destination)
	if err != nil {
		return err
	}
	return e.Executor.Exec(command, args...)
}

func (e Editor) getEditorCommand(path string) (string, []string, error) {
	switch e.Detector.Detect(path) {

	case detect.NormalFile, detect.NonexistentFile:
		return e.NormalEditor, []string{path}, nil

	case detect.UnwritableFile, detect.InaccessibleFile:
		return "sudo", []string{"-e", path}, nil

	case detect.NormalDirectory, detect.UnreadableDirectory:
		return "", nil, ErrWrapperShouldChdir

	default: // UnknownFile or invalid detect.Type
		return "", nil, fmt.Errorf("unknown file type for: %s", path)
	}
}
