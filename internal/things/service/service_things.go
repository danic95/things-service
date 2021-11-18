package service

import (
	"context"
	"errors"

	"github.com/danic95/things-service/internal/things/entity"
)

var ErrInvalidThing error = errors.New("invalid thing")
var ErrThingNotFound error = errors.New("thing not found")

// GetThings gives you all things from storage
func (s *Service) GetThings(ctx context.Context) ([]entity.Thing, error) {

	//TODO: Implement business logic

	return s.repo.GetThings(ctx)
}

// GetThing gives you a thing from storage
func (s *Service) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {

	//TODO: Implement business logic
	thing_got, err := s.repo.GetThing(ctx, id)
	if err != nil {
		err = ErrThingNotFound
	}

	return thing_got, err
}

// CreateThing creates a new thing in storage
func (s *Service) CreateThing(ctx context.Context, thing *entity.Thing) error {

	if thing == nil || len(thing.Name) == 0 || len(thing.Birth_date.String()) == 0 || thing.Start_day == 0 {
		return ErrInvalidThing
	}

	return s.repo.SaveThing(ctx, thing)
}
