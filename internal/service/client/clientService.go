package service

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/dto"
	repositoryclient "github.com/gguedes14/mundo-invest/internal/repository/client"
	"github.com/gguedes14/mundo-invest/internal/service"
	"github.com/google/uuid"
)

func CreateClient(ctx context.Context, s *service.Service, input dto.ClientInput) (*domain.Client, error) {
	if strings.TrimSpace(input.ClienteNome) == "" {
		return nil, errors.New("cliente_nome is required")
	}

	if strings.TrimSpace(input.ClienteEmail) == "" {
		return nil, errors.New("cliente_email is required")
	}

	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(regex, input.ClienteEmail)

	if err != nil {
		return nil, errors.New("error validating email")
	}

	if !matched {
		return nil, errors.New("invalid email format")
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

	if err := repositoryclient.CreateClient(ctx, client, s.Repo); err != nil {
		return nil, err
	}

	return client, nil
}
