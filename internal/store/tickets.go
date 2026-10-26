package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

type Tickets struct{ db *sql.DB }

func NewTickets(ctx context.Context, db *sql.DB) (*Tickets, error) {
	_, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS tickets (
 id INTEGER PRIMARY KEY AUTOINCREMENT,
 title TEXT NOT NULL CHECK(length(trim(title)) > 0),
 description TEXT NOT NULL,
 criteria TEXT NOT NULL,
 state TEXT NOT NULL CHECK(state IN ('todo','in_progress','review','done')),
 created_at TEXT NOT NULL,
 updated_at TEXT NOT NULL
 )`)
	if err != nil {
		return nil, err
	}
	return &Tickets{db: db}, nil
}

const ticketColumns = "id, title, description, criteria, state, created_at, updated_at"

func scanTicket(row interface{ Scan(...any) error }) (domain.Ticket, error) {
	var t domain.Ticket
	var criteria, created, updated string
	err := row.Scan(&t.ID, &t.Title, &t.Description, &criteria, &t.State, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return t, domain.ErrTicketNotFound
	}
	if err != nil {
		return t, err
	}
	if err = json.Unmarshal([]byte(criteria), &t.Criteria); err != nil {
		return t, err
	}
	t.CreatedAt, err = time.Parse(time.RFC3339Nano, created)
	if err != nil {
		return t, err
	}
	t.UpdatedAt, err = time.Parse(time.RFC3339Nano, updated)
	return t, err
}
func (r *Tickets) Create(ctx context.Context, t domain.Ticket) (domain.Ticket, error) {
	criteria, err := json.Marshal(t.Criteria)
	if err != nil {
		return t, err
	}
	return scanTicket(r.db.QueryRowContext(ctx, `INSERT INTO tickets(title,description,criteria,state,created_at,updated_at) VALUES(?,?,?,?,?,?) RETURNING `+ticketColumns, t.Title, t.Description, string(criteria), t.State, t.CreatedAt.Format(time.RFC3339Nano), t.UpdatedAt.Format(time.RFC3339Nano)))
}
func (r *Tickets) Get(ctx context.Context, id int64) (domain.Ticket, error) {
	return scanTicket(r.db.QueryRowContext(ctx, "SELECT "+ticketColumns+" FROM tickets WHERE id=?", id))
}
func (r *Tickets) List(ctx context.Context) ([]domain.Ticket, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT "+ticketColumns+" FROM tickets ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []domain.Ticket{}
	for rows.Next() {
		t, err := scanTicket(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}
func (r *Tickets) Edit(ctx context.Context, id int64, c domain.TicketChanges) (domain.Ticket, error) {
	var criteria any
	if c.Criteria != nil {
		raw, err := json.Marshal(*c.Criteria)
		if err != nil {
			return domain.Ticket{}, err
		}
		criteria = string(raw)
	}
	return scanTicket(r.db.QueryRowContext(ctx, `UPDATE tickets SET title=COALESCE(?,title),description=COALESCE(?,description),criteria=COALESCE(?,criteria),updated_at=? WHERE id=? RETURNING `+ticketColumns, c.Title, c.Description, criteria, time.Now().UTC().Format(time.RFC3339Nano), id))
}
func (r *Tickets) Move(ctx context.Context, id int64, state string) (domain.Ticket, error) {
	return scanTicket(r.db.QueryRowContext(ctx, "UPDATE tickets SET state=?,updated_at=? WHERE id=? RETURNING "+ticketColumns, state, time.Now().UTC().Format(time.RFC3339Nano), id))
}
