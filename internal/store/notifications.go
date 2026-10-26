package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

type Notifications struct{ db *sql.DB }

func NewNotifications(ctx context.Context, db *sql.DB) (*Notifications, error) {
	_, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS notifications (
 id INTEGER PRIMARY KEY AUTOINCREMENT,
 text TEXT NOT NULL CHECK(length(trim(text))>0),
 source TEXT NOT NULL,
 priority INTEGER NOT NULL CHECK(priority BETWEEN 0 AND 3),
 unread INTEGER NOT NULL,
 dismissed INTEGER NOT NULL DEFAULT 0,
 created_at INTEGER NOT NULL,
 queued_at INTEGER NOT NULL
 )`)
	if err != nil {
		return nil, err
	}
	return &Notifications{db: db}, nil
}

const notificationColumns = "id,text,source,priority,unread,dismissed,created_at,queued_at"

func scanNotification(row interface{ Scan(...any) error }) (domain.Notification, error) {
	var n domain.Notification
	var created, queued int64
	err := row.Scan(&n.ID, &n.Text, &n.Source, &n.Priority, &n.Unread, &n.Dismissed, &created, &queued)
	if errors.Is(err, sql.ErrNoRows) {
		return n, domain.ErrNotificationNotFound
	}
	if err != nil {
		return n, err
	}
	n.CreatedAt = time.Unix(0, created).UTC()
	n.QueuedAt = time.Unix(0, queued).UTC()
	return n, nil
}
func (r *Notifications) Add(ctx context.Context, n domain.Notification) (domain.Notification, error) {
	return scanNotification(r.db.QueryRowContext(ctx, "INSERT INTO notifications(text,source,priority,unread,created_at,queued_at) VALUES(?,?,?,?,?,?) RETURNING "+notificationColumns, n.Text, n.Source, n.Priority, n.Unread, n.CreatedAt.UnixNano(), n.QueuedAt.UnixNano()))
}
func (r *Notifications) Pending(ctx context.Context) ([]domain.Notification, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT "+notificationColumns+" FROM notifications WHERE dismissed=0 ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []domain.Notification{}
	for rows.Next() {
		n, err := scanNotification(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, rows.Err()
}
func (r *Notifications) Dismiss(ctx context.Context, id int64) (domain.Notification, error) {
	return scanNotification(r.db.QueryRowContext(ctx, "UPDATE notifications SET dismissed=1,unread=0 WHERE id=? AND dismissed=0 RETURNING "+notificationColumns, id))
}
func (r *Notifications) Requeue(ctx context.Context, id int64, now time.Time) (domain.Notification, error) {
	return scanNotification(r.db.QueryRowContext(ctx, "UPDATE notifications SET priority=MAX(priority-1,0),unread=1,queued_at=? WHERE id=? AND dismissed=0 RETURNING "+notificationColumns, now.UnixNano(), id))
}
