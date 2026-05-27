package model

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nome            string    `gorm:"not null"`
	Email           string    `gorm:"uniqueIndex;not null"`
	TipoSolicitacao string    `gorm:"not null"`
	ValorPatrimonio float64   `gorm:"not null"`
	Status          string    `gorm:"not null"`
	Prioridade      *string
	PipefyCardID    *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
