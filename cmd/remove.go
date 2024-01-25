package cmd

import (
	"errors"
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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "clean",
				Aliases: []string{"c"},
				Usage:   "Clean the hostsfile after adding an entry. See clean command for more details",
			},
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Dry run only, will output contents of the new hostsfile without writing the changes.",
			},
		},
	}
}

func remove(c *cli.Context) error {
	args := c.Args()

	if args.Len() == 0 {
		return errors.New("no input")
	}

	hf, err := loadHostsfile(c, false)
	if err != nil {
		return err
	}

	if args.Len() == 1 { //could be ip or hostname
		return processSingleArg(hf, args.Slice()[0])
	}

	uniqueHosts := map[string]bool{}
	var hostEntries []string

	for i := 1; i < args.Len(); i++ {
		uniqueHosts[args.Slice()[i]] = true
	}

	for key := range uniqueHosts {
		hostEntries = append(hostEntries, key)
	}

	if net.ParseIP(args.Slice()[0]) != nil {
		if hf.HasIP(args.Slice()[0]) {
			err = hf.Remove(args.Slice()[0], hostEntries...)
			if err != nil {
				return err
			}
		}
	} else {
		hostEntries = append([]string{args.Slice()[0]}, hostEntries...)
		for _, value := range hostEntries {
			if err := hf.RemoveByHostname(value); err != nil {
				return err
			}
		}
	}

	if c.Bool("clean") {
		hf.Clean()
	}

	if c.Bool("dry-run") {
		logrus.Debugln("performing a dry run, writing output")
		outputHostsfile(hf, true)
		return debugFooter(c)
	}

	logrus.Debugln("flushing hosts file to disk")
	if err := hf.Flush(); err != nil {
		return cli.Exit(err.Error(), 2)
	}

	logrus.Infof("entry removed: %s\n", strings.Join(hostEntries, " "))
	return debugFooter(c)
}

func processSingleArg(hf *hostsfile.Hosts, arg string) error {
	if net.ParseIP(arg) != nil {
		logrus.Infof("removing ip %s\n", arg)
		hf.RemoveByIP(arg)
		return hf.Flush()
	}

	logrus.Infof("removing hostname %s\n", arg)
	if err := hf.RemoveByHostname(arg); err != nil {
		return err
	}
	return hf.Flush()
}
