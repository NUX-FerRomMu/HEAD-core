# Operations, safety boundaries, and recovery

One local user owns the daemon. Unix socket permissions isolate other users, not arbitrary processes running under the same account. ADE coordinates its own agents/editors; it is not a container sandbox. Provider-native sandbox/approval features should be used when available and verified, but no untested isolation guarantees appear in the UI.

## Installation and configuration

ADE-005 defines a schema-versioned local JSON config with defaults: target branch resolved from local symbolic remote HEAD when available, otherwise explicit selection; provider profiles; branch template; test command argv arrays; per-check timeouts (default 10 minutes); maximum 2 simultaneous developer runs and 2 validation jobs per installation; max provider run duration 60 minutes; queue defaults from its contract; max frame 1 MiB; per-run logs 20 MiB with truncation marker. Starting with no validation commands configured must yield `validation_error/configuration_required`, never green validation. Configuration can be global and overridden per registered project locally. A model/profile change affects future runs only; a validation-command/config digest change invalidates approval eligibility for current candidates.

Only human configuration can supply executable paths, arguments and allowed environment names. Invoke processes using argv, no shell interpolation by default. An explicit shell command profile is a later extension, not implicit evaluation. Provider-required credentials are inherited only from configured environment names or provider-owned auth storage; don't copy credential contents into ADE. Persist credential references, not values. Redact known sensitive environment/argument values in logs and errors. Test redaction without asserting it can identify every arbitrary secret in model prose.

## Process and storage lifecycle

`ade daemon` runs in foreground; optional auto-start by `ade`/bridge starts one user daemon and waits on readiness with a bounded timeout. Use OS locking and race-safe socket discovery; a stale socket alone does not justify killing a process. Structured diagnostics go to stderr/files, never bridge stdout. SIGTERM stops accepting new work, stops/records runs, flushes state, then exits. A forced kill is covered by startup reconciliation.

Schema migrations run before accepting clients under the daemon lock. Back up via SQLite online backup/safe consistent snapshot (not copying a live WAL database file); record schema version and checksum. A failed migration keeps the prior database recoverable and refuses normal startup. Newer-schema databases opened by an older binary are rejected with guidance, not downgraded. ADE-003 establishes the mechanism; ADE-032 documents recovery. Do not auto-delete user data to solve corruption.

At startup inspect pending jobs and leases. Match process identities without trusting PID alone. Mark orphaned runs interrupted; only release leases after confirmed termination. Reconcile worktree creation from Git metadata, validation from exact candidate/result records, and integration from its journal. Unknown ownership creates attention and blocks further writes for that ticket/repository; reads and other projects continue. Retry only idempotent operations with bounded backoff, at most three automatic attempts for transient infrastructure errors, then require intervention. No automatic retry of human decisions, failed tests, stale approvals, or uncertain Git mutations.

## Retention and observability

Persist tickets, decisions, chat metadata and domain events for MVP. Bound provider log files while recording truncation and checksum. Cleanup completed validation worktrees after terminal results and safe process exit; preserve referenced test artifacts and source candidate refs. Explicit maintenance may clean unreferenced owned artifacts after preview; automatic branch deletion is out of scope. Publish disk usage and clear disk-full errors. SQLite busy/disk-full failure must roll back both state and event writes.

Correlate logs using project/ticket/agent/run/job/command IDs; never put high-volume output in event notifications. Expose status and diagnostic export with versions, redacted config, recent errors, unfinished jobs and schema version. Artifacts are accessed by registered ID with size-limited reads. No generic filesystem reader is exposed to clients.

Manual equivalents: human ticket authoring/refinement, explicit assignment/start/stop, configured check execution or provenance-labelled validation recording, human writer lease/commit/candidate publishing, review/forced changes/approval, integration retry and informational queue resolution. All use the same domain guards as automation. Human-entered validation is labelled as such and requires results for every required configured check; it is not a hidden bypass command.
