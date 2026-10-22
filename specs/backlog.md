# ADE implementation backlog

33 implementation tickets across 9 epics. **Start with ADE-001 and ADE-015.** All are unassigned. The user chooses assignment; an agent takes one ticket at a time. All P0 entries are required for the scoped MVP; ordering is determined by dependencies, not the number alone.

Read [execution instructions](AGENTS.md) and [delivery milestones](../Docs/delivery.md). The [JSON graph](backlog.json) has the same IDs/dependencies/gates. A ready ticket can start; blocked tickets become ready only after accepted prerequisites and a satisfied evidence gate. Sizes are relative S/M/L; do not treat them as delivery dates.

| Ticket | Epic | Outcome | Size | Status | Dependencies | Gate |
|---|---|---|---|---|---|---|
| [ADE-001](tickets/ADE-001-go-bootstrap.md) | E01 | Bootstrap the Go repository and local verification | S | ready | — | — |
| [ADE-002](tickets/ADE-002-domain-model.md) | E01 | Implement domain records, invariants and transition rules | M | blocked | ADE-001 | — |
| [ADE-003](tickets/ADE-003-sqlite-state.md) | E01 | Persist state, events, idempotency and durable jobs | M | blocked | ADE-002 | — |
| [ADE-004](tickets/ADE-004-local-rpc.md) | E01 | Serve versioned local RPC with replay and stdio bridge | M | blocked | ADE-003 | — |
| [ADE-005](tickets/ADE-005-projects-config.md) | E02 | Register local projects and validate installation configuration | M | blocked | ADE-004 | — |
| [ADE-006](tickets/ADE-006-backlog-crud.md) | E02 | Create epics and complete ticket specifications | M | blocked | ADE-005 | — |
| [ADE-007](tickets/ADE-007-workflow-commands.md) | E02 | Expose guarded readiness and workflow application services | M | blocked | ADE-006, ADE-002 | — |
| [ADE-008](tickets/ADE-008-attention-store.md) | E03 | Persist and rank the attention queue with aging | M | blocked | ADE-007 | — |
| [ADE-009](tickets/ADE-009-attention-actions.md) | E03 | Implement next, preview, requeue and guarded resolution | S | blocked | ADE-008 | — |
| [ADE-010](tickets/ADE-010-context-projection.md) | E03 | Build durable ticket context and resource summaries | S | blocked | ADE-008, ADE-006 | — |
| [ADE-011](tickets/ADE-011-git-workspaces.md) | E04 | Create isolated ticket worktrees and configurable branches | M | blocked | ADE-005 | — |
| [ADE-012](tickets/ADE-012-writer-leases.md) | E04 | Coordinate fenced writers and supervised processes | M | blocked | ADE-011, ADE-003, ADE-008 | — |
| [ADE-013](tickets/ADE-013-candidates-diffs.md) | E04 | Publish immutable candidates and exact snapshot diffs | M | blocked | ADE-012, ADE-007 | — |
| [ADE-014](tickets/ADE-014-manual-editing.md) | E04 | Hand worktrees to humans and preserve manual commits | M | blocked | ADE-013, ADE-012 | — |
| [ADE-015](tickets/ADE-015-codex-compatibility.md) | E05 | Verify the installed Codex structured integration surface | S | ready | — | — |
| [ADE-016](tickets/ADE-016-provider-port.md) | E05 | Implement provider contracts and a deterministic fake tool | M | blocked | ADE-003, ADE-002 | — |
| [ADE-017](tickets/ADE-017-agent-lifecycle.md) | E05 | Assign roles and orchestrate ticket-scoped agents | M | blocked | ADE-007, ADE-012, ADE-016, ADE-004 | — |
| [ADE-018](tickets/ADE-018-backlog-agent.md) | E05 | Review and atomically apply backlog-agent proposals | M | blocked | ADE-017, ADE-006 | — |
| [ADE-019](tickets/ADE-019-codex-chat.md) | E05 | Integrate verified Codex sessions, chat and interventions | L | blocked | ADE-015, ADE-017 | codex_interface_verified |
| [ADE-020](tickets/ADE-020-validation.md) | E06 | Run exact-candidate validation and publish test evidence | M | blocked | ADE-013, ADE-017, ADE-016 | — |
| [ADE-021](tickets/ADE-021-review-comments.md) | E06 | Anchor review discussions and enforce requested corrections | M | blocked | ADE-013, ADE-017 | — |
| [ADE-022](tickets/ADE-022-review-approval.md) | E06 | Bind human approval to validated candidate identity | S | blocked | ADE-020, ADE-021, ADE-009 | — |
| [ADE-023](tickets/ADE-023-integration.md) | E06 | Integrate approved candidates and reconcile Git journals | L | blocked | ADE-022, ADE-014 | — |
| [ADE-024](tickets/ADE-024-recovery-artifacts.md) | E06 | Recover interrupted work and expose bounded artifacts | M | blocked | ADE-023, ADE-019, ADE-018 | — |
| [ADE-025](tickets/ADE-025-cli-read.md) | E07 | Expose project, backlog, context and review reads in CLI | M | blocked | ADE-010, ADE-013, ADE-004 | — |
| [ADE-026](tickets/ADE-026-cli-workflow.md) | E07 | Expose all human execution, review and recovery commands | M | blocked | ADE-025, ADE-024 | — |
| [ADE-027](tickets/ADE-027-nvim-transport.md) | E08 | Build the Lua plugin transport and reconnecting state cache | M | blocked | ADE-004 | — |
| [ADE-028](tickets/ADE-028-nvim-workspace.md) | E08 | Add ticket Kanban, context and attention navigation | M | blocked | ADE-027, ADE-010, ADE-009, ADE-017 | — |
| [ADE-029](tickets/ADE-029-nvim-review.md) | E08 | Review side-by-side snapshots and edit through human ownership | M | blocked | ADE-028, ADE-014, ADE-021, ADE-022, ADE-023 | — |
| [ADE-030](tickets/ADE-030-nvim-agents-tests.md) | E08 | Add agent chat, interventions, testing and backlog proposal views | M | blocked | ADE-029, ADE-019, ADE-020, ADE-018 | — |
| [ADE-031](tickets/ADE-031-e2e-workflow.md) | E09 | Verify the complete headless and Neovim self-hosting flow | M | blocked | ADE-026, ADE-030, ADE-024 | — |
| [ADE-032](tickets/ADE-032-release-operations.md) | E09 | Package Linux installation and measure operational readiness | M | blocked | ADE-031 | — |
| [ADE-033](tickets/ADE-033-self-hosting-rehearsal.md) | E09 | Use ADE to develop ADE and record MVP acceptance | M | blocked | ADE-032, ADE-019 | real_codex_authenticated |

