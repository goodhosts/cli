package cmd

import (
	"github.com/urfave/cli/v2"
)

func Clean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the hostsfile by combining duplicate ips, sorting and removing duplicate hostnames",
		Action:  clean,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Dry run only, will output contents of new hostsfile without writing the changes.",
			},
		},
	}
}

func clean(c *cli.Context) error {
	hf, err := loadHostsfile(c)
	if err != nil {
		return err
	}

	hf.Clean()
	if c.Bool("dry-run") {
		outputHostsfile(hf, c.Bool("all"))
	} else {
		hf.Flush()
	}

	return debugFooter(c)
}
