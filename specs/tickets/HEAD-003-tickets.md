# HEAD-003 — Create and list tickets

Dependency: HEAD-001. Scope: two application operations against a repository interface.

Implement `NewTicketService` and `application.Tickets` from the [ticket interface](../ticket-interface.md). Create generates an ID and accepts a nonempty trimmed title; description and criteria may be empty while drafting. List returns repository records in insertion order. All storage is through the injected `TicketRepository`.

Owned paths: `internal/application/ticket_service.go` and its tests. An in-memory repository test double lets this ticket finish independently of HEAD-002; no SQL or migrations belong here.

## Acceptance

- Creating two tickets gives distinct IDs and passes the normalized records to the repository. Listing returns its original content and order.
- A blank title fails without calling Insert. Repository failures from either operation propagate as errors; failed creation returns no successful ticket.
- Application/domain code does not import the SQLite driver or CLI package.

Verification: service tests with an in-memory repository and injected storage errors; `go test ./...` and `go vet ./...`. Database persistence is verified by HEAD-002 and end-to-end wiring by HEAD-005. Editing, lifecycle states, project relationships and agent assignment are outside this ticket.
