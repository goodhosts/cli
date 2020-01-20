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
	hostsfile, err := loadHostsfile(c)
	if err != nil {
		return err
	}

	output := c.String("output")
	if output == "" {
		output = filepath.Join(
			filepath.Dir(hostsfile.Path),
			"."+filepath.Base(hostsfile.Path))
	}

	_, err = copyFile(hostsfile.Path, output)
	if err != nil {
		return err
	}
	logrus.Infof("backup complete")
	return debugFooter(c)
}
