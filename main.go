package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/goodhosts/cli/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func formatVersion(version, commit, date string) string {
	return fmt.Sprintf("goodhosts %s@%s built on %s", version, commit, date)
}

func run(args []string) error {
	cmd.Version(formatVersion(version, commit, date))
	return cmd.App.Run(args)
}
