package service

import (
	"context"

	"github.com/danic95/things-service/internal/things/entity"
)

// GetThings gives you all things from storage
func (s *Service) GetThings(ctx context.Context) (*entity.Things, error) {

	//TODO: Implement business logic

	return s.repo.GetThings(ctx)
}

// GetThing gives you a thing from storage
func (s *Service) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {

	//TODO: Implement business logic

	return s.repo.GetThing(ctx, id)
}

// CreateThing creates a new thing in storage
func (s *Service) CreateThing(ctx context.Context, thing *entity.Thing) error {

	//TODO: Implement business logic
	return s.repo.SaveThing(ctx, thing)
}
