package events

import (
	"context"

	"github.com/gguedes14/mundo-invest/internal/domain"
	"github.com/gguedes14/mundo-invest/internal/repository"
)

func CreateEvent(ctx context.Context, event *domain.Events, r *repository.Repository) error {
	result := r.Db.WithContext(ctx).Create(event)

	return result.Error
}

func FindEventByID(ctx context.Context, id string, r *repository.Repository) (*domain.Events, error) {
	var event domain.Events
	result := r.Db.WithContext(ctx).Where("id = ?", id).First(&event)

	if result.Error != nil {
		return nil, result.Error
	}

	return &event, nil
}
