package events

import (
	"context"
	"testing"
	"time"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/dto"
	"github.com/gguedes14/mundo-invest/internal/repository"
	repositoryevents "github.com/gguedes14/mundo-invest/internal/repository/events"
	"github.com/gguedes14/mundo-invest/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *repository.Repository {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("erro ao abrir banco em memória: %v", err)
	}

	err = db.AutoMigrate(
		&domain.Events{},
		&domain.Client{},
	)
	if err != nil {
		t.Fatalf("erro ao migrar schema: %v", err)
	}

	return &repository.Repository{Db: db}
}

func TestFindEventByID_Success(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	event := &domain.Events{
		ID:           uuid.New(),
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		TimeStamp:    time.Now(),
	}

	err := repositoryevents.CreateEvent(context.Background(), event, s.Repo)
	assert.NoError(t, err)

	foundEvent, err := FindEventByID(context.Background(), "event_123", s)

	assert.NoError(t, err)
	assert.NotNil(t, foundEvent)
	assert.Equal(t, "event_123", foundEvent.EventId)
}

func TestFindEventByID_NotFound(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	foundEvent, err := FindEventByID(context.Background(), "nonexistent_event", s)

	assert.Error(t, err)
	assert.Nil(t, foundEvent)
}

func TestCreateEvent_Success(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	client := &domain.Client{
		ID:              uuid.New(),
		ClienteEmail:    "email@email.com",
		ValorPatrimonio: 250000,
	}

	err := s.Repo.Db.Create(client).Error
	assert.NoError(t, err)

	input := dto.EventInput{
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		Timestamp:    time.Now(),
	}

	event, pipefy, err := CreateEvent(context.Background(), s, input)

	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "event_123", event.EventId)
	assert.NotEmpty(t, pipefy)
}

func TestCreateEvent_EventAlreadyExists(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	event := &domain.Events{
		ID:           uuid.New(),
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		TimeStamp:    time.Now(),
	}

	err := repositoryevents.CreateEvent(context.Background(), event, s.Repo)
	assert.NoError(t, err)

	input := dto.EventInput{
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		Timestamp:    time.Now(),
	}

	createdEvent, pipefy, err := CreateEvent(context.Background(), s, input)

	assert.Error(t, err)
	assert.Nil(t, createdEvent)
	assert.Empty(t, pipefy)
}

func TestCreateEvent_ClientNotFound(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	input := dto.EventInput{
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		Timestamp:    time.Now(),
	}

	event, pipefy, err := CreateEvent(context.Background(), s, input)

	assert.Error(t, err)
	assert.Nil(t, event)
	assert.Empty(t, pipefy)
}

func TestCreateEvent_PrioridadeAlta(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	client := &domain.Client{
		ID:              uuid.New(),
		ClienteEmail:    "email@email.com",
		ValorPatrimonio: 250000,
	}

	err := s.Repo.Db.Create(client).Error
	assert.NoError(t, err)

	input := dto.EventInput{
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		Timestamp:    time.Now(),
	}

	event, pipefy, err := CreateEvent(context.Background(), s, input)

	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "event_123", event.EventId)

	assert.Contains(t, pipefy, "mutation")

	var updatedClient domain.Client
	err = s.Repo.Db.Where("cliente_email = ?", "email@email.com").First(&updatedClient).Error
	assert.NoError(t, err)

	assert.Equal(t, "prioridade_alta", updatedClient.Prioridade)
}

func TestCreateEvent_PrioridadeNormal(t *testing.T) {
	s := &service.Service{
		Repo: setupTestDB(t),
	}

	client := &domain.Client{
		ID:              uuid.New(),
		ClienteEmail:    "email@email.com",
		ValorPatrimonio: 150000,
	}

	err := s.Repo.Db.Create(client).Error
	assert.NoError(t, err)

	input := dto.EventInput{
		EventId:      "event_123",
		CardId:       "card_123",
		ClienteEmail: "email@email.com",
		Timestamp:    time.Now(),
	}

	event, pipefy, err := CreateEvent(context.Background(), s, input)

	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "event_123", event.EventId)

	assert.Contains(t, pipefy, "mutation")

	var updatedClient domain.Client
	err = s.Repo.Db.Where("cliente_email = ?", "email@email.com").First(&updatedClient).Error
	assert.NoError(t, err)

	assert.Equal(t, "prioridade_normal", updatedClient.Prioridade)
}
