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
