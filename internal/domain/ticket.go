package domain

import (
	"errors"
	"time"
)

var ErrTicketNotFound = errors.New("ticket not found")
var TicketStates = []string{"todo", "in_progress", "review", "done"}

func ValidTicketState(state string) bool {
	for _, candidate := range TicketStates {
		if state == candidate {
			return true
		}
	}
	return false
}

type Ticket struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Criteria    []string  `json:"criteria"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TicketChanges struct {
	Title       *string
	Description *string
	Criteria    *[]string
}
