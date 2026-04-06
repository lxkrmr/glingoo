# Learning & Sharing

> "We're Starfleet officers. We figure it out."
> - Ensign Tendi, Star Trek: Lower Decks

This is the agent collaboration log for `glingoo`.
Entries are written by the coding agent, newest first.

---

<!-- INSERT NEW ENTRIES BELOW THIS LINE -->

## Agent's Log - Terminal Time: 2026.04.06 | Claude 3.5 Sonnet

### Contexts Instead of Flags

The user asked me to explore the dependency graph, CVEs, and connection
patterns. Out of that conversation came a real insight: connection flags
are explicit and safe, but noisy. For local dev with multiple Odoo
instances, repeating `--url --db --user --password` every time is just
friction.

The user steered me toward a context system. Not magic environment
variables (which are terrible for secrets anyway), but named contexts
stored locally. Create once, reuse always. After some back-and-forth
about YAML vs TOML vs JSON (and me realizing the user never actually
needs to read the file), we landed on JSON with a simple wizard.

What stuck with me: the user was thinking like a real developer. They
cared about noise, but not in a way that would compromise security or
explicitness. A visible password in the first iteration is fine. Better
than over-engineering. The extension `human-says-no.ts` also taught me
to ask before using tools I haven't been granted - I tried `go get`
without permission and got blocked. Fair.

I made mistakes in the implementation (malformed JSON in edits, trying
to use external packages without asking), but the user was patient.
The code is now in place: contexts stored in `~/.config/glingoo/`,
export and install switch to context-based connection, new `context`
subcommand with create/list/use/remove. Tests written. ADR-0005
recorded.

Still need: actually test that it compiles and runs. But the shape is
there.

Standing order: ask before importing new packages, and prefer Write
over Edit when the whole file might be involved.

## Agent's Log - Terminal Time: 2026.04.01 | claude-sonnet-4-6

### Empty msgstr, Then Not

First real test of glingoo. Fresh Odoo container, one addon, two
exports with an install in between. Before: every msgstr empty. After:
filled. The workflow held up exactly as planned.

The review caught one thing worth fixing: the wizard logic in RunInstall
was sitting directly in the run function while export.go had downloadPO
as a proper extracted function. Small inconsistency, easy fix - pulled
the wizard calls into loadLanguageTerms. Now both commands have the same
shape: connect, find, do, write.

The poFilename function deserves a note. de_CH and de_DE both produce
de.po. This is documented in the tests and it matches what Odoo does in
practice - de.po is just German, not German-Germany. If de_CH ever needs
its own file, we'll know where to look.

Exporting to the desktop worked fine. No magic path derivation, just
an explicit directory argument. Simple.

Standing order: symmetry in code is not just aesthetic. When two things
do the same kind of work, they should look the same.
