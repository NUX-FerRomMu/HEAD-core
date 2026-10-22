# ADE-020 — Run exact-candidate validation and publish test evidence

- Epic: [E06](../epics/E06-completion.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-013, ADE-017, ADE-016
- Evidence gate: none
- Requirements: R11, R05, R04

## Outcome

A committed candidate receives reproducible test results which cannot accidentally certify another commit.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [git contract](../contracts/git.md).
- [provider contract](../contracts/provider.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- [configuration contract](../contracts/configuration.md).
- Accepted [ADE-013](./ADE-013-candidates-diffs.md) implementation and evidence.
- Accepted [ADE-017](./ADE-017-agent-lifecycle.md) implementation and evidence.
- Accepted [ADE-016](./ADE-016-provider-port.md) implementation and evidence.

## Scope and implementation boundaries

- Implement validation.start/get/list/record, tester job scheduling on candidate publication/late tester assignment, and isolated validation worktrees.
- Run every required configured argv check with timeout/output artifacts, exit-code status, exact candidate/base/config digest and optional coverage.
- Allow assigned tester findings/test-patch artifacts, but keep runner results authoritative and route patches through developer/human ownership.
- Handle human-entered results with explicit provenance and identical eligibility/completeness rules. Missing commands produce configuration error.

Out of scope: No test-generation patch auto-application, missing-check pass, or claim of universal coverage format support.

## Acceptance criteria

- **AC1** — A current candidate with assigned tester runs all required checks on its exact clean snapshot; all pass yields passed validation and review-ready attention.
- **AC2** — A failed current check moves committed to fixing and schedules the assigned nonpaused developer; infrastructure error leaves committed with retry attention.
- **AC3** — A delayed result for an older candidate/base/config is retained as history but cannot move current state or permit approval.
- **AC4** — Tracked/untracked source mutations during tests cause error; suggested test patches cannot change the validated head.
- **AC5** — Without tester assignment no provider validation launches; role-assignment attention appears and human manual recording remains explicit.

## Verification

- Use executable check fixtures for pass/fail/timeout/signal/malformed coverage and source mutation.
- Run old/new candidates concurrently in separate worktrees and verify stale result handling.
- Test duplicate validation starts, missing checks and forged incomplete human records.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
