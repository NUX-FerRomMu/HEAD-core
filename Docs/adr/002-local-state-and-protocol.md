# ADR-002: SQLite and a local protocol

Status: accepted.

Use a per-installation SQLite database owned by one daemon, atomic state/event writes, durable jobs, and a local Unix socket on Linux. Use newline-delimited JSON-RPC 2.0 over sockets and a stdio bridge for Neovim. The pure-Go SQLite driver avoids a CGO compiler requirement. Schema migrations are numbered SQL files; do not use event sourcing.

Alternatives: repository YAML as mutable runtime state makes concurrent writes and local clone independence harder; PostgreSQL adds a service; HTTP adds unnecessary exposure; direct Lua access to SQLite defeats ownership. NDJSON makes the bridge and test fixtures small but requires bounded frames, escaping, backpressure, and explicit replay.

Consequences: installation-local state is not automatically shared or synced; design backups/recovery; retain events for reconnect; platform-specific local transport is required for Windows. RPC is local authenticated-by-OS-user coordination, not a public remote API.