## Epics

- [E01 — Core foundations and protocol](epics/E01-foundation.md)
- [E02 — Projects, backlog and ticket workflow](epics/E02-planning.md)
- [E03 — Attention queue and context recovery](epics/E03-attention.md)
- [E04 — Git workspaces and controlled writing](epics/E04-git.md)
- [E05 — Provider-neutral agents and Codex integration](epics/E05-agents.md)
- [E06 — Validation, review, integration and recovery](epics/E06-completion.md)
- [E07 — Complete headless human control](epics/E07-cli.md)
- [E08 — Neovim ticket workspace](epics/E08-neovim.md)
- [E09 — Operational and self-hosting acceptance](epics/E09-self-hosting.md)

## Evidence gates

- `codex_interface_verified`: ADE-015 records a reproducible supported structured integration. A negative feasibility report can complete the investigation while leaving this gate unsatisfied and ADE-019 blocked.
- `real_codex_authenticated`: ADE-019 records a successful real-tool smoke; tool/auth availability must also hold at rehearsal time. Fake-provider tests cannot satisfy it.

## Deferred backlog

These are themes, not agent-ready implementation tickets. Refine them into specs only after a scope decision.

- **D01** — macOS/Windows transport, process and permissions support.
- **D02** — Remote hosting and issue/PR tracker adapters.
- **D03** — Multi-user collaboration and remote authentication.
- **D04** — Dynamic extension protocol and additional frontends/providers.
- **D05** — Advanced indexing, persistent provider memory and context compaction.
- **D06** — Submodules, LFS, bare/unborn repositories and relocation.
- **D07** — Partial proposal acceptance and richer backlog editing.

## Updating the schedule

After human/reviewer acceptance, mark the ticket done in its file and both backlog views. Recompute dependencies and gates for downstream tickets. Keep blocked reasons explicit. Never replace a missing prerequisite with an unreviewed stub. Use specs/check.py to verify links, graph structure, coverage and status consistency.
