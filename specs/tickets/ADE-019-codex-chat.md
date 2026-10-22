# ADE-019 — Integrate verified Codex sessions, chat and interventions

- Epic: [E05](../epics/E05-agents.md)
- Status: blocked; unassigned
- Size: L
- Dependencies: ADE-015, ADE-017
- Evidence gate: codex_interface_verified
- Requirements: R06, R07, R12, R15

## Outcome

A real Codex tool runs through the generic lifecycle and the human can communicate with the exact ticket agent.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [provider contract](../contracts/provider.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- Accepted [ADE-015](./ADE-015-codex-compatibility.md) implementation and evidence.
- Accepted [ADE-017](./ADE-017-agent-lifecycle.md) implementation and evidence.

## Scope and implementation boundaries

- Implement only the structured tool interface/capability behavior evidenced by ADE-015, with version probing and clear unsupported errors.
- Implement durable agent.send/messages and intervention.respond, pending/sent/acknowledged/failed message delivery and per-run intervention correlation.
- Deliver supported live messages or visibly queue them for a new/resumed turn; build bounded durable handoff context and optional verified resume.
- Map provider errors/approvals to deduplicated attention; redact configured secrets, bound logs and preserve protocol fixtures for adapter tests.

Out of scope: No proprietary API-only integration, TUI scraping, unsupported live messaging or fabricated cost reporting.

## Acceptance criteria

- **AC1** — A supported installed Codex can complete a real developer turn through the port; unsupported versions fail before ticket writes.
- **AC2** — Chat is tied to one agent, persists across reconnect, and never appears delivered until acknowledged; closed agents reject new sends.
- **AC3** — An intervention response reaches only its matching live run/external ID; stale/duplicate responses cannot resume the wrong process.
- **AC4** — Missing/expired resume references visibly fall back to fresh context with mandatory spec/obligations intact.
- **AC5** — Process failure/timeout/cancellation preserves diagnostics and follows the lease uncertainty rule.

## Verification

- Replay redacted versioned protocol fixtures offline for normal/error/intervention/message cases.
- Run opt-in live smoke with exact version/model/profile and record evidence in Docs/evidence/codex-adapter.md.
- Test mandatory-input budget overflow and secret redaction; fake success cannot satisfy the live smoke criterion.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
