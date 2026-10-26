# Initial ticket interface

This small boundary lets storage, application logic and CLI work proceed independently. HEAD-001 declares the shared types; HEAD-002–004 implement separate sides; HEAD-005 connects them.

`internal/domain/ticket.go`:

```go
type TicketInput struct {
    Title, Description string
    AcceptanceCriteria []string
}
type Ticket struct {
    ID string
    TicketInput
}
```

`internal/application/ticket_ports.go` (imports `context` and this module's `internal/domain`):

```go
type TicketRepository interface {
    Insert(context.Context, domain.Ticket) error
    List(context.Context) ([]domain.Ticket, error)
}
type Tickets interface {
    Create(context.Context, domain.TicketInput) (domain.Ticket, error)
    List(context.Context) ([]domain.Ticket, error)
}
```

Insert is atomic; duplicate IDs fail without replacement. List preserves insertion order and an empty repository returns no tickets. The service generates opaque IDs, trims and validates titles, and preserves other input. Storage persists those records without applying workflow rules. No error is reported as a successful result.

Each owning ticket supplies its implementation entry point:

- HEAD-002: `store.Open(ctx context.Context, path string) (*Store, error)`; `*Store` implements `TicketRepository` and `Close() error`.
- HEAD-003: `application.NewTicketService(repo TicketRepository) Tickets`.
- HEAD-004: `cli.OpenTickets` is `func(context.Context, string) (application.Tickets, func() error, error)`; `cli.RunTickets(ctx context.Context, args []string, stdout, stderr io.Writer, open OpenTickets) int` returns an exit code. A successful open provides a nonnil service and cleanup; a failed open cleans up its own resources. Cleanup failure is reported on stderr with a nonzero exit code.

These are contracts for the initial CLI only; provider, daemon and frontend interfaces are not part of this change.
