# ADE-004 — Serve versioned local RPC with replay and stdio bridge

- Epic: [E01](../epics/E01-foundation.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-003
- Evidence gate: none
- Requirements: R01, R15

## Outcome

Any local client can connect, issue bounded commands, and recover domain notifications after disconnection.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [rpc contract](../contracts/rpc.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-003](./ADE-003-sqlite-state.md) implementation and evidence.

## Scope and implementation boundaries

- Own internal/transport/rpc and local plus daemon/bridge wiring; establish the single-daemon OS lock and Linux private socket.
- Implement hello/status, request validation, error mapping, command envelope middleware, method registration and durable event subscriptions.
- Implement frame limits, independent writer queues, bounded slow-client handling and consistent snapshot/replay behavior.
- Implement ade bridge --stdio with JSON only on stdout and safe daemon connection/autostart; no recursive bridge startup.

Out of scope: No feature-specific ticket handlers beyond dispatcher hooks; no public HTTP API or Windows transport.

## Acceptance criteria

- **AC1** — A compatible hello works; an incompatible protocol, invalid JSON, batch request or oversized frame returns the contract error or closes safely.
- **AC2** — Bridge output remains parseable when daemon diagnostics occur; split/coalesced socket reads preserve frames.
- **AC3** — Disconnect/reconnect from the last cursor replays durable events without a gap; duplicate events can be detected by cursor.
- **AC4** — A slow consumer is disconnected without blocking a concurrent mutation, and two daemon launches leave one owner.
- **AC5** — Unsafe socket ownership/permissions/symlinks are rejected; no TCP listener is opened.

## Verification

- Use real temporary Unix sockets and subprocess bridge tests, fragmented frames and Unicode/newlines.
- Inject a mutation during replay handoff and test CURSOR_EXPIRED client behavior with a simulated retention floor.
- Run transport tests and go test -race on concurrency cases.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
