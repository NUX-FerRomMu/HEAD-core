# Local RPC contract — ade.v1

## Transport and connection

JSON-RPC 2.0, UTF-8, one compact JSON object per newline. Strings escape embedded newlines. No batches. Maximum inbound/outbound frame is 1 MiB; larger artifact content is fetched in chunks. Request IDs are strings. Linux transport is a Unix socket in a 0700 user-owned directory, socket mode 0600; reject unsafe owners/permissions, refuse symlink endpoints, and do not expose TCP in MVP. The stdio bridge writes protocol frames only to stdout, diagnostics to stderr, and exits nonzero on connection loss. It does not own business state.

First request must be `system.hello` with `{protocol:"ade.v1", client:{name:string,version:string}}`. Reply includes `{protocol,server_version,capabilities:string[],latest_cursor:string}`. Incompatible versions return `UNSUPPORTED` and close. Unknown methods return JSON-RPC -32601. Invalid JSON/request/params use -32700/-32600/-32602. Other application errors use code -32000 and `data:{kind,details,retryable}`. Kinds: `NOT_FOUND`, `CONFLICT`, `INVALID_STATE`, `VALIDATION_FAILED`, `BUSY`, `FORBIDDEN`, `UNSUPPORTED`, `PROVIDER_UNAVAILABLE`, `GIT_CONFLICT`, `STALE_APPROVAL`, `CURSOR_EXPIRED`, `LIMIT_EXCEEDED`, `INTERNAL`. Redact process secrets from error details.

Every mutation has `command_id` (UUID). Updates have `expected_version` for the primary entity; creates do not. `attention.next` is an atomic selection with no client-selected entity/version; `project.open` is an idempotent registration with no expected version. Proposal application uses the proposal's expected version plus `source_versions`. Writer/candidate/validation commands use ticket version; chat/control use agent version; comments use comment version, except comment creation/request-changes/approval use ticket version. Integration retry uses integration version. Record the normalized request hash and result with command_id. Same ID/same payload returns original result, same ID/different payload is `CONFLICT`. Keep idempotency records for the lifetime of their project in MVP. Expected-version check happens after duplicate lookup. Bulk operations also check referenced versions described by their contract. Read timeouts default 10 s; mutations return accepted durable job IDs promptly and do not wait for provider/Git execution. A client timeout is an unknown outcome: retry the same command ID, never silently issue a new one.

Default result forms: single read `{item:<projection>,snapshot_cursor:string}`; mutation `{item:<updated projection>,job_id:string|null}`; list `{items:[],next_page_token:string|null,snapshot_cursor:string}`. Lists use stable `(created_at,id)` ascending order unless explicitly defined otherwise. `limit` defaults 50/max 200; opaque page tokens capture the filter, sort and snapshot boundary. Invalid/mismatched tokens are rejected. For MVP, cache the ordered list's projected rows in a bounded server snapshot for 60 seconds rather than implement historical SQL queries. Bound cached snapshots to 64 MiB per daemon and reject/expire tokens with `CURSOR_EXPIRED` on eviction/restart; no long-lived SQLite read transaction. Additional pages return that snapshot or require restart, never silently mix revisions. Attention snapshots also retain their evaluation time.

## Method registry

The fields listed below are method-specific `params`; add the mutation envelope above wherever M is shown. `id` field names are explicit to prevent cross-entity confusion. `project_id` filters all project collections. No RPC accepts arbitrary executable commands except human-authored validated configuration.

