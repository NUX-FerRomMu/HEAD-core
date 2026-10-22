# ADE-005 — Register local projects and validate installation configuration

- Epic: [E02](../epics/E02-planning.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-004
- Evidence gate: none
- Requirements: R02, R08, R14

## Outcome

A human opens a Git folder and obtains stable local project identity and validated effective settings.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [rpc contract](../contracts/rpc.md).
- [domain contract](../contracts/domain.md).
- [provider contract](../contracts/provider.md).
- [configuration contract](../contracts/configuration.md).
- [git contract](../contracts/git.md).
- Accepted [ADE-004](./ADE-004-local-rpc.md) implementation and evidence.

## Scope and implementation boundaries

- Own internal/config and project/config application services and RPC methods.
- Canonicalize roots/common Git directories; handle paths with spaces and linked worktrees; store local project IDs independently of the repository.
- Implement global plus per-project local config schema, defaults and redacted output; validation command/profile/branch/queue changes produce new config versions and digests.
- Require an explicit target branch when no unambiguous default exists; reject unsupported repository forms from the Git contract.

Out of scope: No remote cloning, relocation implementation, secrets manager or cloud configuration synchronization.

## Acceptance criteria

- **AC1** — Opening two paths into the same common Git directory returns one project; separate clones produce separate project IDs.
- **AC2** — No operational state/config is written into the repository and two independent data directories do not share settings.
- **AC3** — Invalid argv/profile/branch-template/queue settings fail atomically with field-specific errors; missing test commands remain visibly unconfigured.
- **AC4** — Config updates affect future run snapshots and invalidate approval eligibility when the validation digest changes, without rewriting old results.

## Verification

- Use temp repositories including spaces, linked worktrees, missing branches and unsupported forms.
- Test XDG defaults/overrides, redaction, malformed config and deterministic digest generation.
- Run project/config integration tests via RPC.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
