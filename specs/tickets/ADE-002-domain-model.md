# ADE-002 — Implement domain records, invariants and transition rules

- Epic: [E01](../epics/E01-foundation.md)
- Status: blocked; unassigned
- Size: M
- Dependencies: ADE-001
- Evidence gate: none
- Requirements: R03, R04, R05, R10

## Outcome

Core has one pure, testable definition of ticket and agent state transitions.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).
- [domain contract](../contracts/domain.md).
- Accepted [ADE-001](./ADE-001-go-bootstrap.md) implementation and evidence.

## Scope and implementation boundaries

- Own internal/domain: opaque IDs, enums, entity versions, command decisions and typed errors.
- Implement the domain transition table with guards for ticket kind, dependencies, current candidate and human-only decisions.
- Represent approval/validation eligibility and writer fences as pure policies; adapters supply facts rather than importing Git/SQL.
- Create a deterministic clock/ID port where needed and preserve full opaque Git object IDs.

Out of scope: No database, provider processes, RPC handlers or frontend state machine.

## Acceptance criteria

- **AC1** — Every allowed transition in the contract yields the specified new state and domain facts; every unlisted transition rejects without mutation.
- **AC2** — Provider-origin assignment, approval and done requests fail even if their payload claims actor=human.
- **AC3** — A validation with wrong head, base or config digest cannot make a current candidate eligible for approval.
- **AC4** — Cancelled/done tickets reject reopening and dependency/parent cycles are rejected; planning completion does not require a Git candidate.

## Verification

- Table-driven tests cover each transition and guard, including assigned→doing→committed→fixing→committed→done.
- Test SHA-1 and SHA-256-shaped opaque hashes and unknown enums.
- Run go test ./internal/domain/... and relevant vet checks.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
