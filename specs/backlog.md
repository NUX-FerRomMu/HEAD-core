# Initial backlog

## Tickets

These five tickets establish a small, persistent ticket list through the CLI. They do not include project registration, workflow automation, the daemon or a UI.

| Ticket | Deliverable | Dependency | Status |
|---|---|---|---|
| [HEAD-001](tickets/HEAD-001-command.md) | Go command and shared ticket interfaces | None | Ready |
| [HEAD-002](tickets/HEAD-002-store.md) | SQLite ticket repository | HEAD-001 | Blocked |
| [HEAD-003](tickets/HEAD-003-tickets.md) | Ticket creation and listing logic | HEAD-001 | Blocked |
| [HEAD-004](tickets/HEAD-004-cli.md) | Ticket command parsing and output | HEAD-001 | Blocked |
| [HEAD-005](tickets/HEAD-005-wire-cli.md) | Working CLI backed by SQLite | HEAD-002, HEAD-003, HEAD-004 | Blocked |

All tickets are unassigned. A dependency is satisfied when its work is accepted. Execution instructions are in [AGENTS.md](../AGENTS.md).

## Parallel work

After HEAD-001, HEAD-002, HEAD-003 and HEAD-004 can run concurrently in separate branches. They share the small [ticket interface](ticket-interface.md), not each other's implementation. HEAD-005 joins the three completed pieces and verifies the executable end to end.

| Ticket | Owned paths | Independent verification |
|---|---|---|
| HEAD-002 | `internal/store/`, SQLite dependency in `go.mod` / `go.sum` | Real temporary databases |
| HEAD-003 | `internal/application/ticket_service.go` and its tests | In-memory repository test double |
| HEAD-004 | `internal/cli/` | Fake ticket service and captured output |

The shared types are delivered by HEAD-001; parallel tickets consume them unchanged. Any necessary interface correction is agreed and merged separately before dependent changes. Their actual command wiring belongs to HEAD-005.

## Epics

- Core and local backlog — current tickets.
- Attention queue.
- Agent coordination and provider adapters.
- Git, testing and review.
- Neovim integration.
- Self-hosting.
