package application

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

type NotificationRepository interface {
	Add(context.Context, domain.Notification) (domain.Notification, error)
	Pending(context.Context) ([]domain.Notification, error)
	Dismiss(context.Context, int64) (domain.Notification, error)
	Requeue(context.Context, int64, time.Time) (domain.Notification, error)
}
type Notifications struct {
	Repo NotificationRepository
	Now  func() time.Time
}

func (s Notifications) now() time.Time {
	if s.Now != nil {
		return s.Now().UTC()
	}
	return time.Now().UTC()
}
func (s Notifications) Add(ctx context.Context, text, source string, priority int) (domain.Notification, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return domain.Notification{}, fmt.Errorf("notification text must not be empty")
	}
	if priority < 0 || priority > 3 {
		return domain.Notification{}, fmt.Errorf("priority must be between 0 and 3")
	}
	now := s.now()
	return s.Repo.Add(ctx, domain.Notification{Text: text, Source: source, Priority: priority, Unread: true, CreatedAt: now, QueuedAt: now})
}
func (s Notifications) List(ctx context.Context) ([]domain.Notification, error) {
	items, err := s.Repo.Pending(ctx)
	if err != nil {
		return nil, err
	}
	now := s.now()
	sort.Slice(items, func(i, j int) bool {
		a, b := items[i], items[j]
		pa, pb := a.EffectivePriority(now), b.EffectivePriority(now)
		if pa != pb {
			return pa > pb
		}
		if !a.QueuedAt.Equal(b.QueuedAt) {
			return a.QueuedAt.Before(b.QueuedAt)
		}
		return a.ID < b.ID
	})
	return items, nil
}
func (s Notifications) Next(ctx context.Context) (*domain.Notification, error) {
	items, err := s.List(ctx)
	if err != nil || len(items) == 0 {
		return nil, err
	}
	return &items[0], nil
}
func (s Notifications) Requeue(ctx context.Context, id int64) (domain.Notification, error) {
	return s.Repo.Requeue(ctx, id, s.now())
}
