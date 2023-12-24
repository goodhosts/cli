//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/uwu-tools/magex/xplat"

	//mage:import install
	"github.com/goodhosts/cli/mage/install"

	//mage:import test
	"github.com/goodhosts/cli/mage/test"
)

func init() {
	cwd, _ := os.Getwd()
	xplat.EnsureInPath(filepath.Join(cwd, "bin"))
}

// run everything for ci process (install deps, lint, coverage, build)
func Ci() error {
	fmt.Println("Running Continuous Integration...")
	mg.Deps(install.Dependencies)
	mg.Deps(Lint, test.Coverage)
	mg.Deps(test.Build)
	return nil
}

// run the linter
func Lint() error {
	mg.Deps(install.Golangcilint)
	fmt.Println("Running Linter...")
	return sh.RunV("./bin/golangci-lint", "run")
}

// delete files and paths made by mage
func Clean() error {
	for _, path := range []string{"coverage.txt", "dist", "bin"} {
		if err := sh.Rm(path); err != nil {
			return err
		}
	}

	return nil
}
