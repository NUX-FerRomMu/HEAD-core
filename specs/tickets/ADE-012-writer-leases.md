# ADE-012 — Coordinate fenced writers and supervised processes

- Epic: [E04](../epics/E04-git.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-011, ADE-003, ADE-008
- Evidence gate: none
- Requirements: R10, R15

## Outcome

Core can prove that one ADE-controlled writer owns a ticket and can stop it before ownership changes.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [domain contract](../contracts/domain.md).
- [provider contract](../contracts/provider.md).
- Accepted [ADE-011](./ADE-011-git-workspaces.md) implementation and evidence.
- Accepted [ADE-003](./ADE-003-sqlite-state.md) implementation and evidence.
- Accepted [ADE-008](./ADE-008-attention-store.md) implementation and evidence.

## Scope and implementation boundaries

- Implement persisted writer lease/fence and repository integration lease ports.
- Implement process launch tracking, bounded graceful/forced process-tree termination and identity-aware reconciliation on Linux.
- Use contexts and nonblocking bounded output readers; prevent goroutine leaks and stale PID kills.
- A lease with uncertain process termination blocks new writers and emits recoverable attention through the queue port.

Out of scope: No claim of blocking arbitrary external editors; no automatic provider assignment.

## Acceptance criteria

- **AC1** — Two concurrent writer acquisitions produce one owner; stale fence writes are rejected after handoff.
- **AC2** — An expired heartbeat alone cannot transfer ownership while the old process is alive.
- **AC3** — Cancellation terminates child processes and drains/waits for readers; inability to confirm exit leaves lease uncertain.
- **AC4** — Restart cannot kill an unrelated process whose PID was reused, and does not grant a writer until old ownership is resolved.

## Verification

- Spawn controlled child/grandchild fixtures with ignored termination and bounded escalation.
- Use fake process identity for PID reuse and crash reconciliation cases.
- Run race tests over acquire/release/start/stop and check goroutine/process counts.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
