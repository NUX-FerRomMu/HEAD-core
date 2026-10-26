# HEAD-005 — Persistent backlog and Kanban

Dependency: HEAD-001. Deliver a complete feature through the CLI.

`hea ticket create --title TEXT [--description TEXT] [--criterion TEXT ...]` creates a todo ticket. Provide `get ID`, `list`, `edit ID` with the same fields, `move ID STATE`, and `board`. States: todo, in_progress, review, done. Manual moves between any states are allowed. JSON output is the client contract.

## Acceptance

- Title is trimmed and required; descriptions and repeated criteria preserve Unicode and newlines. IDs are distinct and stable.
- Create, edit and move survive restarting the command. Missing IDs, invalid states and unknown arguments fail without partial writes.
- Board groups all tickets by state; list preserves creation order. Data is isolated by project.
- Temporary SQLite tests cover persistence and validation; executable smoke tests cover the public commands.

No agent assignment, workflow automation or UI in this ticket.
