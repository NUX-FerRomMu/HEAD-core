# Getting started

Initial platform: Linux. Install Go 1.26+ (the pinned toolchain is selected by `go.mod`), Neovim 0.10+, and the Codex CLI.

## Build and connect

```sh
mkdir -p "$HOME/.local/bin"
go build -o "$HOME/.local/bin/hea" ./cmd/hea
export PATH="$HOME/.local/bin:$PATH"
hea --help
codex login
hea agent status
```

HEAD uses Codex's existing configuration, model and login. `agent status` checks local login presence; a task verifies server access. If Codex reports an expired or unrefreshable token, renew its login with `codex login`.

## CLI

Run in a project directory, or put `--project /path/to/project` before the command. Outputs are JSON; agent output is JSONL.

```sh
hea ticket create --title "First task" --description "What to build" --criterion "Observable result"
hea ticket board
hea ticket get 1
hea ticket edit 1 --title "Updated task"
hea ticket move 1 in_progress
hea notification add --text "Review the first task" --priority 2
hea notification next
hea notification requeue 1
hea notification dismiss 1
hea agent run --prompt "Inspect this project and summarize its structure"
```

Use the returned IDs in subsequent commands. Ctrl-C cancels an agent and its child processes. Success, failure and cancellation create a notification. Ticket moves remain manual. Feature help: `hea ticket --help`, `hea notification --help`, `hea agent --help`.

State lives in `$XDG_STATE_HOME/hea` or `~/.local/state/hea`, separated by canonical project root. A Git worktree is a separate project. Override with `hea --data-dir /path/to/state ...`; no state is written into the checkout.

## Neovim

Add the checkout to your Neovim configuration:

```lua
vim.opt.rtp:append("/absolute/path/to/HEAD-core")
require("hea").setup({
  cmd = vim.fn.expand("~/.local/bin/hea"),
  -- project = "/absolute/path/to/project", -- defaults to the current directory
  -- data_dir = "/path/to/state",
  -- model = "your-codex-model",
})
```

- `:HEADBoard`: four Kanban columns. `a` creates; `e` edits title, description or criteria; `m` moves; Enter opens details; `x` runs Codex on the selected ticket. Select cards with normal cursor movement. Field editors save with Ctrl-S; criteria use one line each.
- `:HEADNotifications`: ordered attention items. `a` adds, `n` selects next, Enter previews, `x` dismisses, `u` requeues at lower priority. Waiting items gain an aging point each hour.
- `:HEADRun [prompt]`: run Codex and stream output. `:HEADStop` cancels; exiting Neovim also cancels. One task runs per client at a time.
- `:HEADStatus`: check the installed CLI and local login. In board/queue buffers, `r` refreshes and `q` closes. Open notification views refresh when a task finishes.

This client invokes the Go CLI asynchronously. The planned daemon/RPC transport and persistent conversations are later work.

## Verify

```sh
CGO_ENABLED=0 go test ./...
go vet ./...
./tests/neovim.sh
```

The Neovim test builds the real executable and uses an isolated fixture provider; it requires no Codex account and changes no real project data. Live provider access is a separate check.
