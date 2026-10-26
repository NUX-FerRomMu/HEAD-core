# HEAD-005 — Wire the persistent CLI

Dependencies: HEAD-002, HEAD-003, HEAD-004. Scope: connect the tested components in `cmd/hea`.

Supply the CLI opener using `store.Open`, `application.NewTicketService` and the store's cleanup. Route ticket commands through `cli.RunTickets`, preserving HEAD-001's help/version behavior. No new ticket behavior or parallel implementation belongs here.

## Acceptance

- Build `hea`, create a ticket through the real executable, then list it from a separate process using the same temporary database; its ID and title persist.
- A rejected blank title leaves the database unchanged. Invalid arguments and an inaccessible database return nonzero with diagnostics; help/version still works without opening storage.
- The executable uses the real SQLite repository and service. No test double is reachable through production command wiring.

Verification: subprocess tests in `cmd/hea` against temporary databases, plus `go test ./...`, `go vet ./...` and `go build ./cmd/hea`. This ticket begins after all three parallel implementations are accepted.
