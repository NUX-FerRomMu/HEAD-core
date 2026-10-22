# E09 — Operational and self-hosting acceptance

## Outcome

The installed product develops a real change to itself and has measured recovery/usability evidence.

## Scope and boundary

No claim of completion from a fake demo alone.

Requirements: R01, R02, R03, R04, R05, R06, R07, R08, R09, R10, R11, R12, R13, R14, R15, R16.

External prerequisites: ADE-019, ADE-024, ADE-026, ADE-030. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-031 — Verify the complete headless and Neovim self-hosting flow](../tickets/ADE-031-e2e-workflow.md) (M).
- [ADE-032 — Package Linux installation and measure operational readiness](../tickets/ADE-032-release-operations.md) (M).
- [ADE-033 — Use ADE to develop ADE and record MVP acceptance](../tickets/ADE-033-self-hosting-rehearsal.md) (M).

## Acceptance and demo

Automated suite plus human live-provider rehearsal passes; limitations are explicit. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
