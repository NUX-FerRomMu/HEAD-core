# Initial backlog

## Tickets

These four tickets establish a small, persistent ticket list through the CLI. They do not include project registration, workflow automation, the daemon or a UI.

| Ticket | Deliverable | Dependency | Status |
|---|---|---|---|
| [HEAD-001](tickets/HEAD-001-command.md) | Go module and `hea` help/version | None | Ready |
| [HEAD-002](tickets/HEAD-002-store.md) | Open SQLite and run migrations | HEAD-001 | Blocked |
| [HEAD-003](tickets/HEAD-003-tickets.md) | Create and list stored tickets | HEAD-002 | Blocked |
| [HEAD-004](tickets/HEAD-004-cli.md) | CLI commands for those two operations | HEAD-003 | Blocked |

All tickets are unassigned. A dependency is satisfied when its work is accepted. Execution instructions are in [AGENTS.md](../AGENTS.md).

## Epics

- Core and local backlog — current tickets.
- Attention queue.
- Agent coordination and provider adapters.
- Git, testing and review.
- Neovim integration.
- Self-hosting.
