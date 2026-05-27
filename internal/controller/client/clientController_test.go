package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	serviceclient "github.com/gguedes14/mundo-invest/internal/service/client"

	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/repository"
	"github.com/gguedes14/mundo-invest/internal/service"
)

func TestCreateClient_InvalidEmail(t *testing.T) {
	s := &service.Service{
		Repo: &repository.Repository{},
	}

	input := dto.ClientInput{
		ClienteNome:  "Gabriel",
		ClienteEmail: "email-invalido",
	}

	client, pipefy, err := serviceclient.CreateClient(
		context.Background(),
		s,
		input,
	)

	require.Error(t, err)
	require.Nil(t, client)
	require.Empty(t, pipefy)
	require.Equal(t, "invalid email format", err.Error())
}
