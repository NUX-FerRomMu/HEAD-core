package main

import "github.com/NUX-FerRomMu/HEAD-core/internal/cli"

func init() {
	commands["agent"] = cli.Command{Summary: "check Codex login or run a task with streaming output", Run: cli.AgentCommand}
}
