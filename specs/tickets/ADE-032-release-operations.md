# ADE-032 — Package Linux installation and measure operational readiness

- Epic: [E09](../epics/E09-self-hosting.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-031
- Evidence gate: none
- Requirements: R02, R14, R15, R16

## Outcome

A developer can install, diagnose, back up and recover ADE using documented commands and measured limits.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [rpc contract](../contracts/rpc.md).
- [git contract](../contracts/git.md).
- Accepted [ADE-031](./ADE-031-e2e-workflow.md) implementation and evidence.

## Scope and implementation boundaries

- Document/build Linux binary and Neovim runtime installation without assuming a plugin manager; record supported tool versions.
- Provide status/diagnostic export and a safe backup/restore procedure based on the implemented store mechanism.
- Measure queue/context/list latency at 10,000 items, memory and process/goroutine stability using documented hardware/fixture.
- Document config examples, provider auth references, disk usage, cleanup ownership and common recovery steps.

Out of scope: No unsupported cross-platform claims, external publication or automatic cleanup of dirty worktrees.

## Acceptance criteria

- **AC1** — A clean supported Linux setup can follow installation instructions and connect the Lua plugin to the binary.
- **AC2** — A consistent backup restores tickets/events/decisions and the documented recovery procedure does not copy a live WAL file unsafely.
- **AC3** — Diagnostic export redacts configured secrets and identifies versions/schema/unfinished jobs.
- **AC4** — Performance and 100-run resource measurements record actual numbers against targets; deviations are disclosed with release impact.

## Verification

- Run installation in a clean local test environment with no user data dependence.
- Restore a sample database and compare key entity/event counts and current ticket identity.
- Run benchmark/resource harness and record exact versions/machine/dataset in Docs/evidence/operations.md.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
