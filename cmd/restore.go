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
	hf, err := loadHostsfile(c, false)
	if err != nil {
		// debug only, no problem if file doesn't exist we just need path
		logrus.Debugf("destination hosts file not found: %s", hf.Path)
	}

	input := c.String("input")
	if input == "" {
		input = filepath.Join(
			filepath.Dir(hf.Path),
			"."+filepath.Base(hf.Path))
	}

	if err := copyFile(input, hf.Path); err != nil {
		return err
	}

	return debugFooter(c)
}
