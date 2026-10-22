# E04 — Git workspaces and controlled writing

## Outcome

Concurrent work stays isolated and reviews remain bound to immutable snapshots.

## Scope and boundary

No unreviewed integration or destructive automatic cleanup.

Requirements: R02, R05, R08, R09, R10, R14, R15.

External prerequisites: ADE-003, ADE-005, ADE-007, ADE-008. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-011 — Create isolated ticket worktrees and configurable branches](../tickets/ADE-011-git-workspaces.md) (M).
- [ADE-012 — Coordinate fenced writers and supervised processes](../tickets/ADE-012-writer-leases.md) (M).
- [ADE-013 — Publish immutable candidates and exact snapshot diffs](../tickets/ADE-013-candidates-diffs.md) (M).
- [ADE-014 — Hand worktrees to humans and preserve manual commits](../tickets/ADE-014-manual-editing.md) (M).

## Acceptance and demo

Lease handoff prevents concurrent ADE writers; candidates and manual commits are inspectable. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
