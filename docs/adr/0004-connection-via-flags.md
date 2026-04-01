# ADR 0004: Connection via flags

## Status

Accepted

## Context

Consistent with gindoo and gadoo. No config file, no environment
variables for secrets. Connection flags come before the command.

## Decision

Connection credentials are passed as CLI flags:
`--url`, `--db`, `--user`, `--password`.

A shell alias is recommended for human use. Coding assistants
construct the full command directly.

## Consequences

Positive:
- no config file to manage
- consistent with the tool family
- secure: no env var leakage

Negative:
- flags must be repeated without an alias
