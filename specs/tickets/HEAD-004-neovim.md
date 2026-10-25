# HEAD-004 — Connect Neovim

Epic: E2. Depends on: HEAD-003. Blocked.

## Outcome and scope

A thin Lua plugin displays the same tickets and attention queue as the CLI. Follow [architecture](../../Docs/architecture.md) and [agent instructions](../AGENTS.md).

Introduce `hea daemon`, a private local Unix socket, newline-delimited JSON-RPC and `hea bridge --stdio`. The daemon becomes the sole database owner; route CLI operations through it. Reuse the existing application services. The bridge reserves stdout for protocol messages and stderr for diagnostics.

Define only the methods needed for existing CLI operations; document their request, result and error shapes briefly. Bound frames to 1 MiB. After disconnection, refresh reads on reconnect; report mutation outcomes as unknown rather than automatically retrying them. Durable event replay is later work.

Add `require('hea').setup(...)`, ticket/queue views, preview, next and requeue through `jobstart`. Use explicit refresh initially. Allow buffer-local mapping overrides; install no global mappings by default.

## Acceptance

1. A minimal Neovim setup shows the same stored tickets and queue order as the CLI, with no domain logic or SQLite access in Lua.
2. Preview preserves the editing context; next/requeue use Core behavior. Plugin setup accepts the executable path and mapping overrides.
3. Split/coalesced protocol reads work. A daemon disconnect produces a visible error without freezing Neovim or silently repeating a mutation.
4. A second daemon cannot take ownership of the same database; its socket directory is private to the local user.

## Verification and handoff

Test the transport against a temporary daemon and run a small headless Neovim check for connection and rendering. Record the tested Neovim version and commands. Live agents, chat, Git editing and review UI remain later work.
