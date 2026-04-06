package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/lxkrmr/godoorpc"
)

const installHelp = `Load language terms into Odoo.

Usage:
  glingoo install <lang>

Arguments:
  lang    Odoo language code (e.g. de_DE)

Examples:
  glingoo install de_DE
  glingoo install it_IT

Uses the current context. Set it with: glingoo context use <name>`

// installInput holds the parsed data for an install command.
type installInput struct {
	lang string
}

// parseInstallArgs parses flags and positional args — calculation.
func parseInstallArgs(args []string) (installInput, error) {
	fs := flag.NewFlagSet("install", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(installHelp) }

	if err := fs.Parse(args); err != nil {
		return installInput{}, err
	}

	positional := fs.Args()
	if len(positional) == 0 {
		return installInput{}, fmt.Errorf("lang is required - run 'glingoo install --help'")
	}
	if len(positional) > 1 {
		return installInput{}, fmt.Errorf(
			"unexpected argument %q - install takes exactly one lang code",
			positional[1],
		)
	}

	return installInput{lang: positional[0]}, nil
}

// buildInstallResult shapes the data for the JSON response — pure calculation.
func buildInstallResult(lang string) map[string]any {
	return map[string]any{
		"lang":   lang,
		"result": "installed",
	}
}

// findLangID resolves a language code to its res.lang record ID — side effect.
func findLangID(client *godoorpc.Client, lang string) (int, error) {
	result, err := client.ExecuteKW(
		"res.lang", "search",
		godoorpc.Args{godoorpc.Domain{
			godoorpc.Condition{Field: "code", Op: "=", Value: lang},
		}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		return 0, fmt.Errorf("could not search for language %q: %w", lang, err)
	}

	ids, ok := result.([]any)
	if !ok || len(ids) == 0 {
		return 0, fmt.Errorf(
			"language %q not found in Odoo - is it installed? check Settings > Translations",
			lang,
		)
	}

	id, ok := ids[0].(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected id type for language %q", lang)
	}

	return int(id), nil
}

// loadLanguageTerms loads language terms into Odoo via the language install wizard — side effect.
func loadLanguageTerms(client *godoorpc.Client, langID int, lang string) error {
	wizardID, err := client.ExecuteKW(
		"base.language.install", "create",
		godoorpc.Args{map[string]any{
			"lang_ids":  []any{[]any{6, 0, []int{langID}}},
			"overwrite": true,
		}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		return fmt.Errorf("could not create language install wizard: %w", err)
	}

	wid, ok := wizardID.(float64)
	if !ok {
		return fmt.Errorf("unexpected wizard id type")
	}

	_, err = client.ExecuteKW(
		"base.language.install", "lang_install",
		godoorpc.Args{[]int{int(wid)}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		return fmt.Errorf("could not load language terms for %q: %w", lang, err)
	}

	return nil
}

// RunInstall executes the install command: loads language terms into Odoo.
func RunInstall(args []string) {
	input, err := parseInstallArgs(args)
	if err == flag.ErrHelp {
		os.Exit(0)
	}
	if err != nil {
		write(errorPayload("install", err))
		os.Exit(1)
	}

	_, ctx, err := GetCurrentContext()
	if err != nil {
		write(errorPayload("install", err))
		os.Exit(1)
	}

	conn := ConvertContextToConnFlags(ctx)
	client, err := conn.Connect()
	if err != nil {
		write(errorPayload("install", fmt.Errorf("cannot connect to Odoo at %s - is Odoo running?", conn.URL)))
		os.Exit(1)
	}

	langID, err := findLangID(client, input.lang)
	if err != nil {
		write(errorPayload("install", err))
		os.Exit(1)
	}

	if err := loadLanguageTerms(client, langID, input.lang); err != nil {
		write(errorPayload("install", err))
		os.Exit(1)
	}

	write(successPayload("install", buildInstallResult(input.lang)))
}
