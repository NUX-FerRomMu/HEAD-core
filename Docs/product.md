# Product contract and MVP

ADE coordinates programming agents around tickets. The human directs assignments, reviews code, and decides when changes may enter the target branch. The attention queue organizes pending decisions without restricting navigation. Neovim is the first client; Core must be usable without it.

## Personas and critical journey

One developer, one local installation, several repositories, and several concurrent tickets. A developer opens a repository, creates or refines tickets, explicitly assigns a developer and tester, lets the developer produce a commit, receives validation results, reviews the exact diff, requests fixes or approves, and integrates the approved change. They can switch away and return to a durable summary with pending decisions and relevant resources. A backlog agent can propose decomposition and exit without assigning the resulting tickets.

## Normative requirements

| ID | Requirement | MVP evidence |
|---|---|---|
| R01 | Core owns business rules and persists state independently of Neovim. | Headless create-to-done journey. |
| R02 | Open projects from local Git folders; installation settings stay local and independent between clones/machines. | Path canonicalization and local settings tests. |
| R03 | Persist backlog, epics, ticket refinement, relationships, and Kanban. | Restart and decomposition tests. |
| R04 | Only the human assigns agents or approves merge. | Rejected provider-origin attempts and CLI/UI coverage. |
| R05 | Support assigned → doing → committed → fixing → committed → done with explicit failure and cancellation handling. | State transition table tests. |
| R06 | Per-ticket developer agents, finite backlog/testing runs, live agent chat; no idle developer pool. | Session lifecycle and message tests. |
| R07 | First provider is the Codex tool; provider-neutral lifecycle and an evaluated resumption policy. | Real-tool compatibility report and adapter acceptance test. |
| R08 | Isolate branches/worktrees; branch naming is configurable. | Temporary repository tests. |
| R09 | Review exact commits, side-by-side diffs, file selection, line comments, agent responses, forced correction, and manual commits. | Review integration tests and Neovim checks. |
| R10 | Prevent concurrent ADE-controlled writes by agents and humans; allow review of older snapshots. | Lease and stale-edit tests. |
| R11 | Run tests and report coverage when available; failures return work for correction. | Pass/fail/error/stale validation tests. |
| R12 | Priority queue with aging, next, previews, direct resolution, navigation, and lower-priority unread requeue. | Fake-clock ordering and UI tests. |
| R13 | Restore ticket context, summaries, chats, diff, tests, and pending actions. | Context projection and navigation tests. |
| R14 | Offer manual equivalents of automated operations and configurable commands/keymaps. | Headless human command coverage and Lua mapping override. |
| R15 | Recover safely across daemon/provider failure; neither lose durable decisions nor repeat integration. | Crash/restart scenarios. |
| R16 | MVP can develop ADE itself. | Recorded self-hosting rehearsal on an ADE ticket. |

## Scope decisions

MVP includes one local user; local Git; existing repositories with at least one commit; backlog/epics, Kanban, developer/backlog/testing roles, durable attention, local review and integration, CLI, and Neovim. Multiple projects can be open in one daemon, but `next` is project-scoped by default with an explicit all-projects option. Agents may run concurrently on different ticket worktrees. Repository integrations are serialized.

The testing role can run configured commands and propose test additions. Test additions are applied by the developer or human and cause a new candidate/validation cycle; the tester cannot silently change the candidate it certifies. Coverage is shown as unavailable when the configured tooling supplies none. Initial validation runs the project-configured commands, including coverage/static checks if configured; test selection by an LLM is not a substitute for this suite.

Explicitly deferred: hosted collaboration, multi-user authorization, cloud workspaces, GUI, remote issue synchronization, PR hosting integrations, dynamic Go plugins, marketplace, semantic code indexing, indefinite provider memory, automatic agent assignment, automatic reviewer approval, mobile clients, and universal OS support. Local Git integration does not push to a remote. Deferred work remains visible in the backlog and must not creep into MVP tickets.

## MVP release gate

A human uses Neovim to implement a real small ADE change through ticket creation, explicit developer/tester assignment, Codex development, validation, queue review, one forced correction, revalidation, approval, and local integration. Restart Core during the rehearsal, recover pending attention, and demonstrate a second concurrently executing ticket without shared worktree writes. The same core journey is reproducible with the CLI and deterministic fake providers. Record exact versions, commands, resulting commits, failures, and evidence; fake-provider success alone does not establish Codex readiness.
