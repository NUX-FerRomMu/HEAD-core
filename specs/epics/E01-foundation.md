# E01 — Core foundations and protocol

## Outcome

A small Go executable owns durable local state and offers a frontend-neutral protocol.

## Scope and boundary

No product UI or provider assumptions.

Requirements: R01, R03, R04, R05, R10, R15, R16.

External prerequisites: none. Dependencies are ticket-level; the whole epic need not wait when a member is ready.

## Tickets

- [ADE-001 — Bootstrap the Go repository and local verification](../tickets/ADE-001-go-bootstrap.md) (S).
- [ADE-002 — Implement domain records, invariants and transition rules](../tickets/ADE-002-domain-model.md) (M).
- [ADE-003 — Persist state, events, idempotency and durable jobs](../tickets/ADE-003-sqlite-state.md) (M).
- [ADE-004 — Serve versioned local RPC with replay and stdio bridge](../tickets/ADE-004-local-rpc.md) (M).

## Acceptance and demo

Headless clients can connect and replay durable events; state and jobs are transactional. Show the ticket evidence together in one coherent demonstration. All member tickets must be accepted and any live evidence gate satisfied before the epic is done.

## Failure boundaries

Preserve the shared contracts for idempotency, human authority, exact candidate identity and restart behavior. Surface missing capabilities as blockers; do not substitute stub success.
