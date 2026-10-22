# Delivery sequence

There are 9 implementation epics and 33 bounded tickets. Use [the backlog](../specs/backlog.md) for current status and [the JSON graph](../specs/backlog.json) for dependencies. All tickets are initially unassigned. Estimates are relative: S roughly one focused implementation session; M roughly two to three; L is an integration-heavy change to split further if discovery expands it. These are planning sizes, not delivery promises.

## Milestones

| Milestone | Observable result | Required tickets |
|---|---|---|
| M0: foundation/evidence | Go skeleton and a real Codex feasibility report (or explicit blocker) | ADE-001–004, ADE-015 |
| M1: headless planner | Durable tickets, attention and context through RPC | ADE-005–010 |
| M2: controlled execution | Isolated Git, leases, generic lifecycle, backlog proposals, Codex chat | ADE-011–019 |
| M3: trusted completion | Exact validation/review/approval/integration with restart recovery | ADE-020–024 |
| M4: usable clients | Complete CLI and Neovim ticket/queue/review surfaces | ADE-025–030 |
| M5: self-hosting | E2E, measured packaging/recovery and live human rehearsal | ADE-031–033 |

Milestone rows group outcomes, not a forced serial schedule. The ticket DAG is authoritative: for example CLI reads and Neovim transport can begin once their own prerequisites land. ADE-015 can begin immediately alongside ADE-001 if the user chooses to assign multiple agents. That is a possible future work schedule, not authorization from this task to launch implementation agents now.

Recommended first assignments: ADE-001 to a Go-capable Sol agent; ADE-015 to an agent with the intended Codex tool installation. Then ADE-002 → ADE-003 → ADE-004. After that, follow ready tasks in the graph. Shared contract/dispatcher/migration files need a single integrating owner per batch to avoid competing edits even when dependency graphs allow concurrency.

No code ticket may claim completion by adding stubs for blocked integrations. Fake providers are explicit test adapters. ADE-019 and ADE-033 stay blocked if the real Codex evidence gate fails, while independent Core/client work may proceed. The backlog records this distinction.

## Repository handoff

The repository is [NUX-FerRomMu/HEAD-core](https://github.com/NUX-FerRomMu/HEAD-core), with `master` as the initial/default branch. Review the documentation PR, then assign ADE-001 after this baseline is accepted. Keep uppercase `Docs`. Use `github.com/NUX-FerRomMu/HEAD-core` as the Go module path and GitHub Actions for the initial CI workflow. ADE-001 may add source directories, root `AGENTS.md` pointing to this execution policy, module files and build tooling. This documentation baseline contains only `Docs/` and `specs/`; backlog items remain repository specifications rather than GitHub issues.
