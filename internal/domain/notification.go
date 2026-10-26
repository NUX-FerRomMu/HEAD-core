package domain

import (
	"errors"
	"time"
)

var ErrNotificationNotFound = errors.New("pending notification not found")

type Notification struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	Source    string    `json:"source"`
	Priority  int       `json:"priority"`
	Unread    bool      `json:"unread"`
	Dismissed bool      `json:"dismissed"`
	CreatedAt time.Time `json:"created_at"`
	QueuedAt  time.Time `json:"queued_at"`
}

func (n Notification) EffectivePriority(now time.Time) int {
	age := now.Sub(n.QueuedAt) / time.Hour
	if age < 0 {
		age = 0
	}
	return n.Priority + int(age)
}
