# ADR 0003: JSON-only output

## Status

Accepted

## Context

Consistent with gindoo and gadoo. All tools in this family output JSON.
Machine-readable, agent-friendly, always pretty-printed.

## Decision

`glingoo` always outputs JSON. No text mode, no `--output` flag.

```json
{
  "ok": true,
  "command": "export",
  "data": {
    "addon": "my_addon",
    "lang": "de_DE",
    "path": "/path/to/my_addon/i18n/de.po",
    "created": true
  }
}
```

On error:

```json
{
  "ok": false,
  "command": "export",
  "error": "addon 'my_addon' not found in Odoo - check the name and make sure Odoo has scanned it"
}
```

Output is always pretty-printed. Error messages are specific and
actionable.

## Consequences

Positive:
- consistent with the tool family
- machine-readable and agent-friendly

Negative:
- requires jq for comfortable human reading of large outputs
