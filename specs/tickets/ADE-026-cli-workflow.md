# ADE-026 — Expose all human execution, review and recovery commands

- Epic: [E07](../epics/E07-cli.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-025, ADE-024
- Evidence gate: none
- Requirements: R01, R04, R06, R09, R11, R14

## Outcome

A developer can finish the full ADE workflow without Neovim or an internal debug API.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [rpc contract](../contracts/rpc.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-025](./ADE-025-cli-read.md) implementation and evidence.
- Accepted [ADE-024](./ADE-024-recovery-artifacts.md) implementation and evidence.

## Scope and implementation boundaries

- Add role assignment/start/pause/stop, agent chat/intervention, backlog proposal apply/reject/complete, config update and cancellation commands.
- Add human writer/commit/publish, validation start/record, review comment/resolve/request-changes/approve and integration status/retry commands.
- Support files/stdin for long text/structured validation input without shell interpolation; preserve exact user text.
- Provide actionable blocked/stale states and explicit command IDs for retry after connection loss.

Out of scope: No domain shortcuts, hidden unsafe merge flag or external issue tracker integration.

## Acceptance criteria

- **AC1** — The fake-provider create→assign→commit→test→request-fix→retest→approve→done journey works entirely through CLI.
- **AC2** — Manual candidate publication/validation/comment resolution use the same guards as RPC automation.
- **AC3** — A provider intervention can be answered once and a stale response produces a clear rejection.
- **AC4** — Long multiline ticket/comment input and paths with spaces round-trip correctly.

## Verification

- Black-box workflow test invokes CLI commands and checks DB projections/Git target through public reads.
- Run negative approval/ownership/stale-command cases.
- Complete one manual-only ticket using writer acquire/commit/release, candidate publication and provenance-labelled validation recording, with no provider process launched.
- Audit the public method registry for missing required human command coverage.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
