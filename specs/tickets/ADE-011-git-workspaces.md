# ADE-011 — Create isolated ticket worktrees and configurable branches

- Epic: [E04](../epics/E04-git.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-005
- Evidence gate: none
- Requirements: R08, R02

## Outcome

Every implementation ticket receives a recoverable Git workspace isolated from other tickets.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-005](./ADE-005-projects-config.md) implementation and evidence.

## Scope and implementation boundaries

- Own internal/adapter/git process wrapper and workspace creation/reconciliation.
- Validate supported repositories, branch templates/refs, path boundaries and canonical ownership metadata.
- Journal creation intent before Git mutation and reconcile interrupted creation idempotently.
- Preserve all dirty/user-owned worktrees; expose structured Git errors and bounded command output.

Out of scope: No provider execution, review diff, branch rebasing or automatic cleanup of user files.

## Acceptance criteria

- **AC1** — Two tickets create distinct branches/worktrees from the configured target without changing the user checkout.
- **AC2** — A branch collision, invalid ref, escaping path or unsupported repository fails visibly without deleting anything.
- **AC3** — Crash after branch/worktree creation but before DB completion reconciles the same owned workspace on retry.
- **AC4** — Paths with spaces/Unicode work without shell interpolation and full object IDs remain intact.

## Verification

- Use real Git in temporary repositories and failure injection around journal boundaries.
- Test collision with an existing user branch and a partially created owned worktree.
- Run Git adapter integration tests; assert user branch/dirty files remain unchanged.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
