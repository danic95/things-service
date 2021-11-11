package mysql

import (
	"context"

	"github.com/danic95/things-service/internal/things/entity"
)

// GetThing stands for getting all things from db
func (m *Mysql) GetThings(ctx context.Context) (*entity.Things, error) {
	things := &entity.Things{}

	//TODO: Implement mysql calls
	err := m.db.SelectContext(ctx, &things.Body, "SELECT * FROM THINGS")

	return things, err
}

// GetThing stands for getting a thing from db
func (m *Mysql) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	thing := &entity.Thing{}

	//TODO: Implement mysql calls
	err := m.db.GetContext(ctx, thing, "SELECT * FROM THINGS WHERE id = ?", id)

	return thing, err
}

func (m *Mysql) SaveThing(ctx context.Context, thing *entity.Thing) error {

	//TODO: Implement mysql calls
	const query string = "INSERT INTO THINGS (name) VALUES (?)"
	_, err := m.db.ExecContext(ctx, query, thing.Name)

	return err
}
