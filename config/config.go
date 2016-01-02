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
	// make sure config directory exists
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot create directory: %s: %s", configDir, err)
	}

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

// targetFile returns a target's link file.
func (c Config) targetFile(name string) string {
	return filepath.Join(c.configDir, name)
}

// Lookup finds a target by name and returns its destination.
func (c Config) Resolve(name string) (string, error) {
	// validate target name
	if err := validateNames(name); err != nil {
		return "", err
	}

	// make sure target exists
	if destination, ok := c.targets[name]; ok {
		return destination, nil
	}

	return "", fmt.Errorf("target '%s' does not exist", name)
}

// Add adds a new target to the config directory.
func (c Config) Add(name, destination string) error {
	// validate target name
	if err := validateNames(name); err != nil {
		return err
	}

	targetFile := c.targetFile(name)

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

// Delete removes the given targets from the config directory.
func (c Config) Delete(names ...string) error {
	// validate names
	if err := validateNames(names...); err != nil {
		return err
	}

	// make sure each target exists
	for _, name := range names {
		if _, ok := c.targets[name]; !ok {
			return fmt.Errorf("target '%s' does not exist", name)
		}
	}

	// remove each target
	for _, name := range names {
		err := os.Remove(c.targetFile(name))
		if err != nil {
			return fmt.Errorf("cannot remove target: %s: %s",
				name, err.(*os.PathError).Err)
		}
	}

	return nil
}

// Move changes the destination of a target.
func (c Config) Move(name, destination string) error {
	// validate target name
	if err := validateNames(name); err != nil {
		return err
	}

	// remove old target
	err := c.Delete(name)
	if err != nil {
		return err
	}

	// create target with new destination
	return c.Add(name, destination)
}

// Rename changes a taget's name while leaving its destination the same.
func (c Config) Rename(name, newName string) error {
	// validate both target names
	if err := validateNames(name, newName); err != nil {
		return err
	}

	// make sure old target exists
	if _, ok := c.targets[name]; !ok {
		return fmt.Errorf("target '%s' does not exist", name)
	}

	// make sure new target does not exist
	if _, ok := c.targets[newName]; ok {
		return fmt.Errorf("target '%s' exists", newName)
	}

	// resolve target names
	targetFile := c.targetFile(name)
	newTargetFile := c.targetFile(newName)

	// rename target
	err := os.Rename(targetFile, newTargetFile)
	if err != nil {
		return fmt.Errorf("cannot rename file: %s: %s",
			targetFile, err.(*os.LinkError).Err)
	}

	return nil
}

// validateName returns an error if the target name is invalid
func validateNames(names ...string) error {
	for _, name := range names {
		for _, char := range name {
			if char == os.PathSeparator {
				return fmt.Errorf("invalid target name: %s", name)
			}
		}
	}
	return nil
}
