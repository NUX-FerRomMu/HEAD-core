# ADE-017 — Assign roles and orchestrate ticket-scoped agents

- Epic: [E05](../epics/E05-agents.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-007, ADE-012, ADE-016, ADE-004
- Evidence gate: none
- Requirements: R04, R05, R06, R15

## Outcome

The human can explicitly assign and control ticket agents with durable, finite runs.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [provider contract](../contracts/provider.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-007](./ADE-007-workflow-commands.md) implementation and evidence.
- Accepted [ADE-012](./ADE-012-writer-leases.md) implementation and evidence.
- Accepted [ADE-016](./ADE-016-provider-port.md) implementation and evidence.
- Accepted [ADE-004](./ADE-004-local-rpc.md) implementation and evidence.

## Scope and implementation boundaries

- Implement ticket.assign/cancel and agent start/pause/resume/stop/get/list application services and RPC.
- Create logical agents and run records, snapshot profile/model/spec inputs, acquire developer leases and apply lifecycle facts.
- Require explicit tester/backlog assignment; start=false by default, optional start=true is the human start instruction.
- Serialize turns per agent, enforce installation concurrency limits, preserve pause intent, and distinguish run completion from logical developer closure.
- Treat tester start as a validation run on a committed candidate, and developer discussion in committed as read-only until a correction command authorizes new writing.

Out of scope: No automatic selection/assignment of agents, full chat delivery or real Codex-specific logic.

## Acceptance criteria

- **AC1** — An unassigned ticket cannot launch a provider; assigning a role does not start work unless explicitly requested.
- **AC2** — A developer process can finish its turn while its logical agent remains addressable for review; done/cancel closes it.
- **AC3** — Backlog/testing runs end after output; reassignment preserves history and requires the old role to be stopped.
- **AC4** — Cancellation that cannot confirm writer exit remains pending/uncertain and does not launch a replacement.
- **AC5** — Automatic correction resumes only an assigned developer and respects a human pause.

## Verification

- Drive ready→assigned→doing and pause/stop/reassign/cancel through RPC using fake providers.
- Test simultaneous starts, profile failures and process exit before candidate output.
- Restart queued/running jobs and assert no duplicate run or writer.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
