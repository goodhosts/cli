package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/goodhosts/hostsfile"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/urfave/cli/v2"
)

var (
	version string
	App     = &cli.App{
		Name:   "goodhosts",
		Usage:  "manage your hosts file goodly",
		Action: DefaultAction,
		Commands: append(Commands(), &cli.Command{
			Name:    "version",
			Usage:   "",
			Aliases: []string{"v", "ver"},
			Action: func(c *cli.Context) error {
				logrus.Infof(version)
				return nil
			},
		}),
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.SetFormatter(&logrus.TextFormatter{})
			} else {
				// treat logrus like fmt.Print
				logrus.SetFormatter(&easy.Formatter{
					LogFormat: "%msg%",
				})
			}
			if ctx.Bool("quiet") {
				logrus.SetOutput(io.Discard)
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "",
				Usage:   fmt.Sprintf("override the default hosts: %s", hostsfile.HostsFilePath),
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Turn on verbose debug logging",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Turn on off all logging",
			},
		},
	}
)

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

func Version(v, c, d string) {
	version = fmt.Sprintf("goodhosts %s@%s built on %s", v, c, d)
}

func DefaultAction(c *cli.Context) error {
	return list(c)
}

func loadHostsfile(c *cli.Context, readOnly bool) (*hostsfile.Hosts, error) {
	customHostsfile := c.String("file")
	var hf *hostsfile.Hosts
	var err error

	if customHostsfile != "" {
		logrus.Debugf("loading custom hosts file: %s\n", customHostsfile)
		hf, err = hostsfile.NewCustomHosts(customHostsfile)
	} else {
		logrus.Debugf("loading default hosts file: %s\n", hostsfile.HostsFilePath)
		hf, err = hostsfile.NewHosts()
	}

	if err != nil {
		return hf, err
	}

	if !readOnly && !hf.IsWritable() {
		return hf, fmt.Errorf("Hostsfile %s not writable. Try running with elevated privileges.", hf.Path)
	}

	return hf, nil
}

func outputHostsfile(hf *hostsfile.Hosts, all bool) {
	for _, line := range hf.Lines {
		if !all {
			if line.IsComment() || line.Raw == "" {
				continue
			}
		}

		lineOutput := line.Raw
		if line.IsMalformed() {
			logrus.Debugf("malformed line: %s", line.Err)
			if !line.HasComment() {
				lineOutput = fmt.Sprintf("%s #", lineOutput)
			}
			lineOutput = fmt.Sprintf("%s <<< Malformed!", lineOutput)
		}
		logrus.Info(lineOutput + "\n")
	}
}

func debugFooter(c *cli.Context) error {
	if c.Command.Name != "debug" && !c.Bool("debug") {
		return nil
	}

	hf, err := loadHostsfile(c, true)
	if err != nil {
		return err
	}

	logrus.Infof("hosts file path: %s\n", hf.Path)

	var comments, empty, entry, malformed int
	for _, line := range hf.Lines {

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
		{"lines", fmt.Sprintf("%d", len(hf.Lines))},
		{"entries", fmt.Sprintf("%d", entry)},
		{"comments", fmt.Sprintf("%d", comments)},
		{"empty", fmt.Sprintf("%d", empty)},
		{"malformed", fmt.Sprintf("%d", malformed)},
	}

	buf := &bytes.Buffer{}
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Type", "Count"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	logrus.Info(buf)

	return nil
}

func copyFile(src, dst string) (err error) {
	logrus.Debugf("copying file: src %s, dst %s",
		src, dst)

	var fi os.FileInfo
	fi, err = os.Stat(src)
	if err != nil {
		return err
	}

	if !fi.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	var source *os.File
	source, err = os.Open(src)
	if err != nil {
		return err
	}

	defer func() {
		err = source.Close()
	}()

	var destination *os.File
	destination, err = os.Create(dst)
	if err != nil {
		return err
	}

	defer func() {
		err = destination.Close()
	}()

	_, err = io.Copy(destination, source)
	return
}
