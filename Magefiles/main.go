// Package main provides programmer automations for the project.
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/sh"
)

// init performs some sanity checks before running anything.
func init() {
	mustBeInRoot()
	mustHaveInstalled("docker")
}

// Test tests the whole repo using Ginkgo test runner.
func Test() error {
	if err := sh.Run(
		"go", "run", "-mod=readonly", "github.com/onsi/ginkgo/v2/ginkgo",
		"-p", "-randomize-all", "-repeat=5", "--fail-on-pending", "--race", "--trace",
		"--junit-report=test-report.xml",
		"./...",
	); err != nil {
		return fmt.Errorf("failed to run ginkgo: %w", err)
	}

	return nil
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
