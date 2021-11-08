package service

import (
	"context"
	"github.com/danic95/things-service/internal/things/entity"
)

// GetThing gives you a thing from storage
func (s *Service) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {

	//TODO: Implement business logic

	return s.repo.GetThing(ctx, id)
}

func (s *Service) CreateThing(ctx context.Context, thing *entity.Thing) error {

	//TODO: Implement business logic
	return s.repo.SaveThing(ctx, thing)
}
