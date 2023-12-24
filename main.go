package main

import (
	"fmt"
	"io"
	"os"

	"github.com/goodhosts/cli/cmd"
	"github.com/goodhosts/hostsfile"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var app = &cli.App{
	Name:   "goodhosts",
	Usage:  "manage your hosts file goodly",
	Action: cmd.DefaultAction,
	Commands: append(cmd.Commands(), &cli.Command{
		Name:    "version",
		Usage:   "",
		Aliases: []string{"v", "ver"},
		Action: func(c *cli.Context) error {
			logrus.Infof("goodhosts %s@%s built on %s", version, commit, date)
			return nil
		},
	}),
	Before: func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.SetFormatter(&logrus.TextFormatter{})
		} else {
			// treat logrus like fmt.Print
			logrus.SetFormatter(&easy.Formatter{
				LogFormat: "%msg%",
			})
		}
		if ctx.Bool("quiet") {
			logrus.SetOutput(io.Discard)
		}
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Value:   "",
			Usage:   fmt.Sprintf("override the default hosts: %s", hostsfile.HostsFilePath),
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Turn on verbose debug logging",
		},
		&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "Turn on off all logging",
		},
	},
}

func main() {
	if err := run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(args []string) error {
	return app.Run(args)
}
