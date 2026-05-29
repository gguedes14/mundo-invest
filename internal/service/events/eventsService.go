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
	"github.com/google/uuid"
)

const (
	PrioridadeAlta   = "prioridade_alta"
	PrioridadeNormal = "prioridade_normal"
)

func FindEventByID(ctx context.Context, id string, s *service.Service) (*domain.Events, error) {
	event, err := repositoryevents.FindEventByID(ctx, id, s.Repo)

	if err != nil {
		return nil, fmt.Errorf("error finding event: %w", err)
	}

	return event, nil
}

func CreateEvent(ctx context.Context, s *service.Service, input dto.EventInput) (*domain.Events, string, error) {
	event := &domain.Events{
		ID:           uuid.New(),
		EventId:      input.EventId,
		CardId:       input.CardId,
		ClienteEmail: input.ClienteEmail,
		TimeStamp:    input.Timestamp,
	}

	if event.EventId == "" || event.CardId == "" || event.ClienteEmail == "" {
		return nil, "", errors.New("missing required fields")
	}

	FindEvents, err := repositoryevents.FindEventByID(
		ctx,
		event.EventId,
		s.Repo,
	)

	if err == nil && FindEvents != nil {
		return nil, "", errors.New("event already exists")
	}

	client, err := repositoryClients.FindClientByEmail(
		ctx,
		input.ClienteEmail,
		s.Repo,
	)

	if err != nil {
		return nil, "", fmt.Errorf("error finding client: %w", err)
	}

	prioridade := PrioridadeNormal

	if client.ValorPatrimonio >= 200000 {
		prioridade = PrioridadeAlta
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

	pipefyMutation := updatePipefyFieldCard(event, prioridade)

	return event, pipefyMutation, nil
}

func updatePipefyFieldCard(event *domain.Events, prioridade string) string {
	return fmt.Sprintf(
		pipefy.UpdateCardFieldMutation,
		event.CardId,
		"prioridade",
		prioridade,
	)
}
