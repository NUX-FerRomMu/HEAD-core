# Domain contract — ade.v1

Normative for Core, all clients, and providers. IDs are opaque UUID strings except full Git object IDs and decimal event cursors. Timestamps are RFC3339 UTC with millisecond precision. Entity `version` is a positive integer incremented on every persisted mutation. Empty optional values use JSON `null`; absent mutation fields mean unchanged. No filesystem or provider process is authoritative for ticket state.

## Records

| Record | Required fields and constraints |
|---|---|
| Project | id, canonical_root, common_git_dir, display_name, target_branch, version, created_at. Identity is local common Git directory; configuration is installation-local. |
| Epic | id, project_id, title, description, version, created_at. Progress is derived from child tickets; no parallel epic workflow engine. |
| Ticket | id, project_id, epic_id?, parent_id?, kind (`implementation` or `planning`), title, description, acceptance_criteria[], state, dependency_ids[], version, created_at, updated_at. No cross-project dependency or parent cycles. |
| Assignment | id, ticket_id, role (`developer`, `tester`, `backlog`), provider_profile, model?, assigned_by=`human`, created_at. One active assignment per role per ticket. Model is a profile setting, not hard-coded to Sol. |
| Agent | id, ticket_id, assignment_id, role, state, context_revision, provider_session_ref?, last_error?, version. One logical developer per active assignment. |
| Run | id, agent_id, state, candidate_sha?, context_revision, started_at?, ended_at?, exit_reason?, operation_key, artifact_ids[]. A run is one provider process/turn. |
| Candidate | id, ticket_id, head_sha, base_sha, created_by, created_at. Immutable; base is the target observed for this candidate. |
| Validation | id, ticket_id, candidate_id, head_sha, base_sha, config_digest, status, checks[], coverage?, artifact_ids[], started_at, ended_at?, source (`runner` or `human`). Immutable result once terminal. |
| Comment | id, ticket_id, candidate_id, path?, side? (`base` or `head`), line?, body, author, state (`open`, `answered`, `resolved`, `obsolete`), forced, version. Line is 1-based within that snapshot; binary files allow file comments only. |
| Approval | id, ticket_id, candidate_id, head_sha, base_sha, validation_id, config_digest, actor=`human`, created_at, revoked_at?. Immutable scope; revocation is audited. |
| Integration | id, ticket_id, approval_id, expected_target_sha, candidate_sha, state (`pending`, `running`, `succeeded`, `blocked`, `unknown`), resulting_sha?, error?. Durable journal. |
| WriterLease | ticket_id, owner_type (`developer` or `human`), owner_id, fence, acquired_at, state (`active`, `releasing`, `uncertain`). Only one active/uncertain lease per ticket. |
| Event | cursor, project_id?, type, entity_id, entity_version, occurred_at, payload. Cursor is a global monotonic decimal string; gaps allowed. |

Attention, RPC and provider record details are specified in their separate contracts. SQL tables may split records internally but must preserve these invariants. Store input snapshots (ticket spec version, config digest, model/profile, candidate) on runs so later config edits cannot rewrite history.

All mutable records used by RPC concurrency (including Integration, Assignment, Validation while running, proposals and config) include `version` and `created_at` even when omitted from the abbreviated table. Append-only values such as candidates/terminal results have fixed version 1. Record reads include IDs needed to retrieve related resources; no UI must reconstruct a key from a display label.

## Validation and message payloads

Validation status: `queued`, `running`, `passed`, `failed`, `error`, `cancelled`. A check is `{id,name,argv_snapshot,required,status,exit_code?,duration_ms,artifact_ids[],error_kind?}`; status is `passed`, `failed`, `error`, `cancelled` or `pending`. Pass requires every configured required check to be present and passed, no source mutation, and a matching configuration digest. Nonzero program exit means failed; spawn failure, signal/timeout and invalid required coverage input mean error. An optional check may be reported without gating pass. Human records provide check IDs/status/exit code/evidence text; Core looks up configured command identity rather than accepting invented required checks.

Coverage is nullable or `{format:"lcov",lines:{covered,total},branches:{covered,total}|null,artifact_id}`. Initial parser supports LCOV only; configured absent/malformed required coverage yields error; unconfigured coverage is null. Reject negative counts, covered > total and stale files: the configured coverage path must be removed/verified absent before execution and read afterward inside that validation workspace. Coverage is informational unless a human config explicitly sets a minimum line percentage; that threshold contributes to config digest and pass eligibility. No universal language-specific test discovery is implied.

Message is `{id,agent_id,run_id?,author:"human"|"agent",text,delivery_state:"queued"|"sent"|"acknowledged"|"failed",created_at,sequence,version,error?}`. Agent-authored durable output is acknowledged on ingestion; human input acknowledgement means accepted by the provider, not that it has completed the requested change. Transcript ordering uses per-agent sequence with ID tie-breaker, not uncertain provider clocks. Comment responses are append-only `{id,comment_id,author,body,created_at}`.

## Ticket states and commands

States: `backlog`, `ready`, `assigned`, `doing`, `committed`, `fixing`, `done`, `cancelled`. A terminal ticket cannot be reopened in MVP; create a linked follow-up. Pause/block/interrupted are agent/attention conditions, not additional Kanban states.

