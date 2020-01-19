package cmd

import (
	"fmt"
	"os"

	hosts "github.com/luthermonson/goodhosts/pkg/hostsfile"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	return list(c)
}

func Commands() []*cli.Command {
	return []*cli.Command{
		Check(),
		List(),
		Add(),
		Remove(),
		Debug(),
		Backup(),
		Restore(),
	}
}

func DefaultAction(c *cli.Context) error {
	return list(c)
}

func loadHostsfile(c *cli.Context) (hosts.Hosts, error) {
	customHostsfile := c.String("custom")
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

func debugFooter(c *cli.Context) error {
	if c.Command.Name != "debug" && !c.Bool("debug") {
		return nil
	}

	hostsfile, err := loadHostsfile(c)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", hostsfile.Path)

	var comments, empty, entry, malformed int
	for _, line := range hostsfile.Lines {

		if line.IsComment() {
			comments++
		}

		if line.Raw == "" {
			empty++
		}

		if line.IsMalformed() {
			malformed++
		}

		if line.IsValid() {
			entry++
		}
	}

	data := [][]string{
		[]string{"lines", fmt.Sprintf("%d", len(hostsfile.Lines))},
		[]string{"entries", fmt.Sprintf("%d", entry)},
		[]string{"comments", fmt.Sprintf("%d", comments)},
		[]string{"empty", fmt.Sprintf("%d", empty)},
		[]string{"malformed", fmt.Sprintf("%d", malformed)},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Count"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}
