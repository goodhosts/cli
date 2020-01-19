package hostsfile

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

type Hosts struct {
	Path  string
	Lines []HostsLine
}

// Return a new instance of ``Hosts`` using the default hosts file path.
func NewHosts() (Hosts, error) {
	osHostsFilePath := os.ExpandEnv(filepath.FromSlash(hostsFilePath))

	if env, isset := os.LookupEnv("HOSTS_PATH"); isset && len(env) > 0 {
		osHostsFilePath = os.ExpandEnv(filepath.FromSlash(env))
	}

	return NewCustomHosts(osHostsFilePath)
}

// Return a new instance of ``Hosts`` using a custom hosts file path.
func NewCustomHosts(osHostsFilePath string) (Hosts, error) {
	hosts := Hosts{
		Path: osHostsFilePath,
	}

	if err := hosts.Load(); err != nil {
		return hosts, err
	}

	return hosts, nil
}

// Return ```true``` if hosts file is writable.
func (h *Hosts) IsWritable() bool {
	_, err := os.OpenFile(h.Path, os.O_WRONLY, 0660)
	if err != nil {
		return false
	}

	return true
}

// Load the hosts file into ```l.Lines```.
// ```Load()``` is called by ```NewHosts()``` and ```Hosts.Flush()``` so you
// generally you won't need to call this yourself.
func (h *Hosts) Load() error {
	var lines []HostsLine

	file, err := os.Open(h.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := NewHostsLine(scanner.Text())
		if err != nil {
			return err
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	h.Lines = lines

	return nil
}

// Flush any changes made to hosts file.
func (h Hosts) Flush() error {
	file, err := os.Create(h.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	for _, line := range h.Lines {
		fmt.Fprintf(w, "%s%s", line.Raw, eol)
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return h.Load()
}

// Add an entry to the hosts file.
func (h *Hosts) Add(ip string, hosts ...string) error {
	if net.ParseIP(ip) == nil {
		return errors.New(fmt.Sprintf("%q is an invalid IP address.", ip))
	}

	position := h.getIpPosition(ip)
	if position == -1 {
		endLine := NewHostsLine(buildRawLine(ip, hosts))
		// Ip line is not in file, so we just append our new line.
		h.Lines = append(h.Lines, endLine)
	} else {
		// Otherwise, we replace the line in the correct position
		newHosts := h.Lines[position].Hosts
		for _, addHost := range hosts {
			if itemInSlice(addHost, newHosts) {
				continue
			}

			newHosts = append(newHosts, addHost)
		}
		endLine := NewHostsLine(buildRawLine(ip, newHosts))
		h.Lines[position] = endLine
	}

	return nil
}

// Return a bool if ip/host combo in hosts file.
func (h Hosts) Has(ip string, host string) bool {
	pos := h.getHostPosition(ip, host)

	return pos != -1
}

// Return a bool if hostname in hosts file.
func (h Hosts) HasHostname(host string) bool {
	pos := h.getHostnamePosition(host)

	return pos != -1
}

func (h Hosts) HasIp(ip string) bool {
	pos := h.getIpPosition(ip)

	return pos != -1
}

// Remove an entry from the hosts file.
func (h *Hosts) Remove(ip string, hosts ...string) error {
	var outputLines []HostsLine

	if net.ParseIP(ip) == nil {
		return errors.New(fmt.Sprintf("%q is an invalid IP address.", ip))
	}

	for _, line := range h.Lines {

		// Bad lines or comments just get readded.
		if line.Err != nil || line.IsComment() || line.IP != ip {
			outputLines = append(outputLines, line)
			continue
		}

		var newHosts []string
		for _, checkHost := range line.Hosts {
			if !itemInSlice(checkHost, hosts) {
				newHosts = append(newHosts, checkHost)
			}
		}

		// If hosts is empty, skip the line completely.
		if len(newHosts) > 0 {
			newLineRaw := line.IP

			for _, host := range newHosts {
				newLineRaw = fmt.Sprintf("%s %s", newLineRaw, host)
			}
			newLine := NewHostsLine(newLineRaw)
			outputLines = append(outputLines, newLine)
		}
	}

	h.Lines = outputLines
	return nil
}

// Remove  entries by hostname from the hosts file.
func (h *Hosts) RemoveByHostname(host string) error {
	pos := h.getHostnamePosition(host)
	for pos > -1 {
		if len(h.Lines[pos].Hosts) > 1 {
			h.Lines[pos].Hosts = removeFromSlice(host, h.Lines[pos].Hosts)
			h.Lines[pos].RegenRaw()
		} else {
			h.removeByPosition(pos)
		}
		pos = h.getHostnamePosition(host)
	}

	return nil
}

func (h *Hosts) RemoveByIp(ip string) error {
	pos := h.getIpPosition(ip)
	for pos > -1 {
		h.removeByPosition(pos)
		pos = h.getIpPosition(ip)
	}

	return nil
}

func (h *Hosts) removeByPosition(pos int) {
	h.Lines[pos] = h.Lines[len(h.Lines)-1]
	h.Lines = h.Lines[:len(h.Lines)-1]
}

func (h Hosts) getHostPosition(ip string, host string) int {
	for i := range h.Lines {
		line := h.Lines[i]
		if !line.IsComment() && line.Raw != "" {
			if ip == line.IP && itemInSlice(host, line.Hosts) {
				return i
			}
		}
	}

	return -1
}

func (h Hosts) getHostnamePosition(host string) int {
	for i := range h.Lines {
		line := h.Lines[i]
		if !line.IsComment() && line.Raw != "" {
			if itemInSlice(host, line.Hosts) {
				return i
			}
		}
	}

	return -1
}

func (h Hosts) getIpPosition(ip string) int {
	for i := range h.Lines {
		line := h.Lines[i]
		if !line.IsComment() && line.Raw != "" {
			if line.IP == ip {
				return i
			}
		}
	}

	return -1
}
