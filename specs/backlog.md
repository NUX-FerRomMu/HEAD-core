# Initial backlog

## Tickets

The first usable release delivers three independent capabilities: a ticket board, an attention queue, and Codex execution. Each feature includes its own persistence, commands and tests.

| Ticket | Deliverable | Dependency | Status |
|---|---|---|---|
| [HEAD-001](tickets/HEAD-001-command.md) | Go executable and shared local runtime | None | Ready |
| [HEAD-005](tickets/HEAD-005-kanban.md) | Persistent backlog and Kanban transitions | HEAD-001 | Planned |
| [HEAD-006](tickets/HEAD-006-notifications.md) | Persistent attention queue | HEAD-001 | Planned |
| [HEAD-007](tickets/HEAD-007-codex.md) | Run and cancel Codex tasks | HEAD-001 | Planned |
| [HEAD-008](tickets/HEAD-008-neovim.md) | Neovim board, notifications and agent output | HEAD-005, HEAD-006, HEAD-007 | Planned |

HEAD-002, HEAD-003 and HEAD-004 are superseded by these complete feature tickets; their original specs remain for reference.

## Branches and review

Start independent work from `develop`. While a required change is under review, branch from its implementation and name the prerequisite PR. HEAD-005, HEAD-006 and HEAD-007 share only HEAD-001; none depends on another feature. HEAD-008 combines them. Review status does not imply acceptance or permission to merge.

Execution instructions are in [AGENTS.md](../AGENTS.md).

## Epics

- Backlog and Kanban — HEAD-005.
- Attention queue — HEAD-006.
- Agent execution and provider adapters — HEAD-007.
- Git, testing and review.
- Neovim integration — HEAD-008.
- Self-hosting.
