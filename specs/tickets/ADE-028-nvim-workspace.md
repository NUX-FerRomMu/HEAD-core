# ADE-028 — Add ticket Kanban, context and attention navigation

- Epic: [E08](../epics/E08-neovim.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-027, ADE-010, ADE-009, ADE-017
- Evidence gate: none
- Requirements: R03, R12, R13, R14, R04

## Outcome

The terminal editor centers navigation on tickets and their pending decisions.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [attention contract](../contracts/attention.md).
- [rpc contract](../contracts/rpc.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-027](./ADE-027-nvim-transport.md) implementation and evidence.
- Accepted [ADE-010](./ADE-010-context-projection.md) implementation and evidence.
- Accepted [ADE-009](./ADE-009-attention-actions.md) implementation and evidence.
- Accepted [ADE-017](./ADE-017-agent-lifecycle.md) implementation and evidence.

## Scope and implementation boundaries

- Implement project chooser, ticket/backlog and text Kanban buffers, factual context panel and persistent attention side column.
- Provide configurable :Ade subcommands and optional buffer-local Vim-style mappings for movement, next, preview, open and requeue.
- Provide ticket create/edit/readiness and explicit role/profile assignment actions through RPC.
- Refresh from versioned events and preserve useful cursor/view context; display empty/error/disconnected states.

Out of scope: No sophisticated graphics, model-generated UI sorting or mandatory keybindings.

## Acceptance criteria

- **AC1** — The user can switch tickets and immediately see state, last action, pending decisions and resource links.
- **AC2** — Queue preview leaves the main editing context intact; next opens top unread, and requeue returns it unread with Core-controlled priority.
- **AC3** — Any ticket/activity may be selected directly, even when not top-ranked.
- **AC4** — User-defined mappings override defaults and the plugin does not overwrite global mappings without opt-in.
- **AC5** — Kanban reflects Core state and never permits an arbitrary drag/move that bypasses guards.

## Verification

- Headless view/state tests with seeded ticket/queue snapshots and events.
- Manual smoke records split layout, narrow-terminal behavior and normal Vim navigation.
- Verify all actions map to existing RPC with no local domain transitions.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
