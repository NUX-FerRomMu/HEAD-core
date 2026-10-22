# ADE-013 — Publish immutable candidates and exact snapshot diffs

- Epic: [E04](../epics/E04-git.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-012, ADE-007
- Evidence gate: none
- Requirements: R09, R05, R08

## Outcome

A clean development commit becomes a reviewable immutable candidate with exact base/head identity.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-012](./ADE-012-writer-leases.md) implementation and evidence.
- Accepted [ADE-007](./ADE-007-workflow-commands.md) implementation and evidence.

## Scope and implementation boundaries

- Implement candidate.publish through the application service and Git verification; bind lease fence/clean state/ancestry.
- Implement diff.get and snapshot.read using Git object data, structured hunks and rename/binary metadata.
- Invalidate prior approval eligibility, preserve historical candidates via owned refs and emit candidate.published.
- Schedule validation only through a registered hook when a tester assignment exists; otherwise create assignment-needed attention once that integration is wired.

Out of scope: No comment discussion, approval, live editor, or automatic validation pass.

## Acceptance criteria

- **AC1** — Publishing a clean new descendant commit records exact head/base and moves doing/fixing to committed.
- **AC2** — Dirty, unchanged, detached/wrong-branch, stale-fence or non-descendant publication fails without a candidate.
- **AC3** — Diff and snapshot content stays bound to the original candidate while the developer later changes the branch.
- **AC4** — Rename, deletion, binary, traversal and oversized blob cases return defined metadata/errors without arbitrary filesystem reads.

## Verification

- Build real commit graphs and modify the live worktree after snapshot creation.
- Test line/hunk content and rename metadata against known fixtures.
- Crash/retry publication command and verify one candidate/event.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
