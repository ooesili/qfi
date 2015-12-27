// Package status prints a pretty summary of all targets and their statuses.
package status

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

// Status returns a pretty string summarizing all targets, their destinations,
// and the type of each destination.
func Status(resolver Resolver, detector Detector) string {
	result := &bytes.Buffer{}
	targets := resolver.List()

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
		destination, err := resolver.Resolve(target)
		if err != nil {
			panic("status: List/Resolve mismatch, cannot resolve: " + target)
		}

		// find arrow character
		arrow := arrowChar(detector.Detect(destination))

		// add target to the result
		fmt.Fprintf(result, "%-"+widthStr+"s %c> %s\n",
			target, arrow, destination)
	}

	return result.String()
}

// arrowChar returns the arrow character the status line.
func arrowChar(fileType detect.Type) rune {
	switch fileType {
	case detect.NormalFile:
		return '-'
	case detect.UnwritableFile:
		return '#'
	case detect.InaccessibleFile:
		return '#'
	case detect.NonexistentFile:
		return '-'
	case detect.NormalDirectory:
		return '/'
	case detect.UnreadableDirectory:
		return '/'
	default:
		return '?'
	}
}
