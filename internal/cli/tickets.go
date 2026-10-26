package cli

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/NUX-FerRomMu/HEAD-core/internal/application"
	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
	"github.com/NUX-FerRomMu/HEAD-core/internal/store"
)

const ticketHelp = `Usage: hea [OPTIONS] ticket COMMAND
  create --title TEXT [--description TEXT] [--criterion TEXT ...]
  get ID
  list
  board
  edit ID [--title TEXT] [--description TEXT] [--criterion TEXT ... | --clear-criteria]
  move ID todo|in_progress|review|done`

func TicketCommand(ctx context.Context, e *Env, args []string) error {
	if len(args) == 0 || args[0] == "--help" {
		fmt.Fprintln(e.Out, ticketHelp)
		return nil
	}
	// Parse and validate the complete command before opening storage.
	op := args[0]
	rest := args[1:]
	var id int64
	var err error
	if op == "get" || op == "edit" || op == "move" {
		if len(rest) == 0 {
			return fmt.Errorf("%s requires an ID", op)
		}
		id, err = strconv.ParseInt(rest[0], 10, 64)
		if err != nil || id <= 0 {
			return fmt.Errorf("invalid ticket ID %q", rest[0])
		}
		rest = rest[1:]
	}
	fs := Flags("ticket " + op)
	var title, description string
	criteria := Strings{}
	var clear bool
	if op == "create" || op == "edit" {
		fs.StringVar(&title, "title", "", "title")
		fs.StringVar(&description, "description", "", "description")
		fs.Var(&criteria, "criterion", "acceptance criterion")
		if op == "edit" {
			fs.BoolVar(&clear, "clear-criteria", false, "remove all criteria")
		}
	}
	count := 0
	if op == "move" {
		count = 1
	}
	switch op {
	case "create", "get", "list", "board", "edit", "move":
	default:
		return fmt.Errorf("unknown ticket command %q", op)
	}
	if err := Parse(fs, rest, count); err != nil {
		return err
	}
	if clear && len(criteria) > 0 {
		return fmt.Errorf("--clear-criteria cannot be combined with --criterion")
	}
	db, err := e.DB()
	if err != nil {
		return err
	}
	repo, err := store.NewTickets(ctx, db)
	if err != nil {
		return err
	}
	service := application.Tickets{Repo: repo}
	var result any
	switch op {
	case "create":
		result, err = service.Create(ctx, title, description, criteria)
	case "get":
		result, err = repo.Get(ctx, id)
	case "list":
		result, err = repo.List(ctx)
	case "board":
		result, err = service.Board(ctx)
	case "move":
		result, err = service.Move(ctx, id, fs.Arg(0))
	case "edit":
		var change domain.TicketChanges
		fs.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "title":
				change.Title = &title
			case "description":
				change.Description = &description
			case "criterion":
				v := []string(criteria)
				change.Criteria = &v
			}
		})
		if clear {
			v := []string{}
			change.Criteria = &v
		}
		result, err = service.Edit(ctx, id, change)
	}
	if err != nil {
		return err
	}
	return e.JSON(result)
}
