# Verification strategy

Every ticket supplies observable examples, not just an implementation checklist. Use standard Go `testing`, `go test -race` on supported Linux, temporary SQLite databases, temporary Git repositories and a fake clock. Integration tests may run local Git/process fixtures; ordinary CI must not require network, provider credentials, paid model calls, or a user's repository.

## Layers and commands

ADE-001 establishes these command contracts and pins the tools; they are proposed commands until implementation exists:

| Check | Command contract | Scope |
|---|---|---|
| Format | `gofmt -l` over tracked Go sources, output must be empty | Changed Go files; all in CI |
| Static analysis | `go vet ./...` | Core and CLI |
| Unit/integration | `go test ./...` | Fast deterministic suite, including temporary Git/SQLite |
| Race detection | `go test -race ./...` | Linux concurrency suite |
| Build | `go build ./cmd/ade` | Single binary |
| Lua format | `stylua --check lua plugin tests/nvim` | Added when plugin exists; pin version |
| Neovim tests | `nvim --headless -u tests/nvim/minimal_init.lua -l tests/nvim/run.lua` | Neovim baseline 0.11+, pinned CI version |
| End-to-end | `go test ./tests/e2e -run TestSelfHostingFlow -count=1` | Added in ADE-031; fake providers |
| Live Codex | Explicit opt-in script documented by ADE-019 | Versioned local capability check; never silently part of CI |

Do not run paths/commands before their owning ticket creates them, and do not mark checks passing when missing. Pin exact supported Go/Neovim/Git/driver versions in ADE-001 and evidence. Race detector can require C tooling even though the application SQLite driver does not; document CI prerequisite separately from runtime installation.

## Required scenarios

| Scenario | Assertion | Owner |
|---|---|---|
| Invalid transition, stale version, duplicate command | No partial mutation/event/job; deterministic error/result | ADE-002/003/007 |
| Concurrent writers, stalled process, crash | No second writer until termination verified; old fence rejected | ADE-012/014/024 |
| Target moves during review/integration | No overwrite; approval ineligible; new validation/review needed | ADE-022/023 |
| Tests pass for older candidate | Historical result retained; current candidate remains unvalidated | ADE-020 |
| Provider output attempts assignment/approval | Narrow dispatcher rejects authority escalation | ADE-016/019 |
| Message retry/provider disconnection | Durable pending/acknowledged status, no silent loss | ADE-019 |
| Attention aging/requeue/ties | Exact fixture ordering with fake clock; no age reset | ADE-008/009 |
| Subscription reconnect/slow consumer | Replay without mutation loss or server blocking | ADE-004/027 |
| Human edits while developer active | Editable view waits for ownership; snapshot review stays usable | ADE-014/029 |
| Restart after ref change but before DB completion | Journal converges once; does not integrate twice | ADE-023/024 |
| Backlog decomposition | Invalid/stale/cyclic proposal applies no children | ADE-018 |

Use failure injection at storage/process boundaries rather than arbitrary sleeps. Poll bounded conditions where process startup is genuinely asynchronous. Avoid mocks that merely repeat function calls; assert stored state, Git refs, process exit, protocol frames and user-visible views.

## Performance and usability budgets

Initial engineering targets, to measure rather than claim now: on a documented Linux reference machine with 10,000 tickets and 10,000 unresolved activities, local list/next/context p95 under 100 ms excluding rendering/provider work; terminal input/navigation remains responsive while output streams; daemon idle memory under 150 MiB and no goroutine/process growth after 100 fake runs. ADE-032 records hardware, datasets, p50/p95 and deviations. Tune SQL indexes or buffered output only when measurements justify it; a missed budget requires a documented release decision, not fabricated success.

## Definition of done

Ticket acceptance examples pass; relevant negative/failure cases pass; no unrelated behavior is changed; shared contracts/docs stay consistent; migrations and recovery are reviewed if touched; tests do not use secrets; handoff names exact commands and results. A blocked live provider check remains blocked even if fake tests pass. The final MVP gate additionally requires the human self-hosting rehearsal in ADE-033.
