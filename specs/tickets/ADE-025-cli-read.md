# ADE-025 — Expose project, backlog, context and review reads in CLI

- Epic: [E07](../epics/E07-cli.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-010, ADE-013, ADE-004
- Evidence gate: none
- Requirements: R01, R03, R12, R13, R14

## Outcome

A terminal user can inspect the entire headless planning/review state through the public protocol.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-010](./ADE-010-context-projection.md) implementation and evidence.
- Accepted [ADE-013](./ADE-013-candidates-diffs.md) implementation and evidence.
- Accepted [ADE-004](./ADE-004-local-rpc.md) implementation and evidence.

## Scope and implementation boundaries

- Implement ade project open/list, config get, epic/ticket create/update/list/show/ready, attention list/next/preview/requeue and context.
- Add candidate diff/snapshot reads, --json machine output, pagination/filter flags and stable error/exit conventions.
- Use only RPC for domain actions; human CLI commands create command IDs and preserve them on retry.
- Document command-to-RPC mapping and exit codes: 0 success, 2 invalid usage, 3 domain rejection, 4 connection/internal failure.

Out of scope: No direct SQLite access, full execution/review-write CLI or shell-like command evaluation.

## Acceptance criteria

- **AC1** — A user can open a repository, create/refine an epic and ticket, make it ready and inspect queue/context without Neovim.
- **AC2** — Text mode is legible; --json emits structured results only to stdout and diagnostics to stderr.
- **AC3** — Unknown/stale IDs, invalid flags and connection loss produce documented exit codes without stack traces or silent mutations.
- **AC4** — Read-only diff/snapshot commands use immutable candidate IDs.

## Verification

- Black-box CLI tests against a temporary daemon/DB/repository.
- Golden-output tests only for stable machine/error envelopes, not incidental whitespace.
- Verify help/examples match implemented flags and mutation IDs are retry-safe.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
