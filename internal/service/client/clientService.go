package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/integrations/pipefy"
	repositoryclient "github.com/gguedes14/mundo-invest/internal/repository/client"
	"github.com/gguedes14/mundo-invest/internal/service"
	"github.com/google/uuid"
)

func CreateClient(ctx context.Context, s *service.Service, input dto.ClientInput) (*domain.Client, string, error) {
	if strings.TrimSpace(input.ClienteNome) == "" {
		return nil, "", errors.New("cliente_nome is required")
	}

	if strings.TrimSpace(input.ClienteEmail) == "" {
		return nil, "", errors.New("cliente_email is required")
	}

	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(regex, input.ClienteEmail)

	if err != nil {
		return nil, "", fmt.Errorf("error validating email: %w", err)
	}

	if !matched {
		return nil, "", errors.New("invalid email format")
	}

	existingClient, err := repositoryclient.FindClientByEmail(ctx, input.ClienteEmail, s.Repo)

	if err == nil && existingClient != nil {
		return nil, "", errors.New("client already exists")
	}

	client := &domain.Client{
		ID:              uuid.New(),
		ClienteNome:     input.ClienteNome,
		ClienteEmail:    input.ClienteEmail,
		TipoSolicitacao: input.TipoSolicitacao,
		ValorPatrimonio: input.ValorPatrimonio,
		Status:          "Aguardando Análise",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	pipefyCard := CreatePipefyCard(client)

	if err := repositoryclient.CreateClient(ctx, client, s.Repo); err != nil {
		return nil, "", fmt.Errorf("error creating client: %w", err)
	}

	return client, pipefyCard, nil
}

func CreatePipefyCard(client *domain.Client) string {
	return fmt.Sprintf(
		pipefy.CreateCardMutation,
		client.ClienteNome,
		client.ClienteNome,
		client.ClienteEmail,
		client.ValorPatrimonio,
	)
}

func FindClientByEmail(ctx context.Context, s *service.Service, email string) (*domain.Client, error) {
	client, err := repositoryclient.FindClientByEmail(ctx, email, s.Repo)
	if err != nil {
		return nil, fmt.Errorf("error finding client by email: %w", err)
	}

	return client, nil
}
