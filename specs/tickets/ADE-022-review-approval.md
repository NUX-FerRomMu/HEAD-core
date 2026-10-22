# ADE-022 — Bind human approval to validated candidate identity

- Epic: [E06](../epics/E06-completion.md)
- Status: blocked; unassigned
- Size: S
- Dependencies: ADE-020, ADE-021, ADE-009
- Evidence gate: none
- Requirements: R04, R09, R11

## Outcome

Only a human approval for the exact validated candidate can enqueue integration.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [git contract](../contracts/git.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-020](./ADE-020-validation.md) implementation and evidence.
- Accepted [ADE-021](./ADE-021-review-comments.md) implementation and evidence.
- Accepted [ADE-009](./ADE-009-attention-actions.md) implementation and evidence.

## Scope and implementation boundaries

- Implement review.approve with transactional approval/job/attention mutation.
- Check ticket state, candidate/head/base/config, required passing validation, forced obligations, target ref and absence of a writer.
- Persist immutable approval scope and revoke eligibility when any reviewed input changes.
- Expose approval result and integration job ID without waiting for Git mutation.

Out of scope: No Git integration itself or bypass flag for failed validation.

## Acceptance criteria

- **AC1** — A human approves a current passing candidate with clear obligations and receives one approval/integration job.
- **AC2** — Provider-origin approval and missing/failed/stale validation each reject without an integration job.
- **AC3** — A new candidate, changed validation config or moved target makes the old approval ineligible.
- **AC4** — Retrying the same approval command returns the same result and does not duplicate jobs; competing stale approvals conflict.

## Verification

- Table-driven application tests vary each guard independently.
- Race approval against candidate/config/target changes and verify the worker will still recheck all guards.
- Verify queue resolution and approval/job insertion roll back together.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
