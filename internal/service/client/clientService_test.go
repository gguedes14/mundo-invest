package service

import (
	"context"
	"testing"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/repository"
	"github.com/gguedes14/mundo-invest/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *repository.Repository {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("erro ao abrir banco em memória: %v", err)
	}

	if err := db.AutoMigrate(&domain.Client{}); err != nil {
		t.Fatalf("erro ao migrar schema: %v", err)
	}

	return &repository.Repository{Db: db}
}

func TestCreateClient_Success(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	input := dto.ClientInput{
		ClienteNome:     "Gabriel",
		ClienteEmail:    "gabriel@email.com",
		TipoSolicitacao: "investimento",
		ValorPatrimonio: 100000,
	}

	client, card, err := CreateClient(context.Background(), s, input)

	if err != nil {
		t.Fatalf("não esperava erro: %v", err)
	}
	if client == nil {
		t.Fatal("client não pode ser nil")
	}
	if card == "" {
		t.Fatal("card não pode ser vazio")
	}
}

func TestCreateClient_InvalidEmail(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	input := dto.ClientInput{
		ClienteNome:     "Gabriel",
		ClienteEmail:    "email-invalido",
		TipoSolicitacao: "investimento",
		ValorPatrimonio: 100000,
	}

	client, card, err := CreateClient(context.Background(), s, input)

	if err == nil {
		t.Fatal("esperava erro para email inválido")
	}
	if client != nil {
		t.Fatal("client deve ser nil para email inválido")
	}
	if card != "" {
		t.Fatal("card deve ser vazio para email inválido")
	}
}

func TestCreateClient_EmailRequired(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	input := dto.ClientInput{
		ClienteNome:     "Gabriel",
		TipoSolicitacao: "investimento",
		ValorPatrimonio: 100000,
	}

	client, card, err := CreateClient(context.Background(), s, input)

	if err == nil {
		t.Fatal("esperava erro para email ausente")
	}
	if client != nil {
		t.Fatal("client deve ser nil para email ausente")
	}
	if card != "" {
		t.Fatal("card deve ser vazio para email ausente")
	}
}

func TestCreateClient_NomeRequired(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	input := dto.ClientInput{
		ClienteEmail:    "email@email.com",
		TipoSolicitacao: "investimento",
		ValorPatrimonio: 100000,
	}

	client, card, err := CreateClient(context.Background(), s, input)

	if err == nil {
		t.Fatal("esperava erro para nome ausente")
	}
	if client != nil {
		t.Fatal("client deve ser nil para nome ausente")
	}
	if card != "" {
		t.Fatal("card deve ser vazio para nome ausente")
	}
}
