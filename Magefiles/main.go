// Package main provides programmer automations for the project.
package main

import (
	"os"
	"os/exec"
)

// init performs some sanity checks before running anything.
func init() {
	mustBeInRoot()
	mustHaveInstalled("docker")
}

// mustHaveInstalled checks if the following binaries are installed.
func mustHaveInstalled(names ...string) {
	for _, exe := range names {
		if _, err := exec.LookPath(exe); err != nil {
			panic("required binary not installed (not found in PATH): " + exe)
		}
	}
}

// mustBeInRoot checks that the command is run in the project root.
func mustBeInRoot() {
	var err error
	if _, err = os.ReadFile("go.mod"); err != nil {
		panic("must be in project root, couldn't stat go.mod file: " + err.Error())
	}
}
