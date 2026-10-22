# Provider port and execution contract

Core integrates a tool, not a model-specific API. A profile contains `id`, provider name, executable path, approved argv/config, optional model label, capability requirements, max run duration and environment-name allowlist. Never embed credentials in tickets or logs. A Sol model can be selected where the installed provider actually supports its model identifier; this document does not invent that identifier.

## Go port semantics

The application owns an interface with operations equivalent to:

```text
Probe(ctx, Profile) -> Capabilities + version + actionable error
Start(ctx, RunRequest) -> Handle
Send(ctx, Handle, Message) -> delivery acknowledgement or UNSUPPORTED
Respond(ctx, Handle, InterventionResponse) -> acknowledgement
Stop(ctx, Handle) -> confirmed stopped / uncertain
Events(ctx, Handle) -> ordered stream of ProviderEvent
```

Exact Go signatures belong to ADE-016; retain these semantics. Context cancellation must release readers and terminate/wait for the process tree with bounded graceful/forced phases. Process identity includes start-time/instance token, not PID alone. Do not kill a reused PID during restart recovery. Capabilities: `structured_events`, `live_messages`, `resume`, `interactive_approval`, `cancellation`, `usage_reporting`. Probe results persist with profile/version, and unsupported behavior is visible to the client.

`RunRequest` contains run/agent/ticket IDs, role, working directory, assignment/profile snapshot, spec/handoff text, context revision, candidate/base if relevant, writer fence if developer, and permitted ADE tool actions. Input budget defaults to 64 KiB UTF-8 with deterministic truncation of older chat; never truncate acceptance criteria, current forced obligations, or current validation failures. If mandatory input alone exceeds budget, fail with intervention and links rather than silently omitting requirements.

Events: `started`, `message`, `progress`, `intervention`, `proposal`, `candidate_proposed`, `validation_proposed`, `completed`, `failed`, `usage`. Every event has run_id, monotonic per-run sequence, timestamp and typed payload. Duplicate sequences are ignored; malformed/out-of-order events fail that run with diagnostic artifacts and release/uncertain lease handling. A completed process without a valid role output is failed, not inferred successful from prose. Provider stdout is untrusted data, never executable ADE commands.

Role output schemas:

- Developer: `candidate_proposed:{head_sha,summary,changed_paths,checks_claimed[]}`. Supervisor independently verifies Git and schedules validation; claimed checks are informational only.
- Backlog: `proposal:{schema_version:"ade.v1",planning_ticket_id,source_versions,epics[],tickets[]}`. Ticket entries include local_id, title, description, acceptance_criteria, dependency_local_ids, source_links, constraints. No assignment or ready status fields accepted.
- Tester: `validation_proposed:{candidate_id,summary,findings[],suggested_patch_artifact_id?}`. Configured runner results determine pass/fail; agent prose cannot override an exit code.
- Intervention: `{external_id,kind,prompt,choices:[{id,label}],requires_text:boolean}`. Core maps choices to one durable activity and responses back to exactly that run/intervention. Unknown or stale responses fail safely.

## Session policy

A logical developer is ticket-scoped and addressable through review. For live-message support, deliver to the running turn and record acknowledgement. Otherwise persist chat as queued, show that status, and deliver during the next supported turn; no silent claim of immediate delivery. At most one active turn per agent. An idle developer can receive a new/resumed turn after the human command. Closed backlog/tester agents keep readable transcripts; a new request requires a new assigned run.

Use resume only when the compatibility report verifies it for the pinned version and the profile opts in. Default restart strategy is a fresh provider process with a durable handoff containing ticket spec, current candidate, summary, outstanding obligations, recent chat, test results and artifact links. Store provider session refs separately and treat missing/expired refs as a visible fallback to fresh context. Billing/usage fields are nullable; never estimate cost from missing usage.

## Human authority and trust

Only the human can assign, resolve forced review obligations, and approve merge. Provider-facing tools offer a narrower dispatcher for role outputs and intervention, not the general human RPC command catalog. OS permissions and provider sandbox options are defense in depth. A same-user executable with arbitrary shell access can bypass local controls; ADE does not claim adversarial isolation. Do not grant providers secrets they do not need. Untrusted repositories require a suitable external sandbox, beyond this MVP.

## Codex evidence gate

ADE-015 must inspect the installed executable/help/documented structured interface and produce versioned event fixtures and a capability matrix. Use an actual supported interface such as a tool server protocol if available; its name/flags are not specified until verified. Probe spawn, chat delivery, interruption, approval response, error and process exit. Record which operations require a new turn. ADE-016 implements the generic port/fake regardless; ADE-019 implements the actual adapter only against the verified fixtures. Missing tool/authentication is a blocked live check, not a successful compatibility result.
