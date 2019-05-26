package main

import (
	"log"
	"os"

	"github.com/luthermonson/goodhosts/pkg/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goodhosts"
	app.Usage = "manage your hosts file goodly"
	app.Action = cmd.Run
	app.Commands = cmd.Commands()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "custom, c",
			Value: "",
			Usage: "override the default hosts file",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
