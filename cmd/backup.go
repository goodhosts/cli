package cmd

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
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
	hf, err := loadHostsfile(c, false)
	if err != nil {
		return err
	}

	output := c.String("output")
	if output == "" {
		output = filepath.Join(
			filepath.Dir(hf.Path),
			"."+filepath.Base(hf.Path))
	}

	if err := copyFile(hf.Path, output); err != nil {
		return err
	}

	logrus.Infof("backup complete")
	return debugFooter(c)
}
