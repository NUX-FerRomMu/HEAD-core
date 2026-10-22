# ADE-001 — Bootstrap the Go repository and local verification

- Epic: [E01](../epics/E01-foundation.md)
- Status: ready; unassigned
- Size: S
- Dependencies: none
- Evidence gate: none
- Requirements: R01, R16

## Outcome

A contributor can build a minimal Go ADE binary and run deterministic local checks from a fresh repository.

## Required reading

- [Execution policy](../AGENTS.md), [architecture](../../Docs/architecture.md), and [testing strategy](../../Docs/testing.md).

## Scope and implementation boundaries

- Create go.mod with module path `github.com/NUX-FerRomMu/HEAD-core`, cmd/ade, inward-facing internal packages and a minimal version command; use Go 1.26 and pin a tested patch/toolchain plus dependency versions.
- Verify modernc.org/sqlite compatibility with the pinned Go version, record the minimum tested Git and Neovim versions, and keep CGO unnecessary for the application build.
- Create root AGENTS.md referring to specs/AGENTS.md without overwriting existing instructions; add an offline-by-default test entry point and a GitHub Actions CI workflow for this repository.
- Document build/run/check commands and ignore runtime state, generated binaries and credential files. Preserve Docs and specs.

Out of scope: No daemon behavior, database schema, feature stubs, live provider calls or Neovim UI.

## Acceptance criteria

- **AC1** — From a clean checkout with documented tool prerequisites, go build ./cmd/ade and go test ./... succeed.
- **AC2** — The binary reports its build version and help without launching a provider or touching a user repository.
- **AC3** — Core contains no JavaScript/TypeScript runtime and the initial package graph permits application-to-domain dependencies without adapter imports in domain.
- **AC4** — Pinned dependency/toolchain choices and actual check results are recorded; missing external tools are identified accurately.

## Verification

- Run gofmt verification, go vet ./..., go test ./..., and go build ./cmd/ade.
- Demonstrate the documented build with CGO_ENABLED=0; separately document race-detector build requirements.
- Check that the repository does not track runtime databases, logs or credentials.

## Handoff

Provide changed files, AC-to-test/evidence mapping, exact commands/results, and known limits. Record any schema/contract change and its downstream impact. Update this ticket and both backlog views to `review`; completion is accepted by the human/designated reviewer. Do not mark an unavailable live check passed.
