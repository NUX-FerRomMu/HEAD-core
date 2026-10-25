# Backlog

## First increment

A human can create local tickets, navigate a persistent attention queue, and use both from Neovim. This establishes the foundation before adding agent execution.

| Ticket | Outcome | Depends on | Status |
|---|---|---|---|
| [HEAD-001](tickets/HEAD-001-bootstrap.md) | Go executable and development checks | — | Ready, unassigned |
| [HEAD-002](tickets/HEAD-002-backlog.md) | Persistent local projects and tickets through CLI | HEAD-001 | Blocked |
| [HEAD-003](tickets/HEAD-003-attention.md) | Attention ordering, next and requeue | HEAD-002 | Blocked |
| [HEAD-004](tickets/HEAD-004-neovim.md) | Daemon/bridge and thin Neovim views | HEAD-003 | Blocked |

Assign HEAD-001 first. Dependencies mean accepted work, not merely started work. Use `Ready → In progress → Review → Done`, with Blocked for unmet dependencies; update this table rather than maintaining a duplicate manifest.

## Epics

| Epic | Outcome | Scope |
|---|---|---|
| E1 — Foundation | Usable local backlog | HEAD-001 and HEAD-002 |
| E2 — Attention and frontend | Ticket-centered navigation | HEAD-003 and HEAD-004 |
| E3 — Agents | Explicit assignment, finite roles, chat and accepted backlog proposals | Verify one chosen tool's actual capabilities; then specify its adapter and lifecycle. |
| E4 — Development workflow | Isolated worktrees, testing, review, manual editing and approved integration | Specify one end-to-end ticket flow, including failed validation and interrupted writes. |
| E5 — Self-hosting | Develop a real HEAD change using HEAD | Complete the flow with a real provider, human review and a restart. |

Refine E3–E5 into small specs when their prerequisites exist. Keep provider choice, detailed RPC, schema design and recovery mechanisms with the work that needs them. Later platforms, hosted collaboration and richer extension support remain outside this increment.
