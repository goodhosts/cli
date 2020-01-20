package cmd

import (
	"net"
	"strings"

	"github.com/goodhosts/hostsfile"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Remove() *cli.Command {
	return &cli.Command{
		Name:      "remove",
		Aliases:   []string{"rm", "r"},
		Usage:     "Remove ip or host(s) if exists",
		Action:    remove,
		ArgsUsage: "[IP|HOST] or [IP] [HOST] ([HOST]...)",
	}
}
func remove(c *cli.Context) error {
	args := c.Args()
	hostsfile, err := loadHostsfile(c)
	if err != nil {
		return err
	}

	if args.Len() == 0 {
		return cli.NewExitError("no input", 1)
	}

	if args.Len() == 1 { //could be ip or hostname
		return processSingleArg(hostsfile, args.Slice()[0])
	}

	uniqueHosts := map[string]bool{}
	var hostEntries []string

	for i := 1; i < args.Len(); i++ {
		uniqueHosts[args.Slice()[i]] = true
	}

	for key, _ := range uniqueHosts {
		hostEntries = append(hostEntries, key)
	}

	if net.ParseIP(args.Slice()[0]) != nil {
		if hostsfile.HasIp(args.Slice()[0]) {
			err = hostsfile.Remove(args.Slice()[0], hostEntries...)
			if err != nil {
				return cli.NewExitError(err.Error(), 2)
			}
		}
	} else {
		hostEntries = append(hostEntries, args.Slice()[0])
		for _, value := range hostEntries {
			hostsfile.RemoveByHostname(value)
		}
	}

	err = hostsfile.Flush()
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	logrus.Infof("entry removed: %s\n", strings.Join(hostEntries, " "))
	return debugFooter(c)
}

func processSingleArg(hostsfile hostsfile.Hosts, arg string) error {
	if net.ParseIP(arg) != nil {
		logrus.Infof("removing ip %s\n", arg)
		if err := hostsfile.RemoveByIp(arg); err != nil {
			return err
		}
		if err := hostsfile.Flush(); err != nil {
			return err
		}

		return nil
	}

	logrus.Infof("removing hostname %s\n", arg)

	if err := hostsfile.RemoveByHostname(arg); err != nil {
		return err
	}
	if err := hostsfile.Flush(); err != nil {
		return err
	}

	return nil
}
