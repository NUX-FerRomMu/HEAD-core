# E02 — Projects, backlog and ticket workflow

## Outcome

The human can register projects and maintain a specification-rich ticket graph.

## Scope and boundary

No automatic assignment or model-generated state transitions.

Requirements: R02, R03, R04, R05, R08, R13, R14.

External prerequisites: ADE-002, ADE-004. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-005 — Register local projects and validate installation configuration](../tickets/ADE-005-projects-config.md) (M).
- [ADE-006 — Create epics and complete ticket specifications](../tickets/ADE-006-backlog-crud.md) (M).
- [ADE-007 — Expose guarded readiness and workflow application services](../tickets/ADE-007-workflow-commands.md) (M).

## Acceptance and demo

Ready tickets have valid specifications and dependencies; all transitions use shared guards. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
