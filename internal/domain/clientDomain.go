package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID              uuid.UUID `json:"id"`
	ClienteNome     string    `json:"cliente_nome"`
	ClienteEmail    string    `json:"cliente_email"`
	TipoSolicitacao string    `json:"tipo_solicitacao"`
	ValorPatrimonio float64   `json:"valor_patrimonio"`
	Status          string    `json:"status"`
	Prioridade      *string   `json:"prioridade"`
	PipefyCardID    *string   `json:"pipefyCardId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (Client) TableName() string {
	return "clientes"
}

type ClientDomain interface {
	CreateClient(ctx context.Context, client *Client) error
}
