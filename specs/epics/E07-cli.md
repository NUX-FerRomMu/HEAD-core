# E07 — Complete headless human control

## Outcome

The CLI can execute every required human operation using the same protocol as Neovim.

## Scope and boundary

No direct database shortcuts or secondary workflow engine.

Requirements: R01, R03, R04, R06, R09, R11, R12, R13, R14.

External prerequisites: ADE-004, ADE-010, ADE-013, ADE-024. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-025 — Expose project, backlog, context and review reads in CLI](../tickets/ADE-025-cli-read.md) (M).
- [ADE-026 — Expose all human execution, review and recovery commands](../tickets/ADE-026-cli-workflow.md) (M).

## Acceptance and demo

A full fake-provider journey succeeds through documented CLI commands. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
