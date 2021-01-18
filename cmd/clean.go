package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Clean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"cl"},
		Usage:   "Clean the hostsfile by doing: remove dupe IPs, for each IPs remove dupe hosts and sort, sort all IPs, split hosts per OS limitations",
		Action:  clean,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Dry run only, will output contents of the cleaned hostsfile without writing the changes.",
			},
		},
	}
}

func clean(c *cli.Context) error {
	hostsfile, err := loadHostsfile(c, false)
	if err != nil {
		return err
	}

	hostsfile.Clean()
	if c.Bool("dry-run") {
		logrus.Debugln("performing a dry run, writing output")
		outputHostsfile(hostsfile, true)
		return debugFooter(c)
	}

	if err := hostsfile.Flush(); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return debugFooter(c)
}
