# E03 — Attention queue and context recovery

## Outcome

The user can trust a durable priority queue and resume ticket context quickly.

## Scope and boundary

Queue is advisory; resolving an item cannot bypass its domain action.

Requirements: R06, R12, R13, R14, R15.

External prerequisites: ADE-006, ADE-007. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-008 — Persist and rank the attention queue with aging](../tickets/ADE-008-attention-store.md) (M).
- [ADE-009 — Implement next, preview, requeue and guarded resolution](../tickets/ADE-009-attention-actions.md) (S).
- [ADE-010 — Build durable ticket context and resource summaries](../tickets/ADE-010-context-projection.md) (S).

## Acceptance and demo

Aging/requeue fixtures pass and context is usable without a model call. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
