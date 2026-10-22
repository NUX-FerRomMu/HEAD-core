# Chosen defaults and unresolved evidence

Implementation agents should use these defaults rather than ask broad architecture questions. A real incompatibility is reported with a narrow proposed decision and evidence.

| Topic | Decision/default | Evidence or follow-up |
|---|---|---|
| Language | Go Core/CLI; Lua plugin | Explicit user instruction; not open. |
| Deployment | Local daemon, one OS user, Linux first | Other platforms deferred, interfaces portable. |
| Persistence | SQLite via modernc.org/sqlite, explicit SQL | Bootstrap verifies pinned Go/driver compatibility. |
| Toolchain | Go 1.26 baseline, Neovim 0.11+ baseline | ADE-001 pins tested patch versions; do not claim latest. |
| Provider | Codex tool, structured protocol | ADE-015 is an evidence gate. Tool availability/authentication unknown in this planning task. |
| Sol selection | Provider profile model label | Verify actual available identifier; no invented CLI flag/model name. |
| Provider memory | Durable ADE handoff; resume opt-in when verified | No promise of transparent cross-restart model memory. |
| Agent assignment | Always human, including tester/backlog roles | Automatic validation allowed only after explicit tester assignment. |
| Testing changes | Tester proposes patches; developer/human applies them | Prevent certifying a modified candidate as the old SHA. |
| Integration | Exact-candidate fast-forward after human approval | Changed target requires new candidate/validation/approval. |
| Human editing | Exclusive writer lease with clean handoff | External same-user writes cannot be fully prevented. |
| Queue | Uncapped aging, capped requeue penalty | Fairness under finite load; formula in contract. |
| Ticket specification | Immutable execution fields once assigned | Revise through cancellation/follow-up for initial MVP. |
| Test commands | Human-configured argv required | Missing commands are an error, not a pass. |
| Git edge cases | Reject bare/unborn/submodule/LFS projects initially | Add platform/repo support as separate specs after MVP. |
| Language of docs | English; original Spanish preserved | Product semantics trace back to source, not translations alone. |

Main risks: Codex structured protocol/version churn; Git and SQLite cannot share an atomic transaction; external processes can bypass local leases; validation commands may be slow or nondeterministic; attention noise can overwhelm navigation. Mitigate through the early compatibility gate, durable journals/reconciliation, explicit uncertainty, deterministic tests, deduplication and measured self-hosting. None justifies introducing distributed services or autonomous assignment into this MVP.

Not validated by this planning deliverable: compiling selected dependencies; live Codex behavior; performance budgets; actual OS support; end-to-end usability. Tickets explicitly require evidence before those claims are made.
