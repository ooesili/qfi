// Package summary prints a pretty summary of all targets and their
// destinations.
package summarize

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/ooesili/qfi/detect"
)

// Resolver lists all targets and resolves individual targets' destinations.
// Every target returned by List must be Resolvable without error.
type Resolver interface {
	List() []string
	Resolve(name string) (string, error)
}

// Detector detects a file's type.
type Detector interface {
	Detect(path string) detect.Type
}

// Summarizer returns a pretty string summarizing targets, their destinations,
// and the type of each destination.
type Summarizer struct {
	Detector Detector
	Resolver Resolver
}

// Summary retuns a the summary of all targets in the Resolver.
func (s Summarizer) Summary() string {
	result := &bytes.Buffer{}
	targets := s.Resolver.List()

	// figure width of first column
	width := 0
	for _, target := range targets {
		if len(target) > width {
			width = len(target)
		}
	}
	widthStr := strconv.Itoa(width)

	for _, target := range targets {
		// resolve target
		destination, err := s.Resolver.Resolve(target)
		if err != nil {
			panic("summary: List/Resolve mismatch, cannot resolve: " + target)
		}

		// find arrow character
		arrow := arrowChar(s.Detector.Detect(destination))

		// add target to the result
		fmt.Fprintf(result, "%-"+widthStr+"s %c> %s\n",
			target, arrow, destination)
	}

	return result.String()
}

// arrowChar returns the arrow character the summary line.
func arrowChar(fileType detect.Type) rune {
	switch fileType {
	case detect.NormalFile, detect.NonexistentFile:
		return '-'
	case detect.UnwritableFile, detect.InaccessibleFile:
		return '#'
	case detect.UnreadableDirectory, detect.NormalDirectory:
		return '/'
	default:
		return '?'
	}
}
