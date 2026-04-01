package cmd

import (
	"fmt"

	"github.com/lxkrmr/godoorpc"
)

// findModuleID resolves an addon name to its ir.module.module record ID.
func findModuleID(client *godoorpc.Client, addon string) (int, error) {
	result, err := client.ExecuteKW(
		"ir.module.module", "search",
		godoorpc.Args{godoorpc.Domain{
			godoorpc.Condition{Field: "name", Op: "=", Value: addon},
		}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		return 0, fmt.Errorf("could not search for addon %q: %w", addon, err)
	}

	ids, ok := result.([]any)
	if !ok || len(ids) == 0 {
		return 0, fmt.Errorf(
			"addon %q not found in Odoo - check the name and make sure Odoo has scanned it",
			addon,
		)
	}

	id, ok := ids[0].(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected id type for addon %q", addon)
	}

	return int(id), nil
}
