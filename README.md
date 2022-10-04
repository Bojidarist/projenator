# Projenator

A simple CLI tool used to generate stuff

# Table of contents

- [Projenator](#projenator)
- [Table of contents](#table-of-contents)
- [Install](#install)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [Build instructions](#build-instructions)
  - [Windows](#windows)
  - [Linux](#linux)
  - [Mac](#mac)


# Install

You can get the latest version from the [releases page](https://github.com/Bojidarist/projenator/releases).

# Usage

```
projenator --help

projenator - A simple CLI tool used to generate stuff

Usage:
  projenator [command]

Available Commands:
  completion       Generate the autocompletion script for the specified shell
  electron-web-app Generate electron web app from a url
  help             Help about any command
  version          Display the current version

Flags:
  -h, --help   help for projenator

Use "projenator [command] --help" for more information about a command.
```

# Dependencies

- go 1.18

# Build instructions

## Windows

```bash
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o builds/projenator-amd64-windows.exe
```

## Linux

```bash
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o builds/projenator-amd64-linux
```

## Mac

```bash
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o builds/projenator-amd64-darwin
```

