# ADE-030 — Add agent chat, interventions, testing and backlog proposal views

- Epic: [E08](../epics/E08-neovim.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-029, ADE-019, ADE-020, ADE-018
- Evidence gate: none
- Requirements: R06, R07, R11, R12, R13, R03

## Outcome

All daily agent/test/refinement decisions are reachable from the ticket workspace.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [provider contract](../contracts/provider.md).
- [rpc contract](../contracts/rpc.md).
- [attention contract](../contracts/attention.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-029](./ADE-029-nvim-review.md) implementation and evidence.
- Accepted [ADE-019](./ADE-019-codex-chat.md) implementation and evidence.
- Accepted [ADE-020](./ADE-020-validation.md) implementation and evidence.
- Accepted [ADE-018](./ADE-018-backlog-agent.md) implementation and evidence.

## Scope and implementation boundaries

- Implement per-agent transcript/input with live/queued/failed delivery and explicit closed/paused/running status.
- Render typed provider interventions and acknowledge response only after Core success.
- Show validation checks, coverage-unavailable state, bounded logs/artifacts and suggested test patches as reviewable resources.
- Add backlog proposal diff/accept/reject and planning completion; link everything through context and attention.

Out of scope: No generic detached global chat, provider terminal scraping or automatic test-patch application.

## Acceptance criteria

- **AC1** — Selecting a ticket agent opens only that agent transcript and preserves queued-message status across reconnect.
- **AC2** — Closed agents reject live sends visibly; unsupported live-message capability is described accurately.
- **AC3** — A failed/stale validation is clearly distinguished from the current candidate and coverage is never fabricated.
- **AC4** — Accepting a proposal shows resulting backlog children without automatic assignments; failed/stale proposals remain unapplied.
- **AC5** — Intervention response, validation retry and review navigation can be reached from queue actions.

## Verification

- Headless component tests for all delivery/run/validation/proposal states.
- Integration smoke with fake providers and one opt-in real Codex chat/intervention path.
- Check bounded log scrolling and malformed/unknown action fallbacks.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
