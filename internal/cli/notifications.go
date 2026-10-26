package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NUX-FerRomMu/HEAD-core/internal/application"
	"github.com/NUX-FerRomMu/HEAD-core/internal/store"
)

func NotificationCommand(ctx context.Context, e *Env, args []string) error {
	if len(args) == 0 || args[0] == "--help" {
		fmt.Fprintln(e.Out, `Usage: hea [OPTIONS] notification COMMAND
  add --text TEXT [--priority 0..3] [--source TEXT]
  list
  next
  dismiss ID
  requeue ID`)
		return nil
	}
	op := args[0]
	fs := Flags("notification " + op)
	var text, source string
	var priority int
	if op == "add" {
		fs.StringVar(&text, "text", "", "notification text")
		fs.StringVar(&source, "source", "manual", "source")
		fs.IntVar(&priority, "priority", 1, "priority: 0..3")
	}
	count := 0
	switch op {
	case "add", "list", "next":
	case "dismiss", "requeue":
		count = 1
	default:
		return fmt.Errorf("unknown notification command %q", op)
	}
	if err := Parse(fs, args[1:], count); err != nil {
		return err
	}
	var id int64
	var err error
	if count == 1 {
		id, err = strconv.ParseInt(fs.Arg(0), 10, 64)
		if err != nil || id <= 0 {
			return fmt.Errorf("invalid notification ID %q", fs.Arg(0))
		}
	}
	db, err := e.DB()
	if err != nil {
		return err
	}
	repo, err := store.NewNotifications(ctx, db)
	if err != nil {
		return err
	}
	service := application.Notifications{Repo: repo}
	var result any
	switch op {
	case "add":
		result, err = service.Add(ctx, text, source, priority)
	case "list":
		result, err = service.List(ctx)
	case "next":
		result, err = service.Next(ctx)
	case "dismiss":
		result, err = repo.Dismiss(ctx, id)
	case "requeue":
		result, err = service.Requeue(ctx, id)
	}
	if err != nil {
		return err
	}
	return e.JSON(result)
}
