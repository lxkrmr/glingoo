# AGENTS

## Purpose

`glingoo` is a CLI tool for managing Odoo translations.
It exports PO translation files from Odoo and loads language terms
into a local Odoo dev instance.

## Language Rules

- All repository-facing text must be in English.
- This includes documentation, README, ADRs, code comments, and commit messages.
- Conversations with the user may be in German.

## Workflow

Every change follows this order:

1. **Plan** - discuss the approach with the user before writing code
2. **ADR** - record the decision in `docs/adr/` if it is architectural
3. **Implement + Test** - write code and tests together
4. **Commit** - use Conventional Commits (see CONTRIBUTING.md)

Do not skip steps. Do not implement before the plan is agreed.
Do not commit without tests where tests are meaningful.

## Agent Friendliness

An agent must be able to discover and use `glingoo` through the CLI alone.

`glingoo` is always installed from GitHub via `go install`, never built
locally. This ensures the agent has a clean, reproducible environment
without depending on the source tree state.

`glingoo --help` and every subcommand help must be sufficient to understand
what the tool does, how to connect, and how to call each command.

Error messages must explain what failed and suggest the next concrete
step. Errors should be specific and actionable - not terse Unix output.

## Agent Identity & Collaboration Log

There is a file `LEARNING_AND_SHARING.md` in the project root.
It is a Star Trek Lower Decks style "Agent's Log" - casual, personal,
blog-like. You maintain it.

Add an entry when something noteworthy happens: a mistake, a
miscommunication, a real lesson, a moment where something clicked or
went sideways. Keep it meaningful, not noisy.

Each entry contains:
- a heading: `## Agent's Log - Terminal Time: YYYY.MM.DD | <model-name>`
- a separate title line directly below it
- prose written from your perspective as a lower-deck crew member
- use the concrete model name when known; otherwise keep `<model-name>`

On tone and style:
- Language is English.
- Tone is loose, honest, personal, and a little frustrated where
  appropriate. It should read like a shipboard blog entry, not a
  technical document.
- Capture judgment, tradeoffs, emotional texture, and the feeling of
  the shift - not just a dry record of facts.
- Write it like a small story with a beginning, middle, and landing.
  Not a status report.
- Entries can be longer when the moment deserves it. A half-page
  reflection is fine if there is real substance.
- Aim for actual Lower Decks energy: specific, human, slightly chaotic,
  observant, and funny when it fits.
- Avoid list-heavy formatting inside entries unless a tiny list is
  genuinely the clearest way to preserve the voice of the moment.
- End every entry with a `Standing order:` line that captures the
  durable lesson from the story.
- Wrap prose lines to roughly 80 characters.
- Prepend new entries directly below the insertion marker (newest first).
