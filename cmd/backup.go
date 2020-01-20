package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Backup() *cli.Command {
	return &cli.Command{
		Name:   "backup",
		Usage:  "Backup hosts file",
		Action: backup,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "File location to store backup, default <hostsdir>/.hosts",
			},
		},
	}
}

func backup(c *cli.Context) error {
	fmt.Println("todo")
	return debugFooter(c)
}
