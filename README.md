# glingoo

A CLI tool for managing Odoo translations.

## What it is

`glingoo` is a dev tool built for local Odoo development instances.
Do not use production credentials with glingoo.

```sh
# Create a connection context (one time)
glingoo context create mydev

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

## Setup

Before using `export` or `install`, create a connection context:

```sh
glingoo context create mydev
```

This will prompt for:
- URL (e.g. http://localhost:8069)
- Database name
- Login user
- Password

The context is saved to `~/.config/glingoo/contexts.json` and can be
reused. If you have multiple Odoo instances:

```sh
glingoo context create mydev
glingoo context create staging
glingoo context list
glingoo context use staging   # switch between contexts
```

## Usage

```sh
# Manage contexts
glingoo context create <name>   # Create a new connection context
glingoo context list            # Show all contexts (current marked with *)
glingoo context use <name>      # Set as current context
glingoo context remove <name>   # Delete a context

# Work with translations
glingoo export <addon> <lang> <output-dir>
glingoo install <lang>
```

All output is JSON.

Run `glingoo <command> --help` for command-specific usage.

## License

MIT
