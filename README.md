# goodhosts cli, fork of Lex Toumbourou's project [goodhosts](https://github.com/lextoumbourou/goodhosts)

[![Go Reference](https://pkg.go.dev/badge/github.com/goodhosts/cli.svg)](https://pkg.go.dev/github.com/goodhosts/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/goodhosts/cli)](https://goreportcard.com/report/github.com/goodhosts/cli)

Simple [hosts file](http://en.wikipedia.org/wiki/Hosts_%28file%29) (```/etc/hosts```) management in a Go cli. One simple
interface for any OS or architecture, script and automate hosts file updates using one simple tool.

## Features

- List, add, remove and check hosts file entries from code or the command-line
  - Remove by IP, Host, or IP/Host
  - `check` returns proper exit codes for scripting e.g. `goodhosts check 10.0.5.12 || echo "Missing hosts entry for 10.0.5.12"`
- Clean hostsfile command will
 - Consolidate duplicate IPs
 - Remove duplicate hosts
 - Alpha sort Hosts
 - Sort IPs
 - Help with OS limitations, e.g. 9 hosts per IP line in windows
- linux/darwin/windows support
- Custom hosts file support
- Backup/Restore
- Quick inline editor (vim/nano)

## Installation

Visit the [releases](https://github.com/goodhosts/cli/releases/) page and download the proper binary for your
architecture. Unzip and run in place, put in your system path (linux: `/usr/local/bin` win: `~/bin`) for easier access.

## Usage

For full usage directions simply call `goodhosts -h`

```shell
$ goodhosts -h
  NAME:
    goodhosts - manage your hosts file goodly

  USAGE:
     goodosts [global options] command [command options] [arguments...]
  
  COMMANDS:
     add, a         Add an entry to the hostsfile
     backup         Backup hosts file
     check, c       Check if ip or host exists
     clean, cl      Clean the hostsfile by doing: remove dupe IPs, for each IPs remove dupe hosts and sort, sort all IPs, split hosts per OS limitations
     debug, d       Show debug table for hosts file
     edit, e        Open hosts file in an editor, default vim
     list, ls       List all entries in the hostsfile
     remove, rm, r  Remove ip or host(s) if exists
     restore        Restore hosts file from backup
     version
     help, h        Shows a list of commands or help for one command
  
  GLOBAL OPTIONS:
     --file value, -f value  override the default hosts: ${SystemRoot}/System32/drivers/etc/hosts
     --debug, -d             Turn on verbose debug logging (default: false)
     --quiet, -q             Turn on off all logging (default: false)
     --help, -h              show help (default: false)
```

Each sub-command can be called with a `-h` option to see detailed help information.
```shell

 $ ./goodhosts list -h
 NAME:
    goodhosts list - List all entries in the hostsfile

 USAGE:
    goodhosts list [command options] [arguments...]

 OPTIONS:
    --all       Show all entries in the hosts file including commented lines. (default: false)
    --help, -h  show help (default: false)
```

## License

[MIT](LICENSE)

