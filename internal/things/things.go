package things

import (
	"context"

	"github.com/danic95/things-service/internal/things/entity"
)

// Service declares and summarizes the functionality a
// service in the containing package will implement
type Service interface {
	GetThings(ctx context.Context) (*entity.Things, error)
	GetThing(ctx context.Context, id uint) (*entity.Thing, error)
	CreateThing(ctx context.Context, thing *entity.Thing) error
}

// Repository declares and summarizes the functionality a
// repository in the containing package will implement
type Repository interface {
	GetThings(ctx context.Context) (*entity.Things, error)
	GetThing(ctx context.Context, id uint) (*entity.Thing, error)
	SaveThing(ctx context.Context, thing *entity.Thing) error
}
