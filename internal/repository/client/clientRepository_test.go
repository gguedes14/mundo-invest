package client

import (
	"context"
	"testing"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/repository"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRepository_CreateClient(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Client{})

	repo := repository.NewClientRepository(db)

	client := &domain.Client{
		ID:           uuid.New(),
		ClienteNome:  "Gabriel",
		ClienteEmail: "gabriel@email.com",
	}
	err = repo.Db.WithContext(context.Background()).Create(client).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if client.ID == uuid.Nil {
		t.Fatal("ID não foi preenchido")
	}
}

func TestRepository_FindClientByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Client{})

	repo := repository.NewClientRepository(db)

	client := &domain.Client{
		ID:           uuid.New(),
		ClienteNome:  "Gabriel",
		ClienteEmail: "email@email.com",
	}
	err = repo.Db.WithContext(context.Background()).Create(client).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	foundClient, err := FindClientByID(context.Background(), client.ID, repo)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if foundClient.ID != client.ID {
		t.Fatalf("IDs não coincidem: esperado %v, obtido %v", client.ID, foundClient.ID)
	}
}

func TestRepository_FindClientByID_NotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Client{})

	repo := repository.NewClientRepository(db)

	_, err = FindClientByID(context.Background(), uuid.New(), repo)

	if err == nil {
		t.Fatal("esperava um erro, mas obteve nil")
	}
}

func TestRepository_UpdateClientStatus(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Client{})

	repo := repository.NewClientRepository(db)

	client := &domain.Client{
		ID:           uuid.New(),
		ClienteNome:  "Gabriel",
		ClienteEmail: "email@email.com",
		Status:       "ativo",
	}
	err = repo.Db.WithContext(context.Background()).Create(client).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	err = UpdateClientStatus(context.Background(), client.ID, "inativo", repo)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	var updatedClient domain.Client
	err = repo.Db.WithContext(context.Background()).Where("id = ?", client.ID).First(&updatedClient).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if updatedClient.Status != "inativo" {
		t.Fatalf("status não foi atualizado: esperado 'inativo', obtido '%s'", updatedClient.Status)
	}
}

func TestRepository_UpdateClientPriority(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Client{})

	repo := repository.NewClientRepository(db)

	client := &domain.Client{
		ID:           uuid.New(),
		ClienteNome:  "Gabriel",
		ClienteEmail: "email@email.com",
		Prioridade:   "baixa",
	}
	err = repo.Db.WithContext(context.Background()).Create(client).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	err = UpdateClientPriority(context.Background(), client.ID, "alta", repo)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	var updatedClient domain.Client
	err = repo.Db.WithContext(context.Background()).Where("id = ?", client.ID).First(&updatedClient).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if updatedClient.Prioridade != "alta" {
		t.Fatalf("prioridade não foi atualizada: esperado 'alta', obtido '%s'", updatedClient.Prioridade)
	}
}