| Command/event | From → to | Guard and effect |
|---|---|---|
| human `ticket.ready` | backlog → ready | Title, description, at least one acceptance criterion; dependencies exist and graph is acyclic. |
| human `ticket.unready` | ready → backlog | No assignments. |
| human `ticket.assign` developer/backlog | ready → assigned | All dependencies done. Role matches kind. User chooses provider profile; tester can be assigned in same transaction or later. |
| human `ticket.assign` tester | assigned/doing/committed/fixing → same | Implementation kind only; may enqueue validation if current candidate exists. Never infer tester assignment. |
| developer/backlog `agent.start` | assigned → doing | Explicit role assignment, available provider, worktree/lease if developer. Human may start now or choose `start=true` on assignment. |
| developer `agent.start/resume` | doing/fixing → same | Explicit active assignment and fresh writer lease; paused state cleared only on human resume. An idle developer in committed may receive read-only discussion; writing requires request_changes or refresh_base first. |
| tester `agent.start` | committed → committed | Assigned tester and current immutable candidate; equivalent to scheduling validation for that candidate, never sets ticket doing. |
| human `writer.acquire` | ready → doing; committed → fixing; doing/fixing → same | Implementation only. Ready requires all dependencies done and provisions the normal isolated workspace; no provider assignment required for manual work. Committed revokes approval and invalidates validation eligibility until a new candidate is published. Grant only after previous writer is confirmed stopped. |
| `candidate.publish` | doing/fixing → committed | Clean ticket worktree; full commit differs from base and descends from recorded base; writer stopped/released; candidate captures exact head/base. Invalidate older approval/validation eligibility; schedule validation only for an assigned tester. |
| validation fails | committed → fixing | Result belongs to current head/base/config; enqueue developer correction and attention. Old results remain history without moving state. |
| validation error/cancel | committed → committed | Infrastructure failure does not prove code failure; attention asks retry/configuration. |
| human `review.request_changes` | committed → fixing | Must target current candidate; forced comments are obligations. Revoke approval. Resume assigned developer only if not paused by human. |
| human `candidate.refresh_base` | committed → fixing | Target advanced. Revoke approval; record target-update reason. Developer/human updates branch before publishing a fresh candidate. |
| human `review.approve` | committed → committed | Current passing validation and exact candidate/base/config, target still at base, no open forced comments, no writer. Persist approval and schedule integration atomically. |
| integration success | committed → done | Target ref equals approved candidate; journal verifies the approved target was the base. Close developer logical agent and resolve associated pending activities. |
| human `planning.complete` | doing → done | Planning ticket only; backlog proposals accepted/rejected; no uncommitted proposal application. No Git validation/merge required. |
| human `ticket.cancel` | any nonterminal → cancelled | Stop runs and confirm writers stopped first. If unable, cancellation remains a pending command/job and lease uncertain; do not pretend work is stopped. Keep history/worktrees. |

Metadata edits use optimistic versions. Editing description/acceptance criteria/dependencies on `assigned`, `doing`, `committed`, or `fixing` returns `INVALID_STATE`; pause/cancel and create a revised ticket for MVP. Labels/title/epic links may change without rewriting the execution spec. Dependencies cannot become done through cancellation. Human can reassign only an idle/blocked role after stopping old runs and releasing its lease; preserves assignment/agent history. Reassignment does not change ticket state.

`candidate.publish` is an application command available to a human or an ADE-controlled supervisor with the relevant run/lease; provider output is a proposal, never a directly authorized database command. A provider cannot approve, assign, integrate, or mark done. Actor identity comes from application dispatch, never a caller-supplied JSON `actor` string. Local-user RPC is a human control interface; it is not a sandbox against an arbitrary process running as that OS user.

## Agent/run lifecycle

Agent states: `idle`, `running`, `waiting_input`, `paused`, `interrupted`, `closed`. Run states: `queued`, `running`, `waiting_input`, `succeeded`, `failed`, `cancelled`, `interrupted`. Developer logical agents persist through review and correction, then close on done/cancel/reassignment. Finite backlog/testing agents close after results, with a new run/agent as appropriate for a later assigned job. Assignment persists until explicitly replaced/removed. Chat targets a specific agent ID; closed agents are readable but do not accept new live messages.

Only one running/waiting-input run per logical agent. Ticket writes require the exact lease fence. A run can be logically interrupted at restart while its OS process may still exist: lease remains uncertain until reconciliation proves it dead. No automatic takeover on elapsed heartbeat alone.

## Backlog proposals

Backlog agents output a versioned proposal with `proposal_id`, `planning_ticket_id`, input ticket versions, proposed epics/tickets, local proposal IDs for references, acceptance criteria and dependencies. Core validates size, references and acyclicity. Human `backlog.apply` previews and atomically inserts/updates the chosen proposal as a whole; no partial graph acceptance in MVP. Reject stale source versions. Children carry inherited source links, relevant constraints, acceptance criteria, and explicit scope; a title alone is not sufficient. Proposals cannot assign agents, make tickets ready, or modify in-progress specs.

## Mandatory invariants

1. State mutation, audit event, and background job insertion succeed in one DB transaction or not at all.
2. At most one development writer per ticket, and one integration operation per repository.
3. Validation and approval refer to head **and** base **and** validation configuration digest; any difference makes them ineligible.
4. No done implementation ticket without a verified successful integration journal.
5. No provider event can manufacture human authority.
6. Retries do not duplicate tickets, attention activities, runs, or integrations.
7. Every interrupted/failed operation has an inspectable result and, when action is needed, a deduplicated attention item.
