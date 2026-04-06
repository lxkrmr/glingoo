# ADR 0005: Contexts for connection management

## Status

Accepted (supersedes ADR-0004)

## Context

Connection via flags (ADR-0004) is secure and explicit, but creates
unnecessary noise. For developers working with multiple Odoo instances
locally, repeating `--url --db --user --password` on every command is
tedious. A context system allows connection details to be stored once
and reused.

## Decision

Introduce a `context` subcommand to manage named connection profiles:

```sh
glingoo context create mydev          # Interactive wizard
glingoo context list                  # Show all contexts
glingoo context use mydev             # Set as default
glingoo context remove mydev          # Delete context
```

Contexts are stored in `~/.config/glingoo/contexts.json`:

```json
{
  "contexts": {
    "mydev": {
      "url": "http://localhost:8069",
      "db": "mydb",
      "user": "admin",
      "password": "secret"
    }
  },
  "current_context": "mydev"
}
```

Export and install commands use the current context automatically:

```sh
glingoo export my_addon de_DE /path
glingoo install de_DE
```

## Consequences

Positive:
- Reduces noise: no repeated flags
- Supports multiple Odoo instances: can switch contexts easily
- Agent-friendly: credentials stay on user's system, not shared
- Explicit: `glingoo context list` shows what's active
- Minimal dependencies: JSON only, no external packages

Negative:
- Password stored in plain text JSON file (acceptable for local dev)
- User must run `context create` before first use
- Changes the mental model from "pure flags" to "stateful contexts"

## Notes

Context is stored locally under `~/.config/glingoo/` with permissions
restricted to the user. File is not meant to be version controlled or
shared. For development purposes with default credentials, the risk is
acceptable.
