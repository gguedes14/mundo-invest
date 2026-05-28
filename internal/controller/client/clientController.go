package client

import (
	"context"

	"github.com/gguedes14/mundo-invest/internal/controller"
	"github.com/gguedes14/mundo-invest/internal/dto"
	serviceclient "github.com/gguedes14/mundo-invest/internal/service/client"
	"github.com/google/uuid"
)

func CreateClient(ctx context.Context, c *controller.Controller, input dto.ClientInput) (*dto.ClientResponse, error) {
	client, pipefyCard, err := serviceclient.CreateClient(ctx, c.Service, input)

	if err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		Status:     client.Status,
		Message:    "Solicitação de investimento criada com sucesso",
		PipefyCard: pipefyCard,
	}, nil
}

func FindClientByID(ctx context.Context, c *controller.Controller, id uuid.UUID) (*dto.ClientResponse, error) {
	client, err := serviceclient.FindClientByID(ctx, id, c.Service)

	if err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		ID:         client.ID,
		Status:     client.Status,
		Message:    "Cliente encontrado com sucesso",
		PipefyCard: "",
	}, nil
}
