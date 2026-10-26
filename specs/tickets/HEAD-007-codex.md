# HEAD-007 — Codex task execution

Dependency: HEAD-001. Independent of tickets and notifications.

`hea agent status` reports Codex availability and login readiness. `hea agent run --prompt TEXT [--model NAME]` invokes the installed Codex CLI in the selected project using its existing authentication and configuration. Stream JSONL events to stdout and diagnostics to stderr. An adapter keeps process details out of the command's orchestration. Do not add privilege or approval bypass flags.

## Acceptance

- A task reaches `codex exec --json` with its prompt through stdin and the selected working directory. Optional model is passed unchanged; shell interpolation is never used.
- Output streams before completion. Nonzero provider exits and missing executables are failures; successful completion is distinguishable.
- Ctrl-C/termination cancels the child and descendants and waits for cleanup. No orphan process on cancellation.
- Tests use a controlled executable for arguments, streaming, failure and cancellation. Report a live authenticated smoke test separately when available.

No persistent conversations, autonomous assignment or parallel writers in this ticket.
