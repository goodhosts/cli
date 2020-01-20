package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Restore() *cli.Command {
	return &cli.Command{
		Name:   "restore",
		Usage:  "Restore hosts file from backup",
		Action: restore,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"o"},
				Usage:   "File location to restore a backup from, default <hostsdir>/.hosts",
			},
		},
	}
}

func restore(c *cli.Context) error {
	fmt.Println("todo")
	return debugFooter(c)
}
