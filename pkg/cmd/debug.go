package cmd

import (
	"github.com/urfave/cli/v2"
)

func Debug() *cli.Command {
	return &cli.Command{
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "Show debug table for hosts file",
		Action:  debug,
	}
}

func debug(c *cli.Context) error {
	return debugFooter(c)
}
