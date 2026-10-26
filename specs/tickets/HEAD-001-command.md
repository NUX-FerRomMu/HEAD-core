# HEAD-001 — Go command

Dependency: none. Scope: module, executable and the shared interfaces needed for parallel work.

Create module `github.com/NUX-FerRomMu/HEAD-core` with entry point `cmd/hea`. Pin a supported Go toolchain. Provide `hea --help` and `hea --version`, with an explicit development version for local builds.

Declare the records in `internal/domain/ticket.go` and interfaces in `internal/application/ticket_ports.go` from the [ticket interface](../ticket-interface.md). These are the only shared code prerequisites for HEAD-002, HEAD-003 and HEAD-004; no storage, service implementation or ticket commands yet.

## Acceptance

- `go build -o /tmp/hea ./cmd/hea` succeeds from a fresh checkout.
- Help and version exit successfully without filesystem changes or starting services.
- An unknown flag returns a nonzero exit code and a readable error.
- The ticket records and interfaces compile without a database driver, CLI package or placeholder implementations.

Verification: focused command tests, `go test ./...`, `go vet ./...` and gofmt. Database dependencies, CI setup and unrelated package scaffolding are outside this ticket.
