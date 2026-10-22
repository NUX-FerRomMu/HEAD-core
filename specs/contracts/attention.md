# Attention and context contract

Activity fields: `id,project_id,ticket_id?,agent_id?,kind,dedupe_key,severity,state,created_at,seen_at?,resolved_at?,requeue_count,version,summary,detail_ref?,actions[],context_ref`. States: `unread`, `seen`, `resolved`. Kinds: `provider_intervention`, `validation_failed`, `validation_error`, `review_ready`, `changes_requested`, `target_moved`, `integration_blocked`, `run_interrupted`, `assignment_needed`, `information`. Actions are typed intents `{id,label,method,target_ids}`; clients must render only known methods, never execute arbitrary code from an activity.

## Ordering

Defaults: critical=300, high=200, normal=100, low=0. Aging interval=300 seconds; aging step=10; requeue penalty=50 per requeue, with total penalty capped at 200. At evaluation time `now`:

```text
age = max(0, floor((now - created_at) / aging_interval))
score = severity_weight + age * aging_step - min(requeue_count * requeue_penalty, max_requeue_penalty)
order = score DESC, created_at ASC, id ASC
```

Age is never reset by preview, seen, deduplication, or requeue. Score is not stored as authority; compute from injected clock using persisted timestamps. Negative age clamps to zero after clock regression. Aging step/interval must be positive; severities must retain critical > high > normal > low; penalties nonnegative and capped. Config changes affect subsequent evaluations and are audited. Uncapped aging means a low item can eventually outrank a new critical item; see ADR-005. No bound against an infinite overload is promised.

`attention.list` includes unresolved unread+seen by default. Its pagination freezes `evaluated_at` and ordered IDs/versions for the page sequence. `next` selects the highest-ranked **unread** item, marks it seen atomically, and returns it without executing its action. If there are no unread items, return null with unresolved_seen_count; UI offers the seen filter. `requeue` marks an unresolved item unread and increments its penalty count. Requeue never changes the underlying domain condition. User can select any item directly. Quick preview is read-only; explicit navigation marks seen. `attention.seen` is idempotent.

Example at 2026-10-22T12:00:00Z: a low item created at 09:00 has score 360; a new critical item scores 300. A new normal item scores 100; after one requeue it scores 50 and remains unread. Equal scores break by oldest creation then lexicographic ID.

## Deduplication and resolution

One unresolved item per `(project_id,dedupe_key)`. Upsert refreshes summary/resource links and may raise severity; preserve original created_at and requeue_count. Example keys: `provider:<run_id>:<external_intervention_id>`, `validation:<candidate_id>:failed`, `review:<candidate_id>`, `integration:<integration_id>:blocked`. A resolved item remains history; a genuinely new occurrence uses a new occurrence ID/key. Idempotent retry of the same occurrence does not resurrect it.

An actionable item resolves only after its domain action succeeds: review_ready on approval/new candidate/cancel; provider intervention on acknowledged response or terminal run; validation_failed on successor candidate/cancel; assignment_needed on role assignment/cancel. Error/recovery items resolve when a retry starts or human explicitly acknowledges a terminal condition as described by their offered action. Acknowledging an informational item may use `attention.resolve`; it cannot approve code, suppress a required correction, or make a failed check pass. Pending responses remain visible until acknowledged. Domain state and attention resolution share a transaction.

## Context projection

`context.get(ticket_id)` returns `ticket_version`, generated_at, a deterministic factual summary (title, state, last successful action, current block, next human actions), current candidate/base, active/idle agents, latest validation, open review obligations, unresolved activities, and typed resource refs `{kind,id,label}` for chat/diff/tests/artifacts. An optional agent-authored narrative is clearly attributed with its timestamp/context revision; it never overwrites factual state. Projection must work without any model call.

The Neovim client keeps the latest context panel reachable while switching views, stores per-session cursor/window preferences locally in the client, and refreshes on events. Core retains domain context across sessions; it does not persist Neovim window internals. Missing/deleted resources display an unavailable state with a link back to the ticket, not a crashed panel.
