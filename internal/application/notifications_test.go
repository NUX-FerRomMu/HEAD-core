package application_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/application"
	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
	"github.com/NUX-FerRomMu/HEAD-core/internal/store"
)

func TestNotificationQueue(t *testing.T) {
	ctx := context.Background()
	path := filepath.Join(t.TempDir(), "state.sqlite")
	db, err := store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	repo, err := store.NewNotifications(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	service := application.Notifications{Repo: repo, Now: func() time.Time { return now }}
	if next, err := service.Next(ctx); err != nil || next != nil {
		t.Fatalf("empty: %+v %v", next, err)
	}
	old, err := service.Add(ctx, "old", "manual", 0)
	if err != nil {
		t.Fatal(err)
	}
	now = now.Add(5 * time.Hour)
	recent, err := service.Add(ctx, "recent", "agent", 3)
	if err != nil {
		t.Fatal(err)
	}
	next, err := service.Next(ctx)
	if err != nil || next.ID != old.ID {
		t.Fatalf("aging: %+v %v", next, err)
	}
	if _, err = service.Requeue(ctx, old.ID); err != nil {
		t.Fatal(err)
	}
	next, err = service.Next(ctx)
	if err != nil || next.ID != recent.ID {
		t.Fatalf("requeue: %+v %v", next, err)
	}
	changed, err := service.Requeue(ctx, recent.ID)
	if err != nil || changed.Priority != 2 || !changed.Unread {
		t.Fatalf("%+v %v", changed, err)
	}
	if _, err = repo.Dismiss(ctx, recent.ID); err != nil {
		t.Fatal(err)
	}
	if _, err = service.Requeue(ctx, recent.ID); !errors.Is(err, domain.ErrNotificationNotFound) {
		t.Fatal(err)
	}
	if _, err = repo.Dismiss(ctx, 99999); !errors.Is(err, domain.ErrNotificationNotFound) {
		t.Fatal(err)
	}
	for _, priority := range []int{-1, 4} {
		if _, err = service.Add(ctx, "invalid", "", priority); err == nil {
			t.Fatal("accepted priority")
		}
	}
	if _, err = service.Add(ctx, "  ", "", 1); err == nil {
		t.Fatal("accepted blank")
	}
	db.Close()
	db, err = store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repo, err = store.NewNotifications(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	service.Repo = repo
	pending, err := service.List(ctx)
	if err != nil || len(pending) != 1 || pending[0].ID != old.ID || pending[0].Priority != 0 {
		t.Fatalf("reopen: %+v %v", pending, err)
	}
	// Equal timestamps/priorities fall back to insertion order.
	newer, err := service.Add(ctx, "tie", "", 0)
	if err != nil {
		t.Fatal(err)
	}
	pending, err = service.List(ctx)
	if err != nil || pending[0].ID != old.ID || pending[1].ID != newer.ID {
		t.Fatalf("tie: %+v %v", pending, err)
	}
}
