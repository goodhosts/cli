package cmd

import (
	"github.com/goodhosts/hostsfile"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Clean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"cl"},
		Usage:   "Clean the hostsfile by doing: remove dupe IPs, for each IPs remove dupe hosts and sort, sort all IPs, split hosts per OS limitations",
		Action:  clean,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Dry run only, will output contents of the cleaned hostsfile without writing the changes",
			},
			&cli.BoolFlag{
				Name:    "all",
				Usage:   "Perform all Cleanup Jobs",
				Aliases: []string{"A"},
			},
			&cli.BoolFlag{
				Name:    "remove-duplicate-ips",
				Usage:   "Remove all duplicate ips",
				Aliases: []string{"rdi"},
			},
			&cli.BoolFlag{
				Name:    "remove-duplicate-hosts",
				Usage:   "Remove all duplicate hosts",
				Aliases: []string{"rdh"},
			},
			&cli.BoolFlag{
				Name:    "sort-hosts",
				Usage:   "Sort each ip's hosts alphabetically",
				Aliases: []string{"sh"},
			},
			&cli.BoolFlag{
				Name:    "sort-ips",
				Usage:   "Sort all ips numerically",
				Aliases: []string{"si"},
			},
			&cli.IntFlag{
				Name:    "hosts-per-line",
				Usage:   "Number of hosts allowed per line",
				Value:   hostsfile.HostsPerLine,
				Aliases: []string{"hpl"},
			},
		},
	}
}

func clean(c *cli.Context) error {
	h, err := loadHostsfile(c, false)
	if err != nil {
		return err
	}

	if c.Int("hosts-per-line") != hostsfile.HostsPerLine {
		hostsfile.HostsPerLine = c.Int("hosts-per-line")
	}

	if c.Bool("all") {
		h.Clean()
	} else {
		if c.Bool("remove-duplicate-ips") {
			h.RemoveDuplicateIps()
		}
		if c.Bool("remove-duplicate-hosts") {
			h.RemoveDuplicateHosts()
		}
		if c.Bool("sort-hosts") {
			h.SortHosts()
		}
		if c.Bool("sort-ips") {
			h.SortByIp()
		}
		// needed for windows for 9/line, -1 default for linux will noop but if passed by cli we will run
		h.HostsPerLine(hostsfile.HostsPerLine)
	}

	if c.Bool("dry-run") {
		logrus.Debugln("performing a dry run, writing output")
		outputHostsfile(h, true)
		return debugFooter(c)
	}

	if err := h.Flush(); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return debugFooter(c)
}
