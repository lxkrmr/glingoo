package main

import (
	"fmt"
	"os"

	"github.com/lxkrmr/glingoo/internal/cmd"
)

const help = `glingoo - Odoo translation CLI

Usage:
  glingoo <command> [args]

Commands:
  context   Manage connection contexts
  export    Export a PO translation file from Odoo
  install   Load language terms into Odoo

Examples:
  glingoo context create mydev
  glingoo context list
  glingoo context use mydev
  glingoo export my_addon de_DE /path/to/my_addon/i18n
  glingoo install de_DE

Run 'glingoo <command> --help' for command-specific usage.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(0)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "context":
		cmd.RunContext(args)
	case "export":
		cmd.RunExport(args)
	case "install":
		cmd.RunInstall(args)
	case "help":
		fmt.Println(help)
	default:
		cmd.WriteError("", fmt.Errorf("unknown command %q - run glingoo --help", command))
		os.Exit(1)
	}
}
