// Package summary prints a pretty summary of all targets and their
// destinations.
package summarize

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/fatih/color"
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
	targets := s.Resolver.List()
	width := strconv.Itoa(lengthOfLongestString(targets))

	result := &bytes.Buffer{}
	for _, target := range targets {
		line := s.getSummaryLine(target, width)
		fmt.Fprintln(result, line)
	}

	return result.String()
}

func (s Summarizer) getSummaryLine(target string, width string) string {
	destination, err := s.Resolver.Resolve(target)
	if err != nil {
		panic("summary: List/Resolve mismatch, cannot resolve: " + target)
	}
	arrow, colorFunc := s.getSummaryLineStyles(destination)
	return fmt.Sprintf(colorFunc("%-"+width+"s %c> %s"),
		target, arrow, destination)
}

func (s Summarizer) getSummaryLineStyles(path string) (rune, func(a ...interface{}) string) {
	fileType := s.Detector.Detect(path)
	arrow := arrowCharForFileType(fileType)
	colorFunc := colorForFileType(fileType)
	return arrow, colorFunc
}

func lengthOfLongestString(strs []string) int {
	width := 0
	for _, str := range strs {
		if len(str) > width {
			width = len(str)
		}
	}
	return width
}

func arrowCharForFileType(fileType detect.Type) rune {
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

func colorForFileType(fileType detect.Type) func(a ...interface{}) string {
	var fgColor color.Attribute

	switch fileType {
	case detect.NormalFile:
		fgColor = color.FgGreen
	case detect.UnwritableFile:
		fgColor = color.FgYellow
	case detect.NormalDirectory:
		fgColor = color.FgBlue
	case detect.UnreadableDirectory, detect.InaccessibleFile, detect.NonexistentFile:
		fgColor = color.FgRed
	default:
		fgColor = color.FgMagenta
	}

	return color.New(fgColor).SprintFunc()
}
