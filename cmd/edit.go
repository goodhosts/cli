package cmd

import (
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func Edit() *cli.Command {
	return &cli.Command{
		Name:    "edit",
		Aliases: []string{"e"},
		Usage:   "Open hosts file in an editor, default vim",
		Action:  edit,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "editor",
				Aliases: []string{"e"},
				Usage:   "Which file editor to use, defaults vim",
				Value:   "vim",
			},
		},
	}
}

func edit(c *cli.Context) error {
	hostsfile, err := loadHostsfile(c, false)
	if err != nil {
		return err
	}

	cmd := exec.Command(c.String("editor"), hostsfile.Path)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
