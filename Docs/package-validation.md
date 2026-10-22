# Planning package validation

Date: 2026-10-22. This report concerns the design/specification package only; ADE has not been implemented or run by this task.

Validated with `python3 specs/check.py`:

- 33 unique implementation tickets, grouped exactly once into 9 epics.
- All 16 MVP requirements have ticket owners.
- Dependency graph is acyclic and all referenced ticket/gate IDs exist.
- Ticket status, dependencies and evidence gates agree between individual specs, Markdown backlog and JSON manifest.
- Local Markdown links resolve within the package.
- Shared queue examples match the specified formula/order; workflow examples have complete given/when/then structure.
- Only ADE-001 and ADE-015 are initially ready to assign.

The document checker also passed Python syntax compilation. The original attached Spanish source was copied unchanged to `Docs/source/ADE_CONTEXTO.md` and checked byte-for-byte.

Manual contract review clarified: Go/Lua ownership; explicit tester assignment; read-only developer discussion during committed state; human-only editing without a provider; candidate publication reservation; exact head/base/config validation; forced obligations across revisions; target-move reapproval; event replay; and uncertain-process recovery.

Not verified here: Go compilation, database driver/toolchain compatibility, actual Codex capabilities, application behavior, performance budgets, or a live self-hosting rehearsal. Those checks belong to their implementation tickets. Passing this checker is not evidence that the future application satisfies its specs.

The original archive contains only `Docs/` and `specs/`. Those directories are now prepared for review in `NUX-FerRomMu/HEAD-core`, with the repository handoff text updated and the checker rerun. Start at [README.md](README.md) and [the backlog](../specs/backlog.md).
