package cmd

import (
	"github.com/luthermonson/goodhosts/pkg/hosts"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
	return list(c)
}

func Commands() []cli.Command {
	return []cli.Command{
		Check(),
		List(),
		Add(),
		Remove(),
	}
}

func loadHostsfile(c *cli.Context) (hosts.Hosts, error) {
	customHostsfile := c.GlobalString("custom")
	var hostsfile hosts.Hosts
	var err error

	if customHostsfile != "" {
		hostsfile, err = hosts.NewCustomHosts(customHostsfile)
	} else {
		hostsfile, err = hosts.NewHosts()
	}

	if err != nil {
		return hostsfile, cli.NewExitError(err, 1)
	}

	if !hostsfile.IsWritable() {
		return hostsfile, cli.NewExitError("Host file not writable. Try running with elevated privileges.", 1)
	}

	return hostsfile, nil
}
