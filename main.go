package main

import (
	"io/ioutil"
	"os"

	"github.com/luthermonson/goodhosts/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "goodhosts",
		Usage:    "manage your hosts file goodly",
		Action:   cmd.DefaultAction,
		Commands: cmd.Commands(),
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
			}
			if ctx.Bool("quiet") {
				logrus.SetOutput(ioutil.Discard)
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "custom, c",
				Value: "",
				Usage: "override the default hosts file",
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

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
