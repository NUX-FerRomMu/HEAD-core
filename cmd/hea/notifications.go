package main

import "github.com/NUX-FerRomMu/HEAD-core/internal/cli"

func init() {
	commands["notification"] = cli.Command{Summary: "add, order, requeue and dismiss attention items", Run: cli.NotificationCommand}
}
