# Architecture

## Decisions to build on

| Area | Choice and reason |
|---|---|
| Core | Go modular monolith. Go is simple to read and maintain, familiar to the maintainer, and provides lightweight concurrency as a first-class feature for coordinating processes and events. |
| Frontend | Lua Neovim plugin. Neovim handles editing and navigation; Go owns business rules and durable state. |
| Storage | SQLite via `database/sql` and a pure-Go driver. One local daemon owns writes; no separate database service. Verify and pin driver/toolchain compatibility during bootstrap. |
| Communication | Local JSON-RPC over a Unix socket initially; `hea bridge --stdio` connects Neovim. CLI and plugin use the same application services. Keep the transport replaceable for other platforms. |
| Providers | Adapters for external tools or APIs, behind a small Go interface. Tool-specific launch, messaging and resumption behavior belongs in adapters, never in the domain. Select and verify the first adapter when agent work begins. |
| Git | Separate development worktrees per ticket, immutable review snapshots, and explicit writer ownership. Design the integration details when implementing that workflow. |

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

Adapters depend on application interfaces; domain imports neither adapters nor UI code. Use ordinary Go packages, not a plugin framework. Configuration and runtime state stay outside the repository in the user's local directories.

## Invariants for later features

- Assignment and merge approval belong to the human. Provider output cannot grant either authority.
- Only one HEAD-controlled writer may own a ticket worktree. Stop and verify the previous process before handing it over; an expired timer alone is insufficient.
- Publish a candidate only from a clean, stable worktree while preventing another writer from acquiring it. Tests and approval bind to its exact commit, base and validation configuration.
- Validation needs at least one configured required check. Empty or missing checks cannot produce a pass. Changed candidates or target bases require fresh validation and approval.
- An agent reply does not resolve a forced correction; the human decides. Resolving a queue item cannot bypass the underlying action.
- Preserve dirty files and uncertain operations on failure. Reconcile interrupted work before restarting writes; never infer successful integration from incomplete evidence.

Specify wire methods, schemas and recovery algorithms in the ticket that introduces them. Do not build those later mechanisms during the starter work.
