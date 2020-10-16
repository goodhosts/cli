package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Version() *cli.Command {
	return &cli.Command{
		Name:   "version",
		Usage:  "",
		Action: version,
	}
}

func version(c *cli.Context) error {
	logrus.Infof("goodhosts %s", c.Context.Value("version"))
	return nil
}
