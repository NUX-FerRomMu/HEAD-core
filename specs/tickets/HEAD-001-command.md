# HEAD-001 — Go command

Dependency: none. Scope: executable and the local runtime shared by independent features.

Create module `github.com/NUX-FerRomMu/HEAD-core` with entry point `cmd/hea`. Pin a supported Go toolchain. Add command registration and project-scoped SQLite access through a pure-Go driver. Provide `hea --help` and `hea --version`, with an explicit development version for local builds.

## Acceptance

- `go build -o /tmp/hea ./cmd/hea` succeeds from a fresh checkout.
- Help and version exit successfully without filesystem changes or starting services.
- An unknown flag returns a nonzero exit code and a readable error.
- `--project` resolves the project root (including Git worktrees); `--data-dir` overrides the state directory. State defaults to `$XDG_STATE_HOME/hea` or `~/.local/state/hea`, outside the checkout. Different projects have separate databases.
- Storage opens lazily with foreign keys and a busy timeout; features own their schemas.

Verification: focused command tests, `go test ./...`, `go vet ./...` and gofmt. Domain features and a daemon are outside this ticket.
