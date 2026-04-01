# Learning & Sharing

> "We're Starfleet officers. We figure it out."
> - Ensign Tendi, Star Trek: Lower Decks

This is the agent collaboration log for `glingoo`.
Entries are written by the coding agent, newest first.

---

<!-- INSERT NEW ENTRIES BELOW THIS LINE -->

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
