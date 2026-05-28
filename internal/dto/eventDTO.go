package dto

import "time"

type EventInput struct {
	EventId     string    `json:"event_id"`
	CardId      string    `json:"card_id"`
	ClientEmail string    `json:"client_email"`
	Timestamp   time.Time `json:"timestamp"`
}

type EventResponse struct {
	EventId     string     `json:"event_id"`
	CardId      string     `json:"card_id"`
	ClientEmail string     `json:"client_email"`
	PipefyCard  string     `json:"pipefy_card"`
	Timestamp   *time.Time `json:"timestamp"`
}
