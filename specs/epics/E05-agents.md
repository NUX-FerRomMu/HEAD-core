# E05 — Provider-neutral agents and Codex integration

## Outcome

Humans assign finite ticket-scoped agents, refine backlog and communicate through a verified tool interface.

## Scope and boundary

No autonomous assignments, API-only lock-in, invented capabilities or permanent developer pool.

Requirements: R03, R04, R05, R06, R07, R12, R15.

External prerequisites: ADE-002, ADE-003, ADE-004, ADE-006, ADE-007, ADE-012. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-015 — Verify the installed Codex structured integration surface](../tickets/ADE-015-codex-compatibility.md) (S).
- [ADE-016 — Implement provider contracts and a deterministic fake tool](../tickets/ADE-016-provider-port.md) (M).
- [ADE-017 — Assign roles and orchestrate ticket-scoped agents](../tickets/ADE-017-agent-lifecycle.md) (M).
- [ADE-018 — Review and atomically apply backlog-agent proposals](../tickets/ADE-018-backlog-agent.md) (M).
- [ADE-019 — Integrate verified Codex sessions, chat and interventions](../tickets/ADE-019-codex-chat.md) (L).

## Acceptance and demo

Fake contract suite plus real Codex evidence; proposals are accepted by humans and chat delivery is explicit. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
