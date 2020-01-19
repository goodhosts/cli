package hostsfile

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type HostsLine struct {
	IP    string
	Hosts []string
	Raw   string
	Err   error
}

const commentChar string = "#"

// Return a new instance of ```HostsLine```.
func NewHostsLine(raw string) HostsLine {
	fields := strings.Fields(raw)
	if len(fields) == 0 {
		return HostsLine{Raw: raw}
	}

	output := HostsLine{Raw: raw}
	if !output.IsComment() {
		rawIP := fields[0]
		if net.ParseIP(rawIP) == nil {
			output.Err = errors.New(fmt.Sprintf("Bad hosts line: %q", raw))
		}

		output.IP = rawIP
		output.Hosts = fields[1:]
	}

	return output
}

// Return ```true``` if the line is a comment.
func (l *HostsLine) IsComment() bool {
	trimLine := strings.TrimSpace(l.Raw)
	isComment := strings.HasPrefix(trimLine, commentChar)
	return isComment
}

func (l *HostsLine) IsValid() bool {
	if l.IP != "" {
		return true
	}
	return false
}

func (l *HostsLine) IsMalformed() bool {
	if l.Err != nil {
		return true
	}
	return false
}

func (l *HostsLine) RegenRaw() {
	l.Raw = fmt.Sprintf("%s %s", l.IP, strings.Join(l.Hosts, " "))
}
