package cmd

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lxkrmr/godoorpc"
)

const exportHelp = `Export a PO translation file from Odoo for an addon.

Usage:
  glingoo [connection flags] export <addon> <lang> <output-dir>

Arguments:
  addon       Technical addon name (e.g. my_addon)
  lang        Odoo language code (e.g. de_DE)
  output-dir  Directory to save the PO file (e.g. /path/to/my_addon/i18n)
              The file is named after the language: de_DE -> de.po

Examples:
  glingoo export my_addon de_DE /path/to/my_addon/i18n
  glingoo export my_addon it_IT /path/to/my_addon/i18n`

// exportInput holds the parsed data for an export command.
type exportInput struct {
	addon     string
	lang      string
	outputDir string
}

// poFilename derives the PO filename from an Odoo language code.
// de_DE -> de.po, it_IT -> it.po, fr_FR -> fr.po
// Pure calculation.
func poFilename(lang string) string {
	parts := strings.SplitN(lang, "_", 2)
	return strings.ToLower(parts[0]) + ".po"
}

// poFilePath builds the full path for the PO file — pure calculation.
func poFilePath(outputDir, lang string) string {
	return filepath.Join(outputDir, poFilename(lang))
}

// parseExportArgs parses flags and positional args — calculation.
func parseExportArgs(args []string) (exportInput, error) {
	fs := flag.NewFlagSet("export", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(exportHelp) }

	if err := fs.Parse(args); err != nil {
		return exportInput{}, err
	}

	positional := fs.Args()
	if len(positional) < 3 {
		return exportInput{}, fmt.Errorf(
			"addon, lang, and output-dir are required - run 'glingoo export --help'",
		)
	}
	if len(positional) > 3 {
		return exportInput{}, fmt.Errorf(
			"unexpected argument %q - export takes exactly: <addon> <lang> <output-dir>",
			positional[3],
		)
	}

	return exportInput{
		addon:     positional[0],
		lang:      positional[1],
		outputDir: positional[2],
	}, nil
}

// buildExportResult shapes the data for the JSON response — pure calculation.
func buildExportResult(input exportInput, path string, created bool) map[string]any {
	return map[string]any{
		"addon":   input.addon,
		"lang":    input.lang,
		"path":    path,
		"created": created,
	}
}

// downloadPO downloads the PO file content from Odoo as raw bytes — side effect.
func downloadPO(client *godoorpc.Client, moduleID int, lang string) ([]byte, error) {
	// Step 1: create export wizard
	wizardID, err := client.ExecuteKW(
		"base.language.export", "create",
		godoorpc.Args{map[string]any{
			"lang":    lang,
			"format":  "po",
			"modules": []any{[]any{6, 0, []int{moduleID}}},
			"state":   "choose",
		}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		return nil, fmt.Errorf("could not create export wizard: %w", err)
	}

	id, ok := wizardID.(float64)
	if !ok {
		return nil, fmt.Errorf("unexpected wizard id type")
	}
	wid := int(id)

	// Step 2: execute wizard
	_, err = client.ExecuteKW(
		"base.language.export", "act_getfile",
		godoorpc.Args{[]int{wid}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		return nil, fmt.Errorf("could not execute export wizard: %w", err)
	}

	// Step 3: read the base64 encoded PO content
	result, err := client.ExecuteKW(
		"base.language.export", "read",
		godoorpc.Args{[]int{wid}},
		godoorpc.KWArgs{"fields": []string{"data"}},
	)
	if err != nil {
		return nil, fmt.Errorf("could not read export result: %w", err)
	}

	records, ok := result.([]any)
	if !ok || len(records) == 0 {
		return nil, fmt.Errorf("empty export result for addon")
	}

	record, ok := records[0].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unexpected export result format")
	}

	encoded, ok := record["data"].(string)
	if !ok || encoded == "" {
		return nil, fmt.Errorf("no translation data in export result")
	}

	return base64.StdEncoding.DecodeString(encoded)
}

// savePO writes the PO content to disk, creating directories as needed — side effect.
func savePO(path string, content []byte) (created bool, err error) {
	_, statErr := os.Stat(path)
	created = os.IsNotExist(statErr)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return false, fmt.Errorf("could not create directory %q: %w", filepath.Dir(path), err)
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return false, fmt.Errorf("could not write file %q: %w", path, err)
	}

	return created, nil
}

// RunExport executes the export command: downloads a PO file from Odoo and saves it.
func RunExport(args []string, conn ConnFlags) {
	input, err := parseExportArgs(args)
	if err == flag.ErrHelp {
		os.Exit(0)
	}
	if err != nil {
		write(errorPayload("export", err))
		os.Exit(1)
	}

	client, err := conn.Connect()
	if err != nil {
		write(errorPayload("export", fmt.Errorf("cannot connect to Odoo at %s - is Odoo running?", conn.URL)))
		os.Exit(1)
	}

	moduleID, err := findModuleID(client, input.addon)
	if err != nil {
		write(errorPayload("export", err))
		os.Exit(1)
	}

	content, err := downloadPO(client, moduleID, input.lang)
	if err != nil {
		write(errorPayload("export", err))
		os.Exit(1)
	}

	path := poFilePath(input.outputDir, input.lang)
	created, err := savePO(path, content)
	if err != nil {
		write(errorPayload("export", err))
		os.Exit(1)
	}

	write(successPayload("export", buildExportResult(input, path, created)))
}
