# Agent instructions

- Read the assigned spec and Docs/architecture.md. Create or update the backlog ticket before branching. Start from `develop` unless the ticket needs an unmerged implementation; in that case branch from it and reference its PR as a prerequisite.
- Keep documentation concise. README describes the project; Docs describes the product and architecture; agent instructions belong here; specs describe concrete work.
- Do not turn default behaviors into permanent restrictions or invent requirements. Ask about missing product decisions instead.
- Keep repository documentation in English. Use HEAD for the project and `hea` for the provisional executable.
- Add only the contracts needed by the assigned change. Avoid empty scaffolding and unrelated features.
- Verify the acceptance criteria with focused tests; report exact commands, results and blockers. Do not claim implementation or live-provider evidence that does not exist.
- Use `FerRomMu <fer.rom.mu@gmail.com>` for commit author and committer, with conventional commit messages.
- Keep ticket status in specs/backlog.md. For authorized implementation work, test, commit, push with normal `git push` and open a PR listing prerequisites. Do not request privilege escalation for pushes. The user or designated reviewer accepts the work; merge only when authorized.
