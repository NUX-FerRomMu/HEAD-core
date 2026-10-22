# ADE-015 — Verify the installed Codex structured integration surface

- Epic: [E05](../epics/E05-agents.md)
- Status: ready; unassigned
- Size: S
- Dependencies: none
- Evidence gate: none
- Requirements: R07, R06

## Outcome

The adapter implementation has real versioned protocol evidence rather than guessed flags or capabilities.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [provider contract](../contracts/provider.md).

## Scope and implementation boundaries

- Inspect the intended Codex executable, version, help and official/local protocol documentation using available access.
- Produce Docs/evidence/codex-compatibility.md with exact launch method, supported version range, auth prerequisites, capability matrix and redacted protocol samples.
- Probe start, event streaming, chat delivery, intervention/approval response, stop, process exit and resume where supported.
- Define gate codex_interface_verified as true only with a supported structured interface and reproducible samples; if unavailable, produce an explicit blocker and a supported-version/tooling recommendation backed by evidence.

Out of scope: No production adapter, API-only replacement, terminal-screen scraping or invented Sol model identifier.

## Acceptance criteria

- **AC1** — The report distinguishes observed, documented-but-unverified and unavailable capabilities; no unsupported command flag is proposed as fact.
- **AC2** — At least one structured start-to-terminal exchange and cancellation is captured when the tool is available; otherwise the relevant probe remains blocked.
- **AC3** — The report specifies how chat works during an active turn, or explicitly states queued/new-turn behavior.
- **AC4** — No credentials are committed and the report identifies the evidence required to unblock ADE-019.

## Verification

- Record exact probe commands, tool version, exit results and redacted fixtures.
- Check sample events against the proposed provider port and name translation gaps.
- Do not run paid/provider-mutating probes without the existing task/environment authorization; report unavailable authentication honestly.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
