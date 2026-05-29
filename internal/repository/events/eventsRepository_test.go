package events

import (
	"context"
	"testing"
	"time"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/repository"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateEvent(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Events{})

	repo := repository.NewClientRepository(db)

	event := &domain.Events{
		ID:           uuid.New(),
		EventId:      "event123",
		CardId:       "card123",
		ClienteEmail: "email@email.com",
		TimeStamp:    time.Now(),
	}
	err = repo.Db.WithContext(context.Background()).Create(event).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if event.ID == uuid.Nil {
		t.Fatal("ID não foi preenchido")
	}
}

func TestFindEventByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&domain.Events{})

	repo := repository.NewClientRepository(db)

	event := &domain.Events{
		ID:           uuid.New(),
		EventId:      "event123",
		CardId:       "card123",
		ClienteEmail: "email@email.com",
		TimeStamp:    time.Now(),
	}
	err = repo.Db.WithContext(context.Background()).Create(event).Error

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	var foundEvent domain.Events
	result := repo.Db.WithContext(context.Background()).Where("id = ?", event.ID).First(&foundEvent)

	if result.Error != nil {
		t.Fatalf("erro inesperado ao buscar evento: %v", result.Error)
	}

	if foundEvent.ID != event.ID {
		t.Fatal("evento encontrado tem ID diferente do esperado")
	}
}
