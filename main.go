package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lxkrmr/glingoo/internal/cmd"
)

const help = `glingoo - Odoo translation CLI

Usage:
  glingoo --url <url> --db <db> --user <user> --password <password> <command> [args]

Commands:
  export    Export a PO translation file from Odoo
  install   Load language terms into Odoo

Connection flags (required, must come before the command):
  --url       Odoo base URL (e.g. http://localhost:8069)
  --db        Database name
  --user      Login user
  --password  Login password

Examples:
  glingoo --url http://localhost:8069 --db mydb --user admin --password secret export my_addon de_DE /path/to/my_addon/i18n
  glingoo --url http://localhost:8069 --db mydb --user admin --password secret install de_DE

Run 'glingoo <command> --help' for command-specific usage.`

func main() {
	fs := flag.NewFlagSet("glingoo", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(help) }

	var conn cmd.ConnFlags
	cmd.RegisterConnFlags(fs, &conn)

	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	remaining := fs.Args()
	if len(remaining) == 0 {
		fmt.Println(help)
		os.Exit(0)
	}

	switch remaining[0] {
	case "export":
		cmd.RunExport(remaining[1:], conn)
	case "install":
		cmd.RunInstall(remaining[1:], conn)
	case "help":
		fmt.Println(help)
	default:
		cmd.WriteError("", fmt.Errorf("unknown command %q - run glingoo --help", remaining[0]))
		os.Exit(1)
	}
}
