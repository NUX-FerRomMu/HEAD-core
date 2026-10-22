# ADE-021 — Anchor review discussions and enforce requested corrections

- Epic: [E06](../epics/E06-completion.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-013, ADE-017
- Evidence gate: none
- Requirements: R09, R04, R13

## Outcome

The human can discuss exact lines and force a correction without losing obligations across revisions.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [git contract](../contracts/git.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-013](./ADE-013-candidates-diffs.md) implementation and evidence.
- Accepted [ADE-017](./ADE-017-agent-lifecycle.md) implementation and evidence.

## Scope and implementation boundaries

- Implement review.comment/comments/respond/resolve/request_changes and comment projections in context.
- Validate immutable candidate path/side/line anchors; support file comments for binary/large diffs.
- Deliver comment/forced-obligation context through the generic agent handoff/event port; replies do not resolve human obligations.
- On a new candidate, retain old anchors as historical/obsolete and carry unresolved forced obligations forward.

Out of scope: No automatic comment-line rebasing, agent self-resolution or partial-file integration.

## Acceptance criteria

- **AC1** — A line comment opens the original side/path/line even after a new candidate is published.
- **AC2** — Agent justification records a response but leaves an open forced correction unresolved until the human resolves it.
- **AC3** — Requesting changes on the current candidate revokes approval and moves committed→fixing; stale candidate requests reject.
- **AC4** — Selecting only some files for review never narrows eventual approval/merge scope.

## Verification

- Test rename/deletion/binary/invalid-line anchors against real snapshot fixtures.
- Exercise forced request→new candidate→unresolved obligation→human resolution.
- Verify context lists current obligations and links old comments correctly.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
