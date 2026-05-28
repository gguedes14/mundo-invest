package client

import (
	"context"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/repository"
	"github.com/google/uuid"
)

func CreateClient(ctx context.Context, client *domain.Client, r *repository.Repository) error {
	result := r.Db.WithContext(ctx).Create(client)

	return result.Error
}

func FindClientByID(ctx context.Context, id uuid.UUID, r *repository.Repository) (*domain.Client, error) {
	var client domain.Client
	result := r.Db.WithContext(ctx).Where("id = ?", id).First(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func UpdateClientStatus(ctx context.Context, id uuid.UUID, status string, r *repository.Repository) error {
	result := r.Db.WithContext(ctx).Model(&domain.Client{}).Where("id = ?", id).Update("status", status)

	return result.Error
}

func UpdateClientPriority(ctx context.Context, id uuid.UUID, status string, r *repository.Repository) error {
	result := r.Db.WithContext(ctx).Model(&domain.Client{}).Where("id = ?", id).Update("prioridade", status)

	return result.Error
}
