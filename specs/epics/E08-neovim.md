# E08 — Neovim ticket workspace

## Outcome

Lua supplies fast ticket, attention, review, chat and test navigation.

## Scope and boundary

No sophisticated GUI or business rules in Lua.

Requirements: R01, R03, R04, R06, R07, R09, R10, R11, R12, R13, R14, R15.

External prerequisites: ADE-004, ADE-009, ADE-010, ADE-014, ADE-017, ADE-018, ADE-019, ADE-020, ADE-021, ADE-022, ADE-023. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-027 — Build the Lua plugin transport and reconnecting state cache](../tickets/ADE-027-nvim-transport.md) (M).
- [ADE-028 — Add ticket Kanban, context and attention navigation](../tickets/ADE-028-nvim-workspace.md) (M).
- [ADE-029 — Review side-by-side snapshots and edit through human ownership](../tickets/ADE-029-nvim-review.md) (M).
- [ADE-030 — Add agent chat, interventions, testing and backlog proposal views](../tickets/ADE-030-nvim-agents-tests.md) (M).

## Acceptance and demo

Core workflow works in Neovim with configurable bindings and responsive views. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
