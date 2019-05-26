package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func List() cli.Command {
	return cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "List all entries in the hostsfile",
		Action:  list,
		Flags: []cli.Flag{
			cli.BoolFlag{
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

	total := 0
	for _, line := range hostsfile.Lines {
		var lineOutput string

		if line.IsComment() && !c.Bool("all") {
			continue
		}

		if line.Raw == "" && !c.Bool("all") {
			continue
		}

		lineOutput = fmt.Sprintf("%s", line.Raw)
		if line.Err != nil {
			lineOutput = fmt.Sprintf("%s # <<< Malformated!", lineOutput)
		}
		total += 1

		fmt.Println(lineOutput)
	}

	fmt.Printf("Total: %d", total)

	return nil
}
