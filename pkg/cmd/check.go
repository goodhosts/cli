package cmd

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func Check() cli.Command {
	return cli.Command{
		Name:      "check",
		Aliases:   []string{"c"},
		Usage:     "Check if ip or host exists",
		Action:    check,
		ArgsUsage: "[IP|HOST]",
	}
}
func check(c *cli.Context) error {
	if len(c.Args()) < 1 {
		fmt.Println("No input, pass an ip address or hostname to check.")
		return nil
	}

	hostsfile, err := loadHostsfile(c)
	if err != nil {
		return err
	}
	input := c.Args().First()

	if net.ParseIP(input) != nil {
		if hostsfile.HasIp(input) {
			fmt.Printf("%s exists in hosts life\n", input)
			return nil
		}
	}

	if hostsfile.HasHostname(input) {
		fmt.Printf("%s exists in hosts life\n", input)
		return nil
	}

	return cli.NewExitError(fmt.Sprintf("%s does not match anything in the hosts file", input), 1)
}
