package cmd

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Check() *cli.Command {
	return &cli.Command{
		Name:      "check",
		Aliases:   []string{"c"},
		Usage:     "Check if ip or host exists. With IP and hosts: check if all hosts are mapped to IP",
		Action:    check,
		ArgsUsage: "[IP|HOST] or [IP] [HOST] ([HOST]...)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Suppress output, only return exit code",
			},
			&cli.BoolFlag{
				Name:  "any",
				Usage: "With multiple hosts: succeed if ANY host matches (default: ALL must match)",
			},
		},
	}
}

func check(c *cli.Context) error {
	if c.Args().Len() < 1 {
		logrus.Infof("No input, pass an ip address or hostname to check.")
		return nil
	}

	hf, err := loadHostsfile(c, true) // readOnly=true, no sudo needed
	if err != nil {
		return err
	}

	args := c.Args().Slice()
	firstArg := args[0]
	quiet := c.Bool("quiet")

	// Single argument: check if IP or hostname exists (backward compatible)
	if c.Args().Len() == 1 {
		if net.ParseIP(firstArg) != nil {
			if hf.HasIP(firstArg) {
				if !quiet {
					logrus.Infof("%s exists in hosts file\n", firstArg)
				}
				return nil
			}
		}

		if hf.HasHostname(firstArg) {
			if !quiet {
				logrus.Infof("%s exists in hosts file\n", firstArg)
			}
			return nil
		}

		return cli.Exit(fmt.Sprintf("%s does not match anything in the hosts file", firstArg), 1)
	}

	// Multiple arguments: first must be IP, rest are hostnames
	if net.ParseIP(firstArg) == nil {
		return cli.Exit(fmt.Sprintf("%s is not a valid IP address. When checking multiple hosts, first argument must be an IP.", firstArg), 2)
	}

	ip := firstArg
	hosts := args[1:]

	if c.Bool("any") {
		// Check if ANY host is mapped to IP
		if hf.HasAny(ip, hosts...) {
			if !quiet {
				logrus.Infof("At least one host is mapped to %s\n", ip)
			}
			return nil
		}
		return cli.Exit(fmt.Sprintf("None of the specified hosts are mapped to %s", ip), 1)
	}

	// Default: Check if ALL hosts are mapped to IP
	if hf.HasAll(ip, hosts...) {
		if !quiet {
			logrus.Infof("All specified hosts are mapped to %s\n", ip)
		}
		return nil
	}

	// Report which hosts are missing
	if !quiet {
		missing := []string{}
		for _, host := range hosts {
			if !hf.Has(ip, host) {
				missing = append(missing, host)
			}
		}
		return cli.Exit(fmt.Sprintf("Missing hosts for %s: %v", ip, missing), 1)
	}

	return cli.Exit("", 1)
}
