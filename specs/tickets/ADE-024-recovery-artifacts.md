# ADE-024 — Recover interrupted work and expose bounded artifacts

- Epic: [E06](../epics/E06-completion.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-023, ADE-019, ADE-018
- Evidence gate: none
- Requirements: R15, R06, R13

## Outcome

Restarting ADE leaves durable work inspectable and never accidentally starts a second writer.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [provider contract](../contracts/provider.md).
- [git contract](../contracts/git.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-023](./ADE-023-integration.md) implementation and evidence.
- Accepted [ADE-019](./ADE-019-codex-chat.md) implementation and evidence.
- Accepted [ADE-018](./ADE-018-backlog-agent.md) implementation and evidence.

## Scope and implementation boundaries

- Coordinate startup reconciliation across jobs, agents, leases, worktrees, validations and integration journals.
- Implement bounded artifact registration/read with ID checks, offsets/checksums and no arbitrary path exposure.
- Add structured correlation logging, redacted diagnostics, disk-full handling and graceful daemon shutdown.
- Apply bounded transient retries and deduplicated recovery activities; preserve uncertain leases and independently usable projects.

Out of scope: No general telemetry service, automatic data destruction or guarantee of sandboxing same-user external tools.

## Acceptance criteria

- **AC1** — Restart during developer/test/backlog/integration work yields the correct interrupted/reconciled state without duplicated output/actions.
- **AC2** — Unknown process ownership blocks only affected writes; another project and read-only review still work.
- **AC3** — Artifact reads reject traversal/unregistered IDs/out-of-range lengths and log truncation remains visible.
- **AC4** — Disk-full/SQLite failure rolls back state/event changes and yields actionable diagnostics without deleting user data.
- **AC5** — Graceful shutdown terminates/records processes and leaves a clean restart path.

## Verification

- Subprocess kill/restart harness at each persisted job boundary with deterministic fake providers.
- Test oversized logs, unavailable artifacts and injected disk/storage errors.
- Exercise bounded retry exhaustion and repeated startup deduplication.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
