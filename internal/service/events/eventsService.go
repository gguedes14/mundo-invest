package events

import (
	"context"
	"errors"
	"fmt"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/integrations/pipefy"
	repositoryClients "github.com/gguedes14/mundo-invest/internal/repository/client"
	repositoryevents "github.com/gguedes14/mundo-invest/internal/repository/events"
	"github.com/gguedes14/mundo-invest/internal/service"
)

func FindEventByID(ctx context.Context, id string, s *service.Service) (*domain.Events, error) {
	event, err := repositoryevents.FindEventByID(ctx, id, s.Repo)

	if err != nil {
		return nil, fmt.Errorf("error finding event: %w", err)
	}

	return event, nil
}

func CreateEvent(ctx context.Context, s *service.Service, input dto.EventInput, clientId dto.ClientResponse) (*domain.Events, string, error) {
	event := &domain.Events{
		EventId:     input.EventId,
		CardId:      input.CardId,
		ClientEmail: input.ClientEmail,
		TimeStamp:   input.Timestamp,
	}

	FindEvents, err := repositoryevents.FindEventByID(
		ctx,
		event.EventId,
		s.Repo,
	)

	if err == nil && FindEvents != nil {
		return nil, "", errors.New("event already exists")
	}

	client, err := repositoryClients.FindClientByID(
		ctx,
		clientId.ID,
		s.Repo,
	)

	if err != nil {
		return nil, "", fmt.Errorf("error finding client: %w", err)
	}

	prioridade := "prioridade_normal"

	if client.ValorPatrimonio >= 200000 {
		prioridade = "prioridade_alta"
	}

	client.Prioridade = prioridade
	client.Status = "Processado"

	if err := repositoryevents.CreateEvent(
		ctx,
		event,
		s.Repo,
	); err != nil {
		return nil, "", fmt.Errorf("error creating event: %w", err)
	}

	if err := repositoryClients.UpdateClientStatus(
		ctx,
		client.ID,
		client.Status,
		s.Repo,
	); err != nil {
		return nil, "", fmt.Errorf("error updating client status: %w", err)
	}

	if err := repositoryClients.UpdateClientPriority(
		ctx,
		client.ID,
		client.Prioridade,
		s.Repo,
	); err != nil {
		return nil, "", fmt.Errorf("error updating client priority: %w", err)
	}

	pipefyMutation := updatePipefyFieldCard(event)

	return event, pipefyMutation, nil
}

func updatePipefyFieldCard(event *domain.Events) string {
	return fmt.Sprintf(
		pipefy.UpdateCardFieldMutation,
		event.CardId,
		event.ClientEmail,
		event.TimeStamp,
	)
}
