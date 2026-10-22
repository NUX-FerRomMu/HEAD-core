# Architecture decision records

All decisions below are **accepted design choices for the first implementation** on 2026-10-22. They can be superseded by a documented decision. Provider capabilities remain evidence-gated rather than assumed.

| ADR | Decision |
|---|---|
| [001](001-go-core-lua-client.md) | Go modular monolith and Lua Neovim client |
| [002](002-local-state-and-protocol.md) | SQLite, durable jobs, local versioned JSON-RPC |
| [003](003-worktrees-and-review.md) | Isolated worktrees, fenced ownership, exact-commit review |
| [004](004-provider-boundary.md) | Codex tool adapter and capability-gated sessions |
| [005](005-attention-policy.md) | Deterministic attention with uncapped aging |
