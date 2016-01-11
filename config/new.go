package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// New creates a new Config by reading symlinks from the given directory.
func New(configDir string) (*Config, error) {
	if err := ensureDirectory(configDir); err != nil {
		return nil, err
	}
	targets, err := getTargets(configDir)
	if err != nil {
		return nil, err
	}
	return &Config{
		configDir: configDir,
		targets:   targets,
	}, nil
}

func ensureDirectory(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("cannot create directory: %s: %s", path, err)
	}
	return nil
}

func getTargets(configDir string) (map[string]string, error) {
	names, err := listFiles(configDir)
	if err != nil {
		return nil, err
	}
	targets := make(map[string]string)
	for _, name := range names {
		destination, err := resolveLink(filepath.Join(configDir, name))
		if err != nil {
			return nil, err
		}
		targets[name] = destination
	}
	return targets, nil
}

func listFiles(path string) ([]string, error) {
	dir, err := os.Open(path)
	if err != nil {
		pathErr := err.(*os.PathError).Err
		return nil, fmt.Errorf("cannot open directory: %s: %s", path, pathErr)
	}
	names, err := dir.Readdirnames(0)
	if err != nil {
		return nil, fmt.Errorf("cannot read directory: %s: %s", path, err)
	}
	return names, nil
}

func resolveLink(path string) (string, error) {
	destination, err := os.Readlink(path)
	if err != nil {
		// EINVAL means argument was not a symlink
		pathErr := err.(*os.PathError).Err
		if pathErr.Error() == "invalid argument" {
			return "", fmt.Errorf("not a symlink: %s", path)
		}
		return "", fmt.Errorf("error following symlink: %s", err)
	}
	return destination, nil
}
