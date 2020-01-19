package cmd

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Aliases:   []string{"a"},
		Usage:     "Add an entry to the hostsfile",
		Action:    add,
		ArgsUsage: "[IP] [HOST] ([HOST]...)",
	}
}
func add(c *cli.Context) error {

	args := c.Args()

	if args.Len() < 2 {
		fmt.Println("Adding an entry requires an ip and a hostname.")
		return nil
	}

	hostsfile, err := loadHostsfile(c)
	if err != nil {
		return err
	}

	ip := args.Slice()[0]
	uniqueHosts := map[string]bool{}
	var hostEntries []string

	for i := 1; i < args.Len(); i++ {
		uniqueHosts[args.Slice()[i]] = true
	}

	for key, _ := range uniqueHosts {
		hostEntries = append(hostEntries, key)
	}

	err = hostsfile.Add(ip, hostEntries...)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	err = hostsfile.Flush()
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	fmt.Printf("Added: %s %s\n", ip, strings.Join(hostEntries, " "))
	return nil
}
