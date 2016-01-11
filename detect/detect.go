// Package detect inspects and summarizes file types and permissions.
package detect

import "os"

// Type represents an abstract file type and permission combination.
type Type int

const (
	NormalFile Type = iota
	UnwritableFile
	InaccessibleFile
	NonexistentFile

	NormalDirectory
	UnreadableDirectory

	UnknownFile
)

// Detector holds the Detect method so that it can be passed into interfaces.
type Detector struct{}

// Detect returns a Type based on the given path.
func (Detector) Detect(path string) Type {
	stat, err := os.Stat(path)
	if err != nil {
		return handleStatErr(err)
	}
	if stat.IsDir() {
		return handleDirectory(path)
	}
	return handleFile(path)
}

func handleStatErr(err error) Type {
	if os.IsPermission(err) {
		return InaccessibleFile
	}
	if os.IsNotExist(err) {
		return NonexistentFile
	}
	return UnknownFile
}

func handleDirectory(path string) Type {
	if fi, err := os.Open(path); err == nil {
		fi.Close()
		return NormalDirectory
	}
	return UnreadableDirectory
}

func handleFile(path string) Type {
	if isNormalFile(path) {
		return NormalFile
	}
	return UnwritableFile
}

func isNormalFile(path string) bool {
	if fi, err := os.OpenFile(path, os.O_WRONLY, 0); err == nil {
		fi.Close()
		return true
	}
	return false
}
