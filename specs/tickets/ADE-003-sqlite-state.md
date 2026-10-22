# ADE-003 — Persist state, events, idempotency and durable jobs

- Epic: [E01](../epics/E01-foundation.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-002
- Evidence gate: none
- Requirements: R01, R03, R15

## Outcome

Committed domain actions survive restart and retries cannot duplicate work.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-002](./ADE-002-domain-model.md) implementation and evidence.

## Scope and implementation boundaries

- Own internal/store/sqlite migrations and transaction/repository ports, with WAL, foreign keys and busy timeout.
- Persist the records needed by contracts, mutation results/request hashes, global event cursors, jobs and integration/lease journals. Later feature migrations may add indexes/fields with reviewed compatibility.
- Provide transaction operations that save aggregate state, event, attention changes and jobs together; no external I/O inside transactions.
- Implement migration version checks and a consistent pre-migration backup strategy, not raw copying of a live WAL database.

Out of scope: No business scheduling, event delivery, speculative ORM or event-sourced replay engine.

## Acceptance criteria

- **AC1** — A forced failure between state and event/job writes rolls back all changes.
- **AC2** — Replaying a command ID with identical normalized input returns its original result; changed input returns CONFLICT even after restart.
- **AC3** — A stale expected_version rejects without any event, and concurrent valid contenders produce one winner.
- **AC4** — An interrupted migration leaves a recoverable prior database; an older binary rejects a newer schema with guidance.

## Verification

- Use temporary SQLite files, forced transaction failures and two contending goroutines.
- Restart store instances and verify jobs, results and cursors persist.
- Test backup restore and migration failure; run store tests with race detection.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
