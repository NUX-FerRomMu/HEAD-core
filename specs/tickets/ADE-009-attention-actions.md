# ADE-009 — Implement next, preview, requeue and guarded resolution

- Epic: [E03](../epics/E03-attention.md)
- Status: blocked; unassigned
- Size: S
- Dependencies: ADE-008
- Evidence gate: none
- Requirements: R12, R14

## Outcome

The user can choose, postpone and resolve attention without bypassing domain actions.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [attention contract](../contracts/attention.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-008](./ADE-008-attention-store.md) implementation and evidence.

## Scope and implementation boundaries

- Implement attention.next/seen/requeue/resolve with optimistic versions and command idempotency.
- Keep next advisory: mark the top unread item seen and return its typed resource/actions.
- Restrict direct resolution to permitted informational/acknowledgement activities; actionable outcomes resolve with domain success.
- Provide project-scoped and explicit all-project queue behavior, plus seen counts when no unread item remains.

Out of scope: No automatic execution of next actions or UI shortcuts.

## Acceptance criteria

- **AC1** — Two simultaneous next calls select distinct unread items, or one returns null when only one existed.
- **AC2** — Requeue returns an unresolved item to unread, applies capped penalty and preserves original age; retry does not increment twice.
- **AC3** — Preview/list does not mark seen, and users may select any item without changing priority policy.
- **AC4** — Direct attention.resolve on review_ready cannot approve or merge a ticket; failed domain action leaves its activity pending.

## Verification

- Use fake-clock formula cases and concurrent SQLite calls.
- Test idempotent actions and rejected resolution for every actionable kind.
- Exercise project/all filters through RPC.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
