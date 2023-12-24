package install

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/uwu-tools/magex/pkg"
)

const (
	GoreleaserVersion = "v1.22.1"
)

// Dependencies install all dependencies
func Dependencies() error {
	fmt.Println("Installing Dependencies...")
	mg.Deps(Golangcilint, Goreleaser)

	return nil
}

// Golangcilint install golangci-lint
func Golangcilint() error {
	fmt.Println("Installing GolangCI Lint...")
	opts := pkg.EnsurePackageOptions{
		Name:           "github.com/golangci/golangci-lint/cmd/golangci-lint",
		DefaultVersion: "v1.54.2",
		VersionCommand: "version",
		Destination:    "bin",
	}
	return pkg.EnsurePackageWith(opts)
}

// Goreleaser install goreleaser
func Goreleaser() error {
	fmt.Println("Installing goreleaser...")
	opts := pkg.EnsurePackageOptions{
		Name:           "github.com/goreleaser/goreleaser",
		DefaultVersion: "v1.22.1",
		VersionCommand: "--version",
		Destination:    "bin",
	}
	return pkg.EnsurePackageWith(opts)
}
