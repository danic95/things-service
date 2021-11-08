package mysql

import (
	"context"
	"github.com/danic95/things-service/internal/things/entity"
)

// GetThing stands for getting a thing from db
func (m *Mysql) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	thing := &entity.Thing{}

	//TODO: Implement mysql calls
	return thing, nil
}

func (m *Mysql) SaveThing(ctx context.Context, thing *entity.Thing) error {

	//TODO: Implement mysql calls
	return nil
}