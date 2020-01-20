package cmd

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
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
				Aliases: []string{"i"},
				Usage:   "File location to restore a backup from, default <hostsdir>/.hosts",
			},
		},
	}
}

func restore(c *cli.Context) error {
	hostsfile, err := loadHostsfile(c)
	if err != nil {
		// debug only, no problem if file doesn't exist we just need path
		logrus.Debugf("destination hosts file not found: %s", hostsfile.Path)
	}

	input := c.String("input")
	if input == "" {
		input = filepath.Join(
			filepath.Dir(hostsfile.Path),
			"."+filepath.Base(hostsfile.Path))
	}

	_, err = copyFile(input, hostsfile.Path)
	if err != nil {
		return err
	}

	return debugFooter(c)
}
