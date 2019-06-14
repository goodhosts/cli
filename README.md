# Goodhosts

Simple [hosts file](http://en.wikipedia.org/wiki/Hosts_%28file%29) (```/etc/hosts```) management in Go (golang).

## Features

* List, add, remove and check hosts file entries from code or the command-line
* Windows support
* Custom hosts file support

## Command-Line Usage

### Installation

#### Linux

Download the [binary](https://github.com/lextoumbourou/goodhosts/releases/download/v2.1.0/goodhosts-linux) and put it in your path.

```bash
$ wget -O goodhosts https://github.com/lextoumbourou/goodhosts/releases/download/v2.1.0/goodhosts-linux
$ chmod +x goodhosts
$ export PATH=$(pwd):$PATH
$ goodhosts --help
```

#### Windows

Download the [binary](https://github.com/lextoumbourou/goodhosts/releases/download/v2.1.0/goodhosts-windows) and do Windowsy stuff with it (doc PR welcome).


### List entries

```bash
$ goodhosts list
127.0.0.1 localhost
10.0.0.5 my-home-server xbmc-server
10.0.0.6 my-desktop
```

Add ```--all``` flag to include comments.

### Check for an entry

```bash
$ goodhosts check 127.0.0.1 facebook.com
```

### Add an entry

```bash
$ goodhosts add 127.0.0.1 facebook.com
```

Or *entries*.

```bash
$ goodhosts add 127.0.0.1 facebook.com twitter.com gmail.com
```

### Remove an entry

```bash
$ goodhosts rm 127.0.0.1 facebook.com
```

Or *entries*.

```bash
$ goodhosts rm 127.0.0.1 facebook.com twitter.com gmail.com
```

### More

```bash
$ goodhosts --help
```

## License

[MIT](LICENSE)

