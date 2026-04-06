package cmd

import (
	"flag"
	"fmt"
	"os"
)

const contextHelp = `Manage connection contexts for Odoo instances.

Usage:
  glingoo context <subcommand> [args]

Subcommands:
  create <name>  Create a new context (interactive wizard)
  list           List all contexts (current marked with *)
  use <name>     Set the context to use by default
  remove <name>  Delete a context

Examples:
  glingoo context create mydev
  glingoo context list
  glingoo context use mydev
  glingoo context remove mydev`

// RunContext handles the context subcommand.
func RunContext(args []string) {
	if len(args) == 0 {
		fmt.Println(contextHelp)
		os.Exit(0)
	}

	subcommand := args[0]
	subargs := args[1:]

	switch subcommand {
	case "create":
		runContextCreate(subargs)
	case "list":
		runContextList()
	case "use":
		runContextUse(subargs)
	case "remove":
		runContextRemove(subargs)
	case "help":
		fmt.Println(contextHelp)
	default:
		WriteError("context", fmt.Errorf("unknown subcommand %q - run 'glingoo context --help'", subcommand))
		os.Exit(1)
	}
}

func runContextCreate(args []string) {
	fs := flag.NewFlagSet("context create", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(contextHelp) }

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		WriteError("context create", err)
		os.Exit(1)
	}

	positional := fs.Args()
	if len(positional) == 0 {
		WriteError("context create", fmt.Errorf("context name is required"))
		os.Exit(1)
	}
	if len(positional) > 1 {
		WriteError("context create", fmt.Errorf("unexpected argument %q", positional[1]))
		os.Exit(1)
	}

	name := positional[0]

	if err := CreateContextInteractive(name); err != nil {
		WriteError("context create", err)
		os.Exit(1)
	}

	write(successPayload("context create", map[string]any{
		"name":    name,
		"message": fmt.Sprintf("Context %q created and set as current", name),
	}))
}

func runContextList() {
	names, current, err := ListContexts()
	if err != nil {
		WriteError("context list", err)
		os.Exit(1)
	}

	if len(names) == 0 {
		write(successPayload("context list", map[string]any{
			"contexts": []string{},
			"message":  "No contexts yet - run 'glingoo context create <name>' to add one",
		}))
		return
	}

	// Sort for consistent output
	// (in practice Go maps are unordered, but we'll leave it as is for now)
	contexts := make([]map[string]any, len(names))
	for i, name := range names {
		marker := " "
		if name == current {
			marker = "*"
		}
		contexts[i] = map[string]any{
			"name":    name,
			"current": name == current,
			"marker":  marker,
		}
	}

	write(successPayload("context list", map[string]any{
		"contexts": contexts,
		"current":  current,
	}))
}

func runContextUse(args []string) {
	fs := flag.NewFlagSet("context use", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(contextHelp) }

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		WriteError("context use", err)
		os.Exit(1)
	}

	positional := fs.Args()
	if len(positional) == 0 {
		WriteError("context use", fmt.Errorf("context name is required"))
		os.Exit(1)
	}
	if len(positional) > 1 {
		WriteError("context use", fmt.Errorf("unexpected argument %q", positional[1]))
		os.Exit(1)
	}

	name := positional[0]

	if err := SetCurrentContext(name); err != nil {
		WriteError("context use", err)
		os.Exit(1)
	}

	write(successPayload("context use", map[string]any{
		"name":    name,
		"message": fmt.Sprintf("Context switched to %q", name),
	}))
}

func runContextRemove(args []string) {
	fs := flag.NewFlagSet("context remove", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(contextHelp) }

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		WriteError("context remove", err)
		os.Exit(1)
	}

	positional := fs.Args()
	if len(positional) == 0 {
		WriteError("context remove", fmt.Errorf("context name is required"))
		os.Exit(1)
	}
	if len(positional) > 1 {
		WriteError("context remove", fmt.Errorf("unexpected argument %q", positional[1]))
		os.Exit(1)
	}

	name := positional[0]

	if err := RemoveContext(name); err != nil {
		WriteError("context remove", err)
		os.Exit(1)
	}

	write(successPayload("context remove", map[string]any{
		"name":    name,
		"message": fmt.Sprintf("Context %q removed", name),
	}))
}
