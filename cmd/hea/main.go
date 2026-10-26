package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/NUX-FerRomMu/HEAD-core/internal/cli"
)

// Feature files register commands during program initialization.
var commands = map[string]cli.Command{}
var version = "dev"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	os.Exit(cli.Run(ctx, os.Args[1:], os.Stdout, os.Stderr, version, commands))
}
