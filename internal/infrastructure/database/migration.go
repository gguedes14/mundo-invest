package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	ClienteNome     string    `gorm:"not null"`
	ClienteEmail    string    `gorm:"uniqueIndex;not null"`
	TipoSolicitacao string    `gorm:"not null"`
	ValorPatrimonio float64   `gorm:"not null"`
	Status          string    `gorm:"not null"`
	Prioridade      *string
	PipefyCardID    *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Events struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventId      string    `gorm:"uniqueIndex;not null"`
	CardId       string    `gorm:"not null"`
	ClienteEmail string    `gorm:"not null"`
	TimeStamp    time.Time `gorm:"not null"`
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&Client{},
		&Events{},
	)
}
