# HEAD-003 — Attention queue

Epic: E2. Depends on: HEAD-002. Blocked.

## Outcome and scope

A human can add, list, preview, select and requeue ticket-linked activities through the CLI. Follow [architecture](../../Docs/architecture.md) and [agent instructions](../AGENTS.md).

Persist activity ID, ticket, summary, priority, creation time, unread/seen/resolved state and requeue count. For this increment, use `score = priority + floor(max(0, age_seconds) / 300) - min(requeue_count, 5)` with priority 0–3. Rank highest score first, then oldest creation time, then ID. Inject the clock; defer configurable scoring and automatic producers.

`next` atomically selects the highest unread item and marks it seen. Preview is read-only. Requeue makes it unread, increments its penalty, and preserves its age. Users may open any item. These initial activities are informational; resolving one changes no ticket state.

## Acceptance

1. At 30 minutes old, priority 0 scores 6 and outranks a new priority-3 item. Requeue lowers its score without resetting age; ties are stable.
2. Repeated `next` skips seen items and returns empty when no unread items remain. Two concurrent selections cannot claim the same unread item.
3. Queue state survives restart; preview changes nothing. Direct navigation remains available regardless of rank.

## Verification and handoff

Use fake-clock ordering tests and temporary SQLite tests for persistence and atomic selection. Run Go tests and vet; document the added CLI commands. Provider interventions and approval actions are later work, not generic queue-resolution shortcuts.
