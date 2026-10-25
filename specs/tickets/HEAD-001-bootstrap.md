# HEAD-001 — Bootstrap

Epic: E1. Depends on: none. Ready, unassigned.

## Outcome and scope

Create a buildable Go project with module `github.com/NUX-FerRomMu/HEAD-core` and executable `hea`. Follow [architecture](../../Docs/architecture.md) and [agent instructions](../AGENTS.md).

Add `cmd/hea`, minimal version/help commands, root AGENTS.md pointing to specs/AGENTS.md, and GitHub Actions for build, formatting, vet and tests. Pin a supported Go toolchain and a compatible pure-Go SQLite driver. Record the tested versions in a short development note; defer empty scaffolding and feature implementations.

## Acceptance

1. A fresh checkout builds with `go build ./cmd/hea`; help/version exits without starting a provider or changing a repository.
2. `go test ./...` and `go vet ./...` pass, and tracked Go files are formatted. Keep CI and local commands consistent.
3. A minimal in-memory SQLite smoke check works with the selected driver and `CGO_ENABLED=0`.

## Verification and handoff

Run the build, vet, formatting and tests. Report pinned versions and exact commands/results. No daemon, UI, agent execution or ticket behavior belongs in this change.
