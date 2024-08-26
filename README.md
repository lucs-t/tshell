# tshell

A simple ssh connection tool for managing multiple servers.

## Installation

```bash
$ git clone https://github.com/lucs-t/tshell.git
$ cd tshell
$ make build
$ ./tshell
```

## Usage

```bash
$ ./tshell -h
Usage: tshell [command]
Description:
  A simple tool to manage SSH connections
Commands:
  cnn     Connect to a SSH connection
  add     Add a new SSH connection
  remove  Remove a SSH connection
  show    Show all SSH connection
  s3      Synchronize ssh connection information to s3

## Add a new SSH connection
$ ./tshell add -h
Usage: tshell add [options]
Description:
  Add a new SSH connection to the list

Example:
  # Add a new SSH connection with password
  tshell add -u user:host:port -p 123456

  # Add a new SSH connection with key
  tshell add -u user:host:port -k /path/to/key

Options:
  -k string
        key path for ssh connection
  -n string
        ssh connection name,if not set,use host as name
  -p string
        password for ssh connection
  -u string
        user:host:port for ssh connection

## Connect to a SSH connection
$ ./tshell cnn {sshName}

## Remove SSH connection
$ ./tshell remove -h
Usage: tshell remove [options]
Description:
  Remove a SSH connection from the list

Example:
  # Remove a SSH connection
  tshell remove -n sshName

  # Remove all SSH connection
  tshell remove --all

Options:
  -all
        remove all ssh connection
  -n string
        ssh connection name

## Show all SSH connection
$ ./tshell show

## Synchronize ssh connection information to s3
$ ./tshell s3 -h
Usage: tshell config [options]
Description:
  Configure the tshell
Options:
  add    Add ssh configuration
  remove Remove ssh configuration
  info   Show ssh configuration info

## Add ssh configuration
$ ./tshell s3 add -h
Usage: tshell config add [options]
Description:
  Add a new configuration to the list

Example:
  # Add a new configuration
  tshell config add -ak accessKey -sk accessSecret -region region -path bucket:updatePath

Options:
  -e string
        endpoint
  -k string
        access key
  -p string
        format: bucket:updatePath
  -r string
        region
  -s string
        access secret

## Remove ssh configuration
$ ./tshell s3 remove

## Show ssh configuration info
$ ./tshell s3 info
```

### hope you enjoy it !

### welcome to contribute !!!