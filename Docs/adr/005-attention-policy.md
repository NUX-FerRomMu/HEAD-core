# ADR-005: Deterministic attention

Status: accepted.

Queue priorities are deterministic: severity weight plus uncapped aging, with stable ties. Deduplicate unresolved items by semantic key, keep durable resolution, and support unread requeue with a finite penalty. No LLM chooses what the user may see or whether a decision is resolved. `next` is advisory navigation, not execution or forced scheduling.

Alternatives: static severity can starve low-priority items; LLM ranking is hard to reproduce; capping aging below urgent priority preserves starvation. Uncapped aging permits an old low-priority item to outrank new urgent work, which is the intended starvation-prevention tradeoff. Critical badges and filters remain visible independently of rank.

The formula, defaults, ordering and fake-clock examples live in the [attention contract](../../specs/contracts/attention.md).
