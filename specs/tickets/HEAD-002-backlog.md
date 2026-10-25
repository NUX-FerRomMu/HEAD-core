# HEAD-002 — Local backlog

Epic: E1. Depends on: HEAD-001. Blocked.

## Outcome and scope

A human can open a local Git project and create, edit and list tickets from `hea`. Follow [architecture](../../Docs/architecture.md) and [agent instructions](../AGENTS.md).

Implement project/ticket application services, SQLite persistence and CLI commands. A ticket has an opaque ID, project ID, title, description, acceptance criteria, draft/ready state and revision. Ready requires a nonempty title, description and at least one nonempty acceptance criterion. Reject execution-state transitions until that workflow exists.

Store settings and the database outside the repository, with a configurable location for tests. Identify a project by its canonical common Git directory. Separate clones are independent. CLI may call the application service directly until HEAD-004 introduces the daemon; serialize writes and use short transactions.

## Acceptance

1. Create a ticket, restart the CLI and retrieve the same contents and ID. Reopening the same project through another path does not duplicate it; a separate clone does.
2. Invalid readiness and stale-revision edits fail without partial writes. A missing project or malformed input produces an actionable error.
3. Nothing is written into the project checkout. `--json` produces machine-readable output without diagnostics mixed into stdout.

## Verification and handoff

Use temporary Git repositories and SQLite files to test persistence, identity and rejected mutations. Document only the introduced command syntax and record fields. Run Go tests and vet. Agent assignment, full Kanban lifecycle and epic decomposition come later.
