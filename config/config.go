// Package config reads target information from a config directory.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Config stores a map from target names to their destinations.
type Config struct {
	configDir string
	targets   map[string]string
}

func (c Config) targetFile(name string) string {
	return filepath.Join(c.configDir, name)
}

// Lookup finds a target by name and returns its destination.
func (c Config) Resolve(name string) (string, error) {
	if err := ensureValid(name); err != nil {
		return "", err
	}
	if destination, ok := c.targets[name]; ok {
		return destination, nil
	}
	return "", fmt.Errorf("target '%s' does not exist", name)
}

// Add adds a new target to the config directory.
func (c Config) Add(name, destination string) error {
	if err := ensureValid(name); err != nil {
		return err
	}

	targetFile := c.targetFile(name)
	err := addAbsLink(destination, targetFile)
	if err != nil {
		return err
	}

	return nil
}

func addAbsLink(from, to string) error {
	absFrom, err := filepath.Abs(from)
	if err != nil {
		return fmt.Errorf("cannot cannonicalize path: %s: %s", to, err)
	}
	if err := os.Symlink(absFrom, to); err != nil {
		linkErr := err.(*os.LinkError).Err
		return fmt.Errorf("cannot create symlink: %s: %s", to, linkErr)
	}
	return nil
}

// List returns all target names in alphabetical order.
func (c Config) List() []string {
	result := make([]string, len(c.targets))
	i := 0
	for k := range c.targets {
		result[i] = k
		i++
	}

	sort.Strings(result)
	return result
}

// Delete removes the given targets from the config directory.
func (c Config) Delete(names ...string) error {
	if err := ensureValid(names...); err != nil {
		return err
	}
	if err := c.ensureTargetsExist(names); err != nil {
		return err
	}
	if err := c.deleteTargets(names); err != nil {
		return err
	}
	return nil
}

func (c Config) ensureTargetsExist(names []string) error {
	for _, name := range names {
		if err := c.ensureTargetExists(name); err != nil {
			return err
		}
	}
	return nil
}

func (c Config) ensureTargetExists(name string) error {
	if _, ok := c.targets[name]; !ok {
		return fmt.Errorf("target '%s' does not exist", name)
	}
	return nil
}

func (c Config) ensureTargetAbsent(name string) error {
	if _, ok := c.targets[name]; ok {
		return fmt.Errorf("target '%s' exists", name)
	}
	return nil
}

func (c Config) deleteTargets(names []string) error {
	for _, name := range names {
		err := os.Remove(c.targetFile(name))
		if err != nil {
			pathErr := err.(*os.PathError).Err
			return fmt.Errorf("cannot remove target: %s: %s", name, pathErr)
		}
	}
	return nil
}

// Move changes the destination of a target.
func (c Config) Move(name, destination string) error {
	if err := ensureValid(name); err != nil {
		return err
	}
	if err := c.Delete(name); err != nil {
		return err
	}
	return c.Add(name, destination)
}

// Rename changes a taget's name while leaving its destination the same.
func (c Config) Rename(name, newName string) error {
	if err := ensureValid(name, newName); err != nil {
		return err
	}
	if err := c.ensureTargetExists(name); err != nil {
		return err
	}
	if err := c.ensureTargetAbsent(newName); err != nil {
		return err
	}
	return c.renameTarget(name, newName)
}

func (c Config) renameTarget(name, newName string) error {
	targetFile := c.targetFile(name)
	newTargetFile := c.targetFile(newName)

	if err := os.Rename(targetFile, newTargetFile); err != nil {
		linkErr := err.(*os.LinkError).Err
		return fmt.Errorf("cannot rename file: %s: %s", targetFile, linkErr)
	}
	return nil
}

func ensureValid(names ...string) error {
	for _, name := range names {
		for _, char := range name {
			if char == os.PathSeparator {
				return fmt.Errorf("invalid target name: %s", name)
			}
		}
	}
	return nil
}
