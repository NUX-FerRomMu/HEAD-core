# ADE-023 — Integrate approved candidates and reconcile Git journals

- Epic: [E06](../epics/E06-completion.md)
- Status: blocked; unassigned
- Size: L
- Dependencies: ADE-022, ADE-014
- Evidence gate: none
- Requirements: R08, R05, R15, R04

## Outcome

Approved work reaches the target branch once, with stale-target and crash cases visible and recoverable.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-022](./ADE-022-review-approval.md) implementation and evidence.
- Accepted [ADE-014](./ADE-014-manual-editing.md) implementation and evidence.

## Scope and implementation boundaries

- Implement durable integration worker/get/retry and candidate.refresh_base following every Git-contract step.
- Use repository lease, candidate staging check, ref CAS for unchecked-out targets and guarded ff-only merge for a clean checked-out target.
- Persist intent before Git mutation, recheck exact approval/validation, verify resulting ref, then mark done and close developer.
- Target movement revokes approval and asks for a branch update/new candidate; preserve conflicts/dirty files and never force reset.

Out of scope: No squash/cherry-pick after review, automatic conflict resolution after approval or remote PR integration.

## Acceptance criteria

- **AC1** — For unchanged target/base, integration fast-forwards to exactly the approved candidate and marks the ticket done once.
- **AC2** — Moved target, dirty checked-out target, active writer or unresolved obligation blocks integration without overwriting files/refs.
- **AC3** — Crash after ref mutation but before DB completion reconciles success exactly once; an unrelated observed ref remains blocked/unknown.
- **AC4** — Refreshing the target base moves the ticket to fixing; a rebased candidate requires new validation and approval.
- **AC5** — Retry of a blocked job rechecks guards and cannot revive stale approval.

## Verification

- Real Git scenarios cover checked-out/unchecked-out target, concurrent target change, dirty worktree and divergent histories.
- Inject failures before journal, before ref mutation and after ref mutation.
- Verify no push, force-update, unintended merge commit or automatic dirty-worktree deletion occurs.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
