# ADE-029 — Review side-by-side snapshots and edit through human ownership

- Epic: [E08](../epics/E08-neovim.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-028, ADE-014, ADE-021, ADE-022, ADE-023
- Evidence gate: none
- Requirements: R09, R10, R04

## Outcome

The user reviews exact commits, comments on lines, makes controlled edits and authorizes integration from Neovim.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [rpc contract](../contracts/rpc.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-028](./ADE-028-nvim-workspace.md) implementation and evidence.
- Accepted [ADE-014](./ADE-014-manual-editing.md) implementation and evidence.
- Accepted [ADE-021](./ADE-021-review-comments.md) implementation and evidence.
- Accepted [ADE-022](./ADE-022-review-approval.md) implementation and evidence.
- Accepted [ADE-023](./ADE-023-integration.md) implementation and evidence.

## Scope and implementation boundaries

- Implement candidate/file chooser and read-only base/head diff windows with binary/large/rename handling.
- Add anchored comments, discussion, force-correction, human resolution and whole-candidate approval views.
- Expose acquire-writer/edit/commit/release/publish flow with explicit ownership and snapshot-vs-live indicators.
- Show validation identity/obligations and integration status, including target-moved recovery links.

Out of scope: No silent worktree writes from snapshot buffers, partial merges or frontend approval logic.

## Acceptance criteria

- **AC1** — Side-by-side buffers display exact base/head snapshots and can limit navigation to selected files without implying partial approval.
- **AC2** — Line comments retain their candidate/side/path and forced obligations remain visible until human resolution.
- **AC3** — Edit action waits for lease grant; while the developer writes only snapshot buffers are editable=false.
- **AC4** — Approval targets the displayed current candidate/validation, rejects stale UI state and shows pending/done/blocked integration.

## Verification

- Headless tests verify buffer flags/anchors and exact RPC payloads using deterministic snapshots.
- Integration smoke covers agent writing while old diff remains open and human handoff to a new commit.
- Test a newer candidate arriving while the approval view is open.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
