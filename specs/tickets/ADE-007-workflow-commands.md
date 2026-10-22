# ADE-007 — Expose guarded readiness and workflow application services

- Epic: [E02](../epics/E02-planning.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-006, ADE-002
- Evidence gate: none
- Requirements: R03, R04, R05

## Outcome

Core applies domain transitions consistently from every command source.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-006](./ADE-006-backlog-crud.md) implementation and evidence.
- Accepted [ADE-002](./ADE-002-domain-model.md) implementation and evidence.

## Scope and implementation boundaries

- Implement ticket.ready/unready and reusable application dispatch for transition commands and domain facts.
- Load dependency states, entity versions and trusted actor capability from the dispatch path; never trust a JSON actor field.
- Expose internal supervisor command ports for later candidate/validation/integration work without granting human-only actions.
- Persist transition events/jobs transactionally and provide explicit invalid-state details.

Out of scope: No implementation of later agent, validation or Git side effects; those callers consume this service.

## Acceptance criteria

- **AC1** — A draft missing description/criteria cannot become ready; a valid draft can and unready restores backlog when unassigned.
- **AC2** — Assignment policy rejects unfinished/cancelled dependencies and role-kind mismatches when invoked through the internal service.
- **AC3** — A rejected transition persists no new state/event/job; replay of an accepted command returns the same result.
- **AC4** — The same pure guard is used by human and supervisor entry points, with different permitted actions.

## Verification

- Test RPC readiness and internal future-workflow dispatch with stored entity facts.
- Cover stale versions, unauthorized source and duplicate command IDs.
- Run application/state integration tests; do not fake a complete provider lifecycle.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
