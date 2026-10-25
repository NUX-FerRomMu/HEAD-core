# Working from a spec

Read [the architecture](../Docs/architecture.md), [backlog](backlog.md), and the assigned ticket. Follow these instructions when implementing any ticket in this directory.

- Work on one assigned ticket; start only after its dependencies are accepted.
- Use Go for Core and Lua for Neovim. Keep business rules in Core and provider details in adapters.
- Implement the acceptance examples and test observable behavior. Use temporary databases, repositories and fake processes where useful; never substitute fake results for a live-tool claim.
- Document only the small contracts your change introduces. Raise a concrete conflict with the baseline instead of silently redesigning it.
- Keep all repository documentation in English. Use HEAD for the project and `hea` for the provisional executable.
- Before committing, verify author and committer are `FerRomMu <fer.rom.mu@gmail.com>`. Use conventional commit messages.
- Finish with changed files, checks/results and remaining blockers. Move the ticket to review; the human accepts completion. Push or merge only when authorized.

Implementation prompt: “Implement the linked ticket using specs/AGENTS.md. Check dependencies, satisfy its acceptance criteria, and report verification evidence.” Bootstrap will add a root AGENTS.md referring here so source-code tasks discover the same instructions.
