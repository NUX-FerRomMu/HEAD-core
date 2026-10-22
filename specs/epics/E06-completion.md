# E06 — Validation, review, integration and recovery

## Outcome

Only exact validated and human-approved work reaches the target, with recoverable failures.

## Scope and boundary

No approval bypass, invisible rebase or remote push.

Requirements: R04, R05, R06, R08, R09, R11, R13, R15.

External prerequisites: ADE-009, ADE-013, ADE-014, ADE-016, ADE-017, ADE-018, ADE-019. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-020 — Run exact-candidate validation and publish test evidence](../tickets/ADE-020-validation.md) (M).
- [ADE-021 — Anchor review discussions and enforce requested corrections](../tickets/ADE-021-review-comments.md) (M).
- [ADE-022 — Bind human approval to validated candidate identity](../tickets/ADE-022-review-approval.md) (S).
- [ADE-023 — Integrate approved candidates and reconcile Git journals](../tickets/ADE-023-integration.md) (L).
- [ADE-024 — Recover interrupted work and expose bounded artifacts](../tickets/ADE-024-recovery-artifacts.md) (M).

## Acceptance and demo

Pass/fail/stale/restart/ref-change scenarios preserve approval and ownership invariants. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
