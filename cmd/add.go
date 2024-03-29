package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Aliases:   []string{"a"},
		Usage:     "Add an entry to the hostsfile",
		Action:    add,
		ArgsUsage: "[IP] [HOST] ([HOST]...)",
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
func add(c *cli.Context) error {
	args := c.Args()

	if args.Len() < 2 {
		logrus.Infof("adding a hostsfile entry requires an ip and a hostname.")
		return nil
	}

	hf, err := loadHostsfile(c, false)
	if err != nil {
		return err
	}

	ip := args.Slice()[0]
	uniqueHosts := map[string]bool{}
	var hostEntries []string

	for i := 1; i < args.Len(); i++ {
		uniqueHosts[args.Slice()[i]] = true
	}

	for key := range uniqueHosts {
		hostEntries = append(hostEntries, key)
	}

	err = hf.Add(ip, hostEntries...)
	if err != nil {
		return cli.Exit(err.Error(), 2)
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

	logrus.Infof("hosts entry added: %s %s\n", ip, strings.Join(hostEntries, " "))
	return debugFooter(c)
}
