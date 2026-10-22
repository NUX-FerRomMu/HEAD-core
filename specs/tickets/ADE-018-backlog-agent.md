# ADE-018 — Review and atomically apply backlog-agent proposals

- Epic: [E05](../epics/E05-agents.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-017, ADE-006
- Evidence gate: none
- Requirements: R03, R06, R04

## Outcome

A finite backlog agent can decompose work into useful proposed tickets which the human explicitly accepts.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [provider contract](../contracts/provider.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-017](./ADE-017-agent-lifecycle.md) implementation and evidence.
- Accepted [ADE-006](./ADE-006-backlog-crud.md) implementation and evidence.

## Scope and implementation boundaries

- Validate and persist versioned backlog proposals with temporary references and source versions.
- Implement proposal preview/get, apply/reject and planning.complete, remapping local IDs atomically.
- Carry source links, constraints and acceptance criteria into child tickets; validate graph/project boundaries.
- Keep resulting tickets in backlog and unassigned; close finite backlog runs without losing proposals.

Out of scope: No autonomous backlog mutation, partial proposal acceptance UI, or model-based scheduling.

## Acceptance criteria

- **AC1** — A valid proposal with an epic, children and internal dependencies previews accurately and applies once on human command.
- **AC2** — A stale, cyclic, malformed or cross-project proposal inserts/updates no partial graph.
- **AC3** — Accepted children retain the required context and criteria and are not automatically ready or assigned.
- **AC4** — A planning ticket completes only after outstanding proposals are explicitly accepted/rejected; restart preserves proposal decisions.

## Verification

- Use fake backlog outputs with forward references and duplicate command retries.
- Inject a DB failure midway through graph insertion and assert full rollback.
- Test invalid source revisions and planning completion guards.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
