package cmd

import (
	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "List all entries in the hostsfile",
		Action:  list,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Show all entries in the hosts file including commented lines.",
			},
		},
	}
}

func list(c *cli.Context) error {
	hf, err := loadHostsfile(c, true)
	if err != nil {
		return err
	}

	outputHostsfile(hf, c.Bool("all"))
	return debugFooter(c)
}
