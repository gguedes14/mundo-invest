package client

import (
	"context"

	"github.com/gguedes14/mundo-invest/internal/controller"
	"github.com/gguedes14/mundo-invest/internal/dto"
	serviceclient "github.com/gguedes14/mundo-invest/internal/service/client"
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

func FindClientByEmail(ctx context.Context, c *controller.Controller, email string) (*dto.ClientResponse, error) {
	client, err := serviceclient.FindClientByEmail(ctx, c.Service, email)

	if err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		ID:      client.ID,
		Status:  client.Status,
		Message: "Cliente encontrado com sucesso",
	}, nil
}
