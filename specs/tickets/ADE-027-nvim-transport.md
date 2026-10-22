# ADE-027 — Build the Lua plugin transport and reconnecting state cache

- Epic: [E08](../epics/E08-neovim.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-004
- Evidence gate: none
- Requirements: R01, R14, R15

## Outcome

Neovim can connect to Core without blocking the editor or duplicating business rules.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-004](./ADE-004-local-rpc.md) implementation and evidence.

## Scope and implementation boundaries

- Create plugin/ade.lua and lua/ade modules with setup, configurable executable path, jobstart stdio bridge and framed RPC client.
- Implement request correlation, timeout/unknown-outcome handling, entity-version cache, event cursor deduplication and reconnect snapshots.
- Marshal UI mutations via vim.schedule and clean up jobs/autocommands on shutdown/reload.
- Set up pinned headless Neovim/Lua formatting tests; no dependency on a particular plugin manager.

Out of scope: No ticket transitions in Lua, editor-specific daemon state or full UI views.

## Acceptance criteria

- **AC1** — A minimal Neovim config starts the plugin and completes hello/status with the local daemon.
- **AC2** — Fragmented/combined frames, escaped newlines and large decimal cursor strings parse without loss.
- **AC3** — Daemon loss displays disconnected state; reconnect snapshots/replay do not duplicate entities or repeat a mutation with a fresh ID.
- **AC4** — Slow provider/event output does not block editor input and reload does not leak bridge jobs/autocommands.

## Verification

- Headless Lua tests with fake transport plus integration against the real bridge.
- Test connection teardown during an outstanding mutation and cursor expiration.
- Run stylua check and documented Neovim test harness.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
