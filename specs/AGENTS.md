# Instructions for implementation agents

This file governs changes beneath `specs/`; implementation prompts also explicitly adopt it as the repository-wide execution procedure. ADE-001 creates a root instruction file referencing it, preserving any existing repository instructions.

1. Read the assigned ticket in full, `Docs/README.md`, `Docs/architecture.md`, and every contract linked by the ticket. Check the accepted dependency implementations and backlog status. Do not start dependent work when a required artifact/capability is absent.
2. Work on exactly the assigned scope. Go owns Core/CLI; Lua owns Neovim. Do not add TypeScript, change the database/protocol, introduce runtime Go plugins, or move business logic into Lua.
3. Derive observable tests from the acceptance criteria before implementing behavior. Use deterministic fake providers, clocks, temporary Git repositories and databases where appropriate. Research-only tickets produce evidence instead of artificial tests.
4. Shared contracts outrank ticket shorthand. If implementation exposes an ambiguity, document a narrow proposed contract amendment with rationale and affected tickets. Routine internal choices are yours; user assignment/approval, commit identity, writer exclusion and provider capability gates cannot be weakened.
5. Read the current repository before choosing file edits. Proposed paths are ownership guidance, not an excuse to duplicate existing modules. Avoid drive-by refactors. Multiple agents must coordinate shared migration/dispatcher/schema changes through the assigned integration owner; do not overwrite another agent's work.
6. Run only relevant existing checks plus required repository checks; report exact commands and outcomes. Do not create hollow tests that mirror implementation or claim a live provider/OS/performance check passed using only a fake.
7. Handle errors and interrupted operations explicitly. Do not discard dirty worktrees, overwrite target refs, erase user data, auto-assign agents, or turn missing validation into green status. Never log credentials or invent tool flags.
8. Keep docs and contract examples consistent with the implementation. A contract change includes regression cases and an impact note for dependent tickets. Do not rewrite historical evidence.
9. Finish with: scope completed; changed paths; acceptance criterion → evidence mapping; commands/results; remaining limitations; migration/recovery impact; suggested next ready tickets. Set status to `review`, not `done`; acceptance/assignment belongs to the human or the user's designated reviewer.
10. Do not push, publish, send external messages, or create external tracker issues unless the user separately authorizes it. Local implementation and tests are the expected task. This planning package does not authorize starting a pool of agents.

## Blocker format

State the missing dependency or observed contradiction, show a minimal reproduction/evidence, name affected acceptance criteria, and propose the smallest next action. Keep useful independent work within the assigned scope moving where possible; do not hide a blocking gate behind a TODO or success claim.
