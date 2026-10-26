# HEAD-008 — Use the three features from Neovim

Dependencies: HEAD-005, HEAD-006, HEAD-007. Integrate the reviewed feature branches without rewriting their implementations.

Add a Lua client for Neovim 0.10+ using asynchronous `hea` subprocesses. `:HEADBoard` displays the four Kanban columns; mappings create, inspect, edit and move tickets. `:HEADNotifications` displays the ordered queue and supports adding, dismissing and requeuing. `:HEADRun` accepts a prompt and streams Codex output; `:HEADStop` cancels the active run. Completion/failure creates a persistent notification. Configuration accepts the `hea` executable, project and data directory.

## Acceptance

- The board and queue reflect persisted Go state after reopening Neovim. Errors are visible and leave buffers usable.
- Agent output remains responsive; completion and failure appear in the queue. Stop and editor exit clean up the child process.
- No business rules or persistence are duplicated in Lua. One active run per client is the initial UI behavior, not a core restriction.
- Headless Neovim tests exercise real CLI state and a controlled provider; document build, setup and commands in a short getting-started guide.

This first client invokes the CLI directly; the planned daemon/RPC transport is deferred. No automatic Git merge or ticket transitions.
