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

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&Client{},
	)
}
