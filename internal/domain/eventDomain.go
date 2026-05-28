package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Events struct {
	ID          uuid.UUID `json:"id"`
	EventId     string    `json:"event_id"`
	CardId      string    `json:"card_id"`
	ClientEmail string    `json:"cliente_email"`
	TimeStamp   time.Time `json:"processed_at"`
}

func (Events) TableName() string {
	return "events"
}

type EventDomain interface {
	CreateEvent(ctx context.Context, event *Events) error
}
