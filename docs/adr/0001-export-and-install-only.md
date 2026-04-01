# ADR 0001: export and install only

## Status

Accepted

## Context

Otto included translate, add-known-translations, and load-language-terms.
In practice the daily workflow is just two operations:

1. Export a PO file from Odoo for an addon
2. Load language terms into Odoo after updating translations

add-known-translations is a convenience that copies translations between
addons. It requires PO file parsing and is rarely used in practice. YAGNI.

Currently only German (de_DE) is translated. A second language can be
added when there is a real need.

## Decision

`glingoo` implements exactly two commands: `export` and `install`.

## Consequences

Positive:
- minimal surface area to maintain
- covers the real daily workflow

Negative:
- add-known-translations is not available, must be done manually
