# HEAD-004 — Ticket CLI commands

Dependency: HEAD-001. Scope: ticket command parsing and output against an injected service.

Implement the handler for `hea --db <path> ticket create --title <text> [--description <text>] [--criterion <text>]` and `hea --db <path> ticket list`; criterion is repeatable. Require an explicit database path for this initial interface. `cli.RunTickets` receives arguments without the executable name, output/error writers and a service-opening function, as defined in the [ticket interface](../ticket-interface.md).

Owned paths: `internal/cli/`. Tests inject a fake `application.Tickets` through the opener, so neither HEAD-002 nor HEAD-003 blocks this work. Wiring the real executable is HEAD-005; this ticket does not edit `cmd/hea` or import SQLite.

## Acceptance

- Create prints the service's returned ID; list prints the returned IDs and titles in order. The opener receives the exact database path and its cleanup runs after success or failure.
- Text containing spaces, Unicode and repeated criteria reaches the service unchanged; title validation/normalization remains in the service.
- Missing database path or invalid arguments fails before opening the service. Opener/service errors return nonzero with diagnostics on stderr and no success output.

Verification: handler tests using fake services and captured writers, plus `go test ./...` and `go vet ./...`. These test command behavior, not persistent CLI integration. JSON output, project discovery, configuration profiles, RPC and Neovim are outside this ticket.
