package main

import "github.com/NUX-FerRomMu/HEAD-core/internal/cli"

func init() {
	commands["ticket"] = cli.Command{Summary: "create, edit, move and view the backlog/Kanban", Run: cli.TicketCommand}
}
