// Package config reads target information from a config directory.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// New creates a new Config by reading symlinks from the given directory.
func New(configDir string) (*Config, error) {
	targets := make(map[string]string)

	// open config directory
	dir, err := os.Open(configDir)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot open directory: %s: %s",
			configDir, err.(*os.PathError).Err,
		)
	}

	// look at all file names in it
	names, err := dir.Readdirnames(0)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot read directory: %s: %s",
			configDir, err,
		)
	}

	// read each link and store result in targets
	for _, name := range names {
		fullName := filepath.Join(configDir, name)

		// read symlink
		destination, err := os.Readlink(fullName)
		if err != nil {
			// EINVAL means argument was not a symlink
			if err.(*os.PathError).Err.Error() == "invalid argument" {
				return nil, fmt.Errorf("not a symlink: %s", fullName)
			}
			return nil, fmt.Errorf("error following symlink: %s", err)
		}

		targets[name] = destination
	}

	return &Config{
		configDir: configDir,
		targets:   targets,
	}, nil
}

// Config stores a map from target names to their destinations.
type Config struct {
	configDir string
	targets   map[string]string
}

// Lookup finds a target by name and returns its destination.
func (c Config) Resolve(name string) string {
	return c.targets[name]
}

// Add adds a new target to the config directory.
func (c Config) Add(name, destination string) error {
	targetFile := filepath.Join(c.configDir, name)

	// find absolute path
	absDestination, err := filepath.Abs(destination)
	if err != nil {
		return fmt.Errorf("cannot cannonicalize path: %s: %s", destination, err)
	}

	// create symlink
	err = os.Symlink(absDestination, targetFile)
	if err != nil {
		return fmt.Errorf("cannot create symlink: %s: %s",
			targetFile, err.(*os.LinkError).Err)
	}

	return nil
}

// List returns all target names in alphabetical order.
func (c Config) List() []string {
	result := make([]string, len(c.targets))

	// add each key
	i := 0
	for k := range c.targets {
		result[i] = k
		i++
	}

	// sort result
	sort.Strings(result)

	return result
}

// Delete removes a target from the config directory.
func (c Config) Delete(name string) error {
	// make sure target exists
	_, ok := c.targets[name]
	if !ok {
		return fmt.Errorf("target '%s' does not exist", name)
	}

	// remove target
	err := os.Remove(filepath.Join(c.configDir, name))
	if err != nil {
		return fmt.Errorf("cannot remove target: %s: %s",
			name, err.(*os.PathError).Err)
	}

	return nil
}
