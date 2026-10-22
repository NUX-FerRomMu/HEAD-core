# ADE-016 — Implement provider contracts and a deterministic fake tool

- Epic: [E05](../epics/E05-agents.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-003, ADE-002
- Evidence gate: none
- Requirements: R06, R07, R04

## Outcome

Core orchestration can be developed and tested against a provider-neutral, deterministic lifecycle.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [provider contract](../contracts/provider.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-003](./ADE-003-sqlite-state.md) implementation and evidence.
- Accepted [ADE-002](./ADE-002-domain-model.md) implementation and evidence.

## Scope and implementation boundaries

- Implement the Go provider port, capability/profile registry, typed requests/events and validation.
- Build a fake provider executable/adapter with scripted messages, interventions, candidates, failures, hangs and cancellation.
- Implement role-restricted output dispatch and per-run sequence/idempotency validation.
- Keep tool-specific details in adapters and preserve unknown usage as null.

Out of scope: No real Codex integration or runtime plugin loader.

## Acceptance criteria

- **AC1** — The same port drives fake developer/backlog/tester scenarios without importing Codex types into application/domain.
- **AC2** — Malformed/out-of-order/duplicate events follow the contract and cannot produce duplicate state changes.
- **AC3** — A fake provider attempting human assignment/approval/integration is rejected by its narrow dispatcher.
- **AC4** — Cancellation and unsupported capability requests return explicit confirmed/uncertain/unsupported results.

## Verification

- Contract tests run all fake role outputs and capability permutations.
- Use scripted process errors and event fragmentation; test resource cleanup.
- Run provider port/fake tests offline with race detection where concurrent.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
