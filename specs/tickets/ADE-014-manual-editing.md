# ADE-014 — Hand worktrees to humans and preserve manual commits

- Epic: [E04](../epics/E04-git.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-013, ADE-012
- Evidence gate: none
- Requirements: R09, R10, R14

## Outcome

A user can pause agent writes, edit the ticket worktree, make their own commit and return ownership safely.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-013](./ADE-013-candidates-diffs.md) implementation and evidence.
- Accepted [ADE-012](./ADE-012-writer-leases.md) implementation and evidence.

## Scope and implementation boundaries

- Implement writer.acquire/release/commit for human clients with the process-stop and fence handshake.
- Expose editable workspace location only after lease grant; preserve snapshot read availability during waiting.
- Use configured human Git author identity; require explicit commit or separate user-controlled handling of dirty files before releasing to an agent.
- Audit manual commits and leave candidate publication explicit.
- Support manual-only ready→doing work without a provider assignment; acquiring ownership from committed moves to fixing and revokes approval before edits.

Out of scope: No destructive discard command or policing arbitrary external Git/editor processes.

## Acceptance criteria

- **AC1** — Requesting a human lease during an active writer does not permit editing until that process is confirmed stopped.
- **AC2** — A human commit records normal configured authorship and does not silently approve/publish/integrate.
- **AC3** — Releasing a dirty worktree rejects with actionable details and preserves files; old fences cannot commit.
- **AC4** — Review of an older immutable candidate works while the agent holds the live worktree lease.
- **AC5** — A ready implementation ticket with completed dependencies can be implemented entirely by a human; acquiring a writer for an already committed ticket invalidates its approval before any edit.

## Verification

- Integration tests combine process fixture, Git worktree and concurrent lease attempts.
- Test missing author configuration and rejected commit without a lease.
- Verify a manual commit can subsequently pass candidate.publish through the ordinary path.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
