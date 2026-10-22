# ADE-008 — Persist and rank the attention queue with aging

- Epic: [E03](../epics/E03-attention.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-007
- Evidence gate: none
- Requirements: R12, R15

## Outcome

Pending human decisions are durable, deduplicated and ordered fairly over time.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [attention contract](../contracts/attention.md).
- [rpc contract](../contracts/rpc.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-007](./ADE-007-workflow-commands.md) implementation and evidence.

## Scope and implementation boundaries

- Implement Activity storage, creation/upsert/resolution port and attention.list.
- Use the exact score/tie formula with an injected clock and validated settings; snapshot pagination freezes evaluation time/order.
- Provide semantic-key deduplication and domain-transaction hooks for future event producers.
- Keep queue ranking in Core, with a SQL/index strategy adequate for the documented 10,000-item target.

Out of scope: No next/requeue actions, Neovim rendering or model-based ranking.

## Acceptance criteria

- **AC1** — The shared queue fixture has exactly the expected scores/order, including low aging past new critical.
- **AC2** — Deduplicating an unresolved occurrence preserves age/requeue count; replaying a resolved occurrence does not reopen it.
- **AC3** — Clock regression never gives negative age; ties resolve by creation time then ID.
- **AC4** — Restart preserves unresolved/seen state and list pagination does not reorder when time advances.

## Verification

- Load specs/fixtures/attention.json with a fake clock and table-driven edge cases.
- Test atomic activity creation with a domain mutation and rollback.
- Test stable paginated order and reopen/reconnect behavior on a temporary DB.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
