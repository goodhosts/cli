package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/goodhosts/hostsfile"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	return list(c)
}

func Commands() []*cli.Command {
	return []*cli.Command{
		Add(),
		Backup(),
		Check(),
		Clean(),
		Debug(),
		Edit(),
		List(),
		Remove(),
		Restore(),
	}
}

func DefaultAction(c *cli.Context) error {
	return list(c)
}

func loadHostsfile(c *cli.Context, readOnly bool) (*hostsfile.Hosts, error) {
	customHostsfile := c.String("file")
	var hfile *hostsfile.Hosts
	var err error

	if customHostsfile != "" {
		logrus.Debugf("loading custom hosts file: %s\n", customHostsfile)
		hfile, err = hostsfile.NewCustomHosts(customHostsfile)
	} else {
		logrus.Debugf("loading default hosts file: %s\n", hostsfile.HostsFilePath)
		hfile, err = hostsfile.NewHosts()
	}

	if err != nil {
		return hfile, cli.Exit(err, 1)
	}

	if !readOnly && !hfile.IsWritable() {
		return hfile, cli.Exit("Host file not writable. Try running with elevated privileges.", 1)
	}

	return hfile, nil
}

func outputHostsfile(hf *hostsfile.Hosts, all bool) {
	for _, line := range hf.Lines {
		if !all {
			if line.IsComment() || line.Raw == "" {
				continue
			}
		}

		lineOutput := fmt.Sprintf("%s\n", line.Raw)
		if line.IsMalformed() {
			lineOutput = fmt.Sprintf("%s # <<< Malformed!\n", lineOutput)
		}

		logrus.Infof(lineOutput)
	}
}

func debugFooter(c *cli.Context) error {
	if c.Command.Name != "debug" && !c.Bool("debug") {
		return nil
	}

	hostsfile, err := loadHostsfile(c, true)
	if err != nil {
		return err
	}

	logrus.Infof("hosts file path: %s\n", hostsfile.Path)

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

func copyFile(src, dst string) (int64, error) {
	logrus.Debugf("copying file: src %s, dst %s",
		src, dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
