package event

import (
	"context"

	"github.com/gguedes14/mundo-invest/internal/controller"
	"github.com/gguedes14/mundo-invest/internal/dto"
	serviceEvent "github.com/gguedes14/mundo-invest/internal/service/events"
)

func CreateEvent(ctx context.Context, c *controller.Controller, input dto.EventInput) (*dto.EventResponse, error) {
	event, pipefyCard, err := serviceEvent.CreateEvent(ctx, c.Service, input)

	if err != nil {
		return nil, err
	}

	return &dto.EventResponse{
		EventId:      event.EventId,
		CardId:       event.CardId,
		ClienteEmail: event.ClienteEmail,
		PipefyCard:   pipefyCard,
		Timestamp:    &event.TimeStamp,
	}, nil
}
