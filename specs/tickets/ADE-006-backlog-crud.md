# ADE-006 — Create epics and complete ticket specifications

- Epic: [E02](../epics/E02-planning.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-005
- Evidence gate: none
- Requirements: R03, R13

## Outcome

The human can persist and browse an epic/ticket graph that provides enough specification context for agents.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-005](./ADE-005-projects-config.md) implementation and evidence.

## Scope and implementation boundaries

- Implement epic and ticket create/get/list/update RPC and storage queries, pagination and filters.
- Persist descriptions, acceptance criteria, parent/epic/dependency links and source context; validate project membership and acyclicity.
- Enforce immutable execution fields after assignment while permitting allowed metadata edits.
- Provide Kanban projections from the authoritative ticket state, with epic progress derived from children.

Out of scope: No agent refinement, role assignment, drag-and-drop UI or generic project-management customization.

## Acceptance criteria

- **AC1** — Creating an epic with child tickets and dependencies survives restart and returns stable list ordering.
- **AC2** — A cyclic, missing or cross-project parent/dependency reference rejects the whole mutation.
- **AC3** — Stale edits return CONFLICT; execution-spec edits to an assigned ticket return INVALID_STATE without losing metadata.
- **AC4** — Kanban counts and epic progress reflect stored ticket transitions rather than client-maintained state.

## Verification

- Exercise create/read/update/list over RPC against SQLite.
- Test dependency cycles of length two and greater, pagination consistency and title-only drafts.
- Verify a draft may exist in backlog but cannot be made ready without acceptance criteria in ADE-007.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