| Method | R/M | Fields → result / rule | Owner ticket |
|---|---|---|---|
| system.hello | R | protocol, client → negotiated handshake | ADE-004 |
| system.status | R | none → daemon version, uptime, jobs summary | ADE-004 |
| project.open | M | path, target_branch? → Project; verify repository and branch | ADE-005 |
| project.list | R | pagination → Project[] | ADE-005 |
| config.get | R | project_id? → redacted effective config + digest | ADE-005 |
| config.update | M | project_id?, patch, expected_version → effective config; no secrets echoed | ADE-005 |
| epic.create / epic.update | M | project_id,title,description / epic_id,patch → Epic | ADE-006 |
| epic.list | R | project_id,pagination → Epic[] with child progress | ADE-006 |
| ticket.create | M | project_id,kind,title,description,acceptance_criteria,epic_id?,parent_id?,dependency_ids? → backlog Ticket | ADE-006 |
| ticket.update | M | ticket_id,patch → Ticket; enforce editable fields/states | ADE-006 |
| ticket.get / ticket.list | R | ticket_id / project_id,state?,epic_id?,pagination → Ticket / Ticket[] | ADE-006 |
| ticket.ready / ticket.unready | M | ticket_id → Ticket | ADE-007 |
| ticket.assign | M | ticket_id,role,provider_profile,model?,start? → Ticket + assignment; start defaults false | ADE-017 |
| ticket.cancel | M | ticket_id,reason → Ticket or pending cancellation job | ADE-017 |
| agent.start / agent.pause / agent.resume / agent.stop | M | agent_id → Agent + job; stop retains history | ADE-017 |
| agent.get / agent.list | R | agent_id / ticket_id,pagination → Agent / Agent[] | ADE-017 |
| agent.send | M | agent_id,text → persisted Message{id,delivery_state}; max 64 KiB | ADE-019 |
| agent.messages | R | agent_id,pagination → Message[] chronological | ADE-019 |
| intervention.respond | M | activity_id,response:{choice_id,text?} → Activity; choice must match offered options | ADE-019 |
| backlog.proposal.get | R | proposal_id → validated proposal + diff | ADE-018 |
| backlog.apply / backlog.reject | M | proposal_id,source_versions / proposal_id,reason → proposal outcome | ADE-018 |
| planning.complete | M | ticket_id → Ticket | ADE-018 |
| candidate.publish | M | ticket_id,lease_fence → Candidate + Ticket + validation job if assigned | ADE-013 |
| candidate.refresh_base | M | ticket_id → Ticket fixing, target-update activity | ADE-023 |
| diff.get | R | candidate_id,path?,context_lines? → changed files, hunks, full head/base hashes; default context 3/max 20 | ADE-013 |
| snapshot.read | R | candidate_id,side,path,offset?,length? → text or binary metadata; length max 64 KiB | ADE-013 |
| writer.acquire / writer.release | M | ticket_id / ticket_id,lease_fence → human lease state | ADE-014 |
| writer.commit | M | ticket_id,lease_fence,message → commit SHA; only human lease; does not publish implicitly | ADE-014 |
| validation.start | M | ticket_id,candidate_id → Validation + job; assigned tester required | ADE-020 |
| validation.record | M | ticket_id,candidate_id,config_digest,checks,coverage?,artifact_ids → human Validation with explicit provenance; same eligibility rules | ADE-020 |
| validation.get / validation.list | R | validation_id / ticket_id,pagination → Validation / Validation[] | ADE-020 |
| review.comment | M | ticket_id,candidate_id,path?,side?,line?,body → Comment | ADE-021 |
| review.comments | R | ticket_id,candidate_id?,pagination → Comment[] including replies and historical anchors | ADE-021 |
| review.respond | M | comment_id,body → Comment + response; response does not resolve | ADE-021 |
| review.resolve | M | comment_id → Comment; human only | ADE-021 |
| review.request_changes | M | ticket_id,candidate_id,comment_ids,force:boolean → Ticket fixing | ADE-021 |
| review.approve | M | ticket_id,candidate_id,validation_id → Approval + integration job | ADE-022 |
| integration.get / integration.retry | R/M | integration_id → Integration / Integration + job; retry still checks approval | ADE-023 |
| attention.list | R | project_id? (omitted=all),state?,pagination → ordered Activity[] + evaluated_at | ADE-008 |
| attention.next | M | project_id? → highest eligible Activity marked seen, or null; no domain action | ADE-009 |
| attention.seen / attention.requeue | M | activity_id → Activity | ADE-009 |
| attention.resolve | M | activity_id,resolution → Activity; informational types only; actionable types require domain command | ADE-009 |
| context.get | R | ticket_id → summary, state, candidate, agents, pending activities, comments, latest validation, resource refs | ADE-010 |
| artifact.read | R | artifact_id,offset,length → base64 chunk,eof,checksum; max 64 KiB | ADE-024 |
| events.subscribe | R | after_cursor,project_id? → subscription_id; then events notifications | ADE-004 |
| events.unsubscribe | R | subscription_id → closed boolean | ADE-004 |

`agent.stop` stops a run, not the assignment or ticket; `agent.pause` also suppresses automatic correction restarts until resume. `review.respond` via RPC is human-authored; provider responses flow through supervisor validation. Role reassignment reuses `ticket.assign` with `expected_version`, allowed only after old role runs stop. `project.relocate` is reserved and returns `UNSUPPORTED` in MVP.

## Events and reconnection

Notification method `events.item`, params `{subscription_id,event:{cursor,project_id,type,entity_id,entity_version,occurred_at,payload}}`. Event types are `<entity>.created`, `<entity>.updated`, and domain facts `candidate.published`, `validation.completed`, `integration.completed`, `run.interrupted`, `message.appended`. Payload is an invalidation `{changed_fields:string[],related_ids:{ticket_id?,agent_id?,candidate_id?,validation_id?,activity_id?}}`, never arbitrary stdout. Clients refetch projections instead of interpreting partial aggregate mutations. Consumers must tolerate new optional fields and unknown event types by refetching affected entities. Deletion is not exposed for durable domain records in MVP; terminal/resolved records remain queryable.

Events are at-least-once. Clients deduplicate by global cursor and ignore entity versions older than their cache. With a project filter cursors can jump; a gap is not evidence of loss. Retain all domain events for MVP, allowing later retention policy; handle `CURSOR_EXPIRED` anyway. On new connect obtain initial snapshots and their cursors, subscribe after the oldest cursor, and apply only newer entity versions. Subscription captures a high-water mark, replays durable events up to it, then tails without a gap. A slow client exceeding 256 queued notifications is disconnected with last acknowledged-by-client cursor retained client-side; reconnect/replay, never drop events silently. Core never blocks a write on a client.

Example mutation:

```json
{"jsonrpc":"2.0","id":"req-17","method":"ticket.ready","params":{"command_id":"9e789d9e-1ae0-40b4-b330-e59c9ef25135","ticket_id":"ticket-example","expected_version":3}}
```

Illustrative IDs in examples may be mnemonic; tests and actual records use UUIDs. RPC JSON numbers never carry 64-bit event cursors or opaque Git hashes. Lua must not lose precision by converting cursor strings to numbers.
