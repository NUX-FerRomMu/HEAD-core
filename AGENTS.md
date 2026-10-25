# Agent instructions

- Read the assigned spec and Docs/architecture.md. Work only on that ticket, after its dependencies are accepted.
- Keep documentation concise. README describes the project; Docs describes the product and architecture; agent instructions belong here; specs describe concrete work.
- Do not turn default behaviors into permanent restrictions or invent requirements. Ask about missing product decisions instead.
- Use Go for Core and Lua for Neovim. Keep provider-specific behavior in adapters.
- Keep repository documentation in English. Use HEAD for the project and `hea` for the provisional executable.
- Add only the contracts needed by the assigned change. Avoid empty scaffolding and unrelated features.
- Verify the acceptance criteria with focused tests; report exact commands, results and blockers. Do not claim implementation or live-provider evidence that does not exist.
- Use `FerRomMu <fer.rom.mu@gmail.com>` for commit author and committer, with conventional commit messages.
- Keep ticket status in specs/backlog.md. Report completed work for review; the user or designated reviewer accepts it. Push or merge only when authorized.
