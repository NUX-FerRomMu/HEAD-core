# HEAD-006 — Persistent attention queue

Dependency: HEAD-001. Independent of tickets and providers.

`hea notification add --text TEXT [--priority 0..3] [--source TEXT]` records an unread activity (default priority 1). Provide `list`, `next`, `dismiss ID` and `requeue ID`. List and next order pending items by base priority plus one aging point per waiting hour, then oldest first. Next previews without removing. Requeue marks unread, lowers base priority to a minimum of zero and restarts aging. Dismissed items leave the queue. Output is JSON.

## Acceptance

- Added items have unique IDs and survive restarting. Empty text, invalid priorities and missing IDs fail cleanly.
- Ordering is deterministic with a controlled clock; aging lets an old low-priority item overtake a new high-priority one.
- Requeue/dismiss persist and never change other items. Empty next returns JSON null.
- Temporary SQLite tests cover persistence, ordering and transitions; state is project-local.

Automatic provider notifications and Neovim are separate integration work.
