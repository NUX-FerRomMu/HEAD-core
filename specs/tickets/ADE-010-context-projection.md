# ADE-010 — Build durable ticket context and resource summaries

- Epic: [E03](../epics/E03-attention.md)
- Status: blocked; unassigned
- Size: S
- Dependencies: ADE-008, ADE-006
- Evidence gate: none
- Requirements: R13, R06

## Outcome

A client can restore a ticket workspace from one factual projection without invoking a model.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [attention contract](../contracts/attention.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-008](./ADE-008-attention-store.md) implementation and evidence.
- Accepted [ADE-006](./ADE-006-backlog-crud.md) implementation and evidence.

## Scope and implementation boundaries

- Implement context.get with state, last action, current block, pending actions, agents, candidate, tests, comments and typed resource references.
- Return absent sections explicitly as empty/null until those domain records exist, not fabricated placeholder results.
- Attribute optional agent narratives with context revision/timestamp and keep factual summary authoritative.
- Read the projection at one DB snapshot/version; expose events that trigger client refresh.

Out of scope: No automatic LLM summarization or storage of Neovim window layout in Core.

## Acceptance criteria

- **AC1** — After creating a ticket and attention activity, context returns the same coherent state before and after restart.
- **AC2** — A stale narrative cannot override current ticket status or newer validation failure.
- **AC3** — A ticket with no agent/candidate/tests yields an intelligible empty-state summary.
- **AC4** — Missing artifacts/resources return unavailable references without breaking the entire projection.

## Verification

- Seed realistic domain records at multiple lifecycle stages and compare the projection.
- Assert context generation makes no provider calls.
- Test concurrent updates do not produce mixed-version state.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
