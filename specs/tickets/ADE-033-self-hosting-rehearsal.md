# ADE-033 — Use ADE to develop ADE and record MVP acceptance

- Epic: [E09](../epics/E09-self-hosting.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-032, ADE-019
- Evidence gate: real_codex_authenticated
- Requirements: R16, R01, R04, R07, R09, R10, R11, R12, R15

## Outcome

The MVP is accepted only after a human completes real ADE development through its own workflow.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [git contract](../contracts/git.md).
- [provider contract](../contracts/provider.md).
- [attention contract](../contracts/attention.md).
- [rpc contract](../contracts/rpc.md).
- Accepted [ADE-032](./ADE-032-release-operations.md) implementation and evidence.
- Accepted [ADE-019](./ADE-019-codex-chat.md) implementation and evidence.

## Scope and implementation boundaries

- Choose one small real ADE change with measurable acceptance criteria and create its ticket through ADE.
- The human explicitly assigns the actual Codex/Sol profile and tester, reviews results, requests one correction and approves final integration.
- Run a second concurrent ticket and restart the daemon at a documented interruption point; inspect queue/context recovery.
- Record versions/model identifier, commands, ticket/candidate/approval IDs, resulting commit, usability findings and remaining release blockers in Docs/evidence/self-hosting.md.

Out of scope: No automatic human approval, fake-only substitute, remote publishing or sweeping redesign during rehearsal.

## Acceptance criteria

- **AC1** — A real Codex tool produces the change and correction, configured checks execute, and the human completes review in Neovim.
- **AC2** — The final target equals the approved validated candidate and the ticket is done with an inspectable audit trail.
- **AC3** — Concurrent tickets stay isolated and restart preserves pending decisions/chat without granting duplicate ownership.
- **AC4** — The human records explicit MVP acceptance or concrete blockers; absence of live credentials/tool support cannot be called accepted.

## Verification

- Perform the product release-gate journey with actual installed supported versions.
- Cross-check the resulting Git refs and persisted approval/validation identities through CLI reads.
- Retain a concise evidence report and follow-up tickets for observed defects.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
