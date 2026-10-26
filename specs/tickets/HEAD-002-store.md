# HEAD-002 — SQLite store

Dependency: HEAD-001. Scope: opening storage and applying migrations.

Add a store package using `database/sql` and a pinned pure-Go SQLite driver. Its constructor accepts a database path and returns an error rather than exiting. Add ordered SQL migrations and schema-version bookkeeping; no ticket tables yet.

## Acceptance

- A database at an explicit temporary path opens, closes and reopens successfully.
- Migrations apply once; reopening does not repeat them. A failed migration rolls back its changes and leaves the previous version intact.
- An inaccessible path produces a useful error. The store does not fall back to a database in the working directory.
- Store tests pass with `CGO_ENABLED=0`.

Verification: temporary-database tests for opening, migration replay and failure. Default path selection, CLI commands and daemon ownership are outside this ticket.
