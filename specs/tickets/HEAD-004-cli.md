# HEAD-004 — Ticket CLI commands

Dependency: HEAD-003. Scope: expose the existing create/list operations.

Add `hea --db <path> ticket create --title <text> [--description <text>] [--criterion <text>]` and `hea --db <path> ticket list`; criterion is repeatable. Require an explicit database path for this initial interface. Commands call the application services from HEAD-003.

## Acceptance

- Create prints the new ID; a later list process shows that ID and title from the same database.
- Quoted text, Unicode and repeated criteria reach the store unchanged, except for the documented title trimming.
- Missing database path, invalid arguments or storage failures return nonzero with diagnostics on stderr, without creating a fallback database.

Verification: subprocess tests against a temporary database, plus `go test ./...` and `go vet ./...`. JSON output, project discovery, configuration profiles, RPC and Neovim are outside this ticket.
