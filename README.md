# glingoo

A CLI tool for managing Odoo translations.

## What it is

`glingoo` is a dev tool built for local Odoo development instances.
Do not use production credentials with glingoo.

```sh
# export a PO translation file from Odoo
glingoo export my_addon de_DE /path/to/my_addon/i18n

# load language terms into Odoo
glingoo install de_DE
```

## Install

```sh
go install github.com/lxkrmr/glingoo@latest
```

Requires Go. The binary lands in `~/go/bin/glingoo`.

If `@latest` resolves to an older version after a new release:

```sh
GOPROXY=direct go install github.com/lxkrmr/glingoo@latest
```

## Usage

Connection flags are required for every command and must come before
the command name:

```sh
glingoo --url <url> --db <db> --user <user> --password <password> <command> [args]
```

If you are a human typing commands, set up a shell alias:

```sh
alias glingoo='glingoo --url http://localhost:8069 --db mydb --user admin --password secret'
glingoo export my_addon de_DE /path/to/my_addon/i18n
glingoo install de_DE
```

If you are a coding assistant, construct the full command with flags
directly - no alias needed.

All output is JSON.

Run `glingoo <command> --help` for command-specific usage.

## License

MIT
