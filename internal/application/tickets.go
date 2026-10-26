package application

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

type TicketRepository interface {
	Create(context.Context, domain.Ticket) (domain.Ticket, error)
	Get(context.Context, int64) (domain.Ticket, error)
	List(context.Context) ([]domain.Ticket, error)
	Edit(context.Context, int64, domain.TicketChanges) (domain.Ticket, error)
	Move(context.Context, int64, string) (domain.Ticket, error)
}

type Tickets struct{ Repo TicketRepository }

func (s Tickets) Create(ctx context.Context, title, description string, criteria []string) (domain.Ticket, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return domain.Ticket{}, fmt.Errorf("title must not be empty")
	}
	if criteria == nil {
		criteria = []string{}
	}
	now := time.Now().UTC()
	return s.Repo.Create(ctx, domain.Ticket{Title: title, Description: description, Criteria: criteria, State: "todo", CreatedAt: now, UpdatedAt: now})
}
func (s Tickets) Edit(ctx context.Context, id int64, change domain.TicketChanges) (domain.Ticket, error) {
	if change.Title != nil {
		title := strings.TrimSpace(*change.Title)
		if title == "" {
			return domain.Ticket{}, fmt.Errorf("title must not be empty")
		}
		change.Title = &title
	}
	if change.Title == nil && change.Description == nil && change.Criteria == nil {
		return domain.Ticket{}, fmt.Errorf("provide at least one field to edit")
	}
	return s.Repo.Edit(ctx, id, change)
}
func (s Tickets) Move(ctx context.Context, id int64, state string) (domain.Ticket, error) {
	if !domain.ValidTicketState(state) {
		return domain.Ticket{}, fmt.Errorf("invalid state %q; use todo, in_progress, review or done", state)
	}
	return s.Repo.Move(ctx, id, state)
}
func (s Tickets) Board(ctx context.Context) (map[string][]domain.Ticket, error) {
	rows, err := s.Repo.List(ctx)
	if err != nil {
		return nil, err
	}
	board := map[string][]domain.Ticket{}
	for _, state := range domain.TicketStates {
		board[state] = []domain.Ticket{}
	}
	for _, ticket := range rows {
		board[ticket.State] = append(board[ticket.State], ticket)
	}
	return board, nil
}
