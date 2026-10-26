# HEAD-002 — SQLite ticket repository

Dependency: HEAD-001. Scope: a working ticket repository, independent of the service and CLI.

Add a store package using `database/sql` and a pinned pure-Go SQLite driver. Implement `Open`, `Close` and `application.TicketRepository` from the [ticket interface](../ticket-interface.md). Its constructor accepts a database path and returns an error rather than exiting. Add ordered SQL migrations, schema-version bookkeeping and the ticket table.

Owned paths: `internal/store/`, `go.mod` and `go.sum` for the SQLite dependency. HEAD-003 and HEAD-004 can proceed concurrently; this ticket needs neither implementation.

## Acceptance

- A database at an explicit temporary path opens, closes and reopens successfully.
- Migrations apply once; reopening does not repeat them. A failed migration rolls back its changes and leaves the previous version intact.
- An inaccessible path produces a useful error. The store does not fall back to a database in the working directory.
- Insert two records directly through the repository, close and reopen it, and list the original IDs, content and insertion order. Duplicate IDs fail without overwriting data; failed inserts leave no partial record.
- Store tests pass with `CGO_ENABLED=0`.

Verification: temporary-database tests for opening, migration replay/failure and ticket persistence; compile-time interface conformance. Default path selection, title validation, ID generation, CLI commands and daemon ownership are outside this ticket.
