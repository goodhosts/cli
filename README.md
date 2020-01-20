# goodhosts cli, fork of Lex Toumbourou's project [goodhosts](https://github.com/lextoumbourou/goodhosts)

Simple [hosts file](http://en.wikipedia.org/wiki/Hosts_%28file%29) (```/etc/hosts```) management in a Go cli. One simple
interface for any OS or architecture, script and automate hosts file updates using one simple tool.

## Features

* List, add, remove and check hosts file entries from code or the command-line
* Windows support
* Custom hosts file support

## Installation

Visit the [releases](https://github.com/goodhosts/cli/releases/) page and download the proper binary for your 
architecture. Unzip and run in place, put in your system path (linux: `/usr/local/bin` win: `~/bin`) for easier access.

## Usage

For full usage directions simply call `goodhosts -h`

```shell
$ ./goodhosts -h
   NAME:
      goodhosts - manage your hosts file goodly
   
   USAGE:
      goodhosts [global options] command [command options] [arguments...]
   
   COMMANDS:
      check, c       Check if ip or host exists
      list, ls       List all entries in the hostsfile
      add, a         Add an entry to the hostsfile
      remove, rm, r  Remove ip or host(s) if exists
      debug, d       Show debug table for hosts file
      backup         Backup hosts file
      restore        Restore hosts file from backup
      help, h        Shows a list of commands or help for one command
   
   GLOBAL OPTIONS:
      --custom value  override the default hosts file
      --debug, -d     Turn on verbose debug logging (default: false)
      --quiet, -q     Turn on off all logging (default: false)
      --help, -h      show help (default: false)
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

