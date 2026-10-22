# Implementation specifications

Start with [backlog.md](backlog.md), select one `ready` ticket whose dependencies are done, then use [AGENTS.md](AGENTS.md). [backlog.json](backlog.json) is the machine-readable schedule; individual [tickets](tickets/) contain the acceptance contract. [Epics](epics/) describe outcomes and boundaries. [Shared contracts](contracts/) are normative; [fixtures](fixtures/) contain cross-component examples.

Statuses: `ready` = available to assign, `blocked` = prerequisite/gate unsatisfied, `in_progress`, `review`, `done`. All work starts unassigned. A dependency is satisfied only when its implementation and required evidence are accepted, not when an agent says it has started. Recompute readiness after each accepted ticket; a gate can keep an otherwise-ready ticket blocked. Update the table, JSON and ticket status together. This planning backlog is distinct from ADE's runtime ticket state machine.

Every implementation ticket includes outcome, dependencies, required reading, boundaries, acceptance scenarios, validation, and handoff. Do not give a Sol agent only the title. Give the exact ticket path and the execution prompt below.

```text
Implement specs/tickets/ADE-XXX-<slug>.md in this repository.
Read specs/AGENTS.md, the ticket, its linked contracts and listed dependencies.
Confirm prerequisites are accepted; if blocked, report the exact blocker.
Implement only the ticket's scope and verify its acceptance examples.
Preserve Go Core/Lua Neovim boundaries and human assignment/approval rules.
Do not invent Codex capabilities or silently change a shared contract.
Finish with changed files, acceptance-to-test evidence, commands/results,
known limits, and any exact follow-up needed. Do not mark yourself accepted.
```

Use the installed provider's actual supported Sol identifier when launching the implementation agent; these specs do not require a model API or hard-code a model alias.
