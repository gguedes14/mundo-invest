package client

import (
	"context"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/repository"
)

func CreateClient(ctx context.Context, client *domain.Client, r *repository.Repository) error {
	result := r.Db.WithContext(ctx).Create(client)

	return result.Error
}
