# Contributing

## Commits

Use Conventional Commits.

Format:

```text
type(scope): short description
```

Examples:

```text
feat(export): add export command
fix(install): return clear error when language not found
docs(adr): add decision for explicit arguments
refactor(cmd): extract savePoFile helper
test(export): cover missing i18n directory case
```

Rules:
- keep commits small and meaningful
- write commit messages in English
- prefer one focused change per commit
- use a scope that matches the main area you changed

Common types:
- `feat`
- `fix`
- `docs`
- `refactor`
- `test`
- `chore`
