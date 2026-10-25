# HEAD-001 — Go command

Dependency: none. Scope: module and executable only.

Create module `github.com/NUX-FerRomMu/HEAD-core` with entry point `cmd/hea`. Pin a supported Go toolchain. Provide `hea --help` and `hea --version`, with an explicit development version for local builds.

## Acceptance

- `go build -o /tmp/hea ./cmd/hea` succeeds from a fresh checkout.
- Help and version exit successfully without filesystem changes or starting services.
- An unknown flag returns a nonzero exit code and a readable error.

Verification: focused command tests, `go test ./...`, `go vet ./...` and gofmt. Database dependencies, CI setup and other package scaffolding are outside this ticket.
