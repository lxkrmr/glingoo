# ADR 0002: Explicit arguments, no magic

## Status

Accepted

## Context

An earlier design derived the addon name and language from the output
path (e.g. inferring "my_addon" from "/path/to/my_addon/i18n/de.po").

This was rejected because it introduced implicit behaviour: the tool
would silently derive things the caller did not state. A caller who
puts the file on the desktop, or who has a directory name that differs
from the technical addon name, would get wrong results.

The primary caller is a coding assistant constructing commands from
context - verbosity is not a burden for an agent.

## Decision

All arguments are explicit:

    glingoo export my_addon de_DE /path/to/my_addon/i18n
    glingoo install de_DE

- `export` takes addon name, language code, and output directory.
  The filename (e.g. `de.po`) is derived from the language code
  following Odoo's PO file naming convention.
- `install` takes the language code.

## Consequences

Positive:
- no surprise behaviour from path conventions
- works regardless of directory naming
- caller controls where the file goes

Negative:
- more arguments to type (acceptable for agent-driven usage)
