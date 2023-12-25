package main

import (
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

func run(args []string) error {
	cmd.Version(version, commit, date)
	return cmd.App.Run(args)
}
