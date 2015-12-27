package detect

import "os"

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

func Detect(path string) Type {
	// try to stat file
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsPermission(err) {
			return InaccessibleFile
		}
		if os.IsNotExist(err) {
			return NonexistentFile
		}
		return UnknownFile
	}

	if stat.IsDir() {
		// see if we can open the directory
		if fi, err := os.Open(path); err == nil {
			fi.Close()
			return NormalDirectory
		}
		return UnreadableDirectory
	}

	// normal, writable file
	if fi, err := os.OpenFile(path, os.O_WRONLY, 0); err == nil {
		fi.Close()
		return NormalFile
	}

	// cannot open file for writing
	return UnwritableFile
}
