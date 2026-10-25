# Architecture

| Component | Decision and reason |
|---|---|
| Core | Go modular monolith. Go is readable, familiar to the maintainer, and has lightweight concurrency suited to coordinating agents and processes. |
| Neovim | Lua plugin for presentation and editor integration. Business rules and durable state live in Core. |
| Storage | Local SQLite through `database/sql` and a pure-Go driver; no separate database server. |
| Clients | A local daemon exposes Core to CLI and Neovim. JSON-RPC over a local socket, with a stdio bridge, is the initial transport choice. |
| Providers | External tools or APIs behind Go interfaces. Domain logic is independent of any provider's protocol or session behavior. |
| Git | Separate ticket worktrees and commit-based review isolate concurrent changes. |

Proposed layout:

```text
cmd/hea/                   executable and CLI
internal/domain/           rules and types; no I/O
internal/application/      use cases and interfaces
internal/store/            SQLite and migrations
internal/transport/        local RPC and stdio bridge
internal/adapter/          Git, processes and providers
lua/hea/, plugin/hea.lua    Neovim client
```

Core contains domain rules, application services and adapters for storage, Git and providers. Adapters implement application interfaces; domain code has no UI or provider dependency. The entry point is `cmd/hea`; the Go module is `github.com/NUX-FerRomMu/HEAD-core`.

Configuration and runtime state live outside the project checkout. Platform-specific transport and process code remain behind adapters.

## Policy and coordination

Manual assignment and review approval are defaults. Automation policy is separate from domain state so those defaults can evolve without replacing the core.

Ticket-scoped agents are the usual execution model, not a restriction on shared or project-wide agents. Provider capabilities determine how each adapter supports conversations and sessions.

Writer ownership coordinates HEAD-controlled changes to a worktree. Validation and review refer to specific commits so results remain attributable when work continues. Interrupted operations retain enough state for recovery.
