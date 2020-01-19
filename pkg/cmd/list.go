package cmd

import (
	"fmt"

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
				Name:  "all",
				Usage: "Show all entries in the hosts file including commented lines.",
			},
		},
	}
}

func list(c *cli.Context) error {
	hostsfile, err := loadHostsfile(c)
	if err != nil {
		return err
	}

	for _, line := range hostsfile.Lines {
		if !c.Bool("all") {
			if line.IsComment() || line.Raw == "" {
				continue
			}
		}

		lineOutput := fmt.Sprintf("%s", line.Raw)
		if line.IsMalformed() {
			lineOutput = fmt.Sprintf("%s # <<< Malformed!", lineOutput)
		}

		fmt.Println(lineOutput)
	}

	return debugFooter(c)
}
