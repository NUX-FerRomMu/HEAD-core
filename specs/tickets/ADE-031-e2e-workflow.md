# ADE-031 — Verify the complete headless and Neovim self-hosting flow

- Epic: [E09](../epics/E09-self-hosting.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-026, ADE-030, ADE-024
- Evidence gate: none
- Requirements: R01, R03, R04, R05, R06, R08, R09, R10, R11, R12, R13, R14, R15, R16

## Outcome

The integrated system has repeatable evidence for the full workflow and its key failure boundaries.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- [rpc contract](../contracts/rpc.md).
- [git contract](../contracts/git.md).
- [provider contract](../contracts/provider.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-026](./ADE-026-cli-workflow.md) implementation and evidence.
- Accepted [ADE-030](./ADE-030-nvim-agents-tests.md) implementation and evidence.
- Accepted [ADE-024](./ADE-024-recovery-artifacts.md) implementation and evidence.

## Scope and implementation boundaries

- Build tests/e2e/TestSelfHostingFlow with real daemon/SQLite/Git and deterministic fake providers.
- Cover two concurrent tickets, refinement, explicit assignment, validation failure/correction, human review, exact integration and restart.
- Exercise CLI and headless Neovim against the same RPC contract; consume shared fixtures.
- Document acceptance evidence and unresolved defects instead of weakening assertions to make the demo green.

Out of scope: No claim that fake-provider coverage proves live Codex readiness or actual self-hosting usability.

## Acceptance criteria

- **AC1** — The happy path and one forced-fix cycle finish at exact approved target commits with no automatic assignment/approval.
- **AC2** — A restart mid-run and after ref mutation recovers once, preserves pending queue/chat and prevents concurrent writers.
- **AC3** — An older test result, stale approval and moved target are rejected in an end-to-end path.
- **AC4** — Both CLI and Neovim can complete the required human actions; ordinary CI requires no live credentials.

## Verification

- Run go test ./tests/e2e -run TestSelfHostingFlow -count=1 and Neovim harness.
- Run the relevant full Go test/race/static checks after integration.
- Store a compact evidence matrix and reproducible failing traces for any blocker.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
