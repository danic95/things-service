package sqlite

import (
	"context"

	"github.com/danic95/things-service/internal/things/entity"
)

func (s *Sqlite) GetThings(ctx context.Context) (*entity.Things, error) {
	things := &entity.Things{}
	err := s.db.SelectContext(ctx, &things.Body, "SELECT * FROM THINGS")

	return things, err
}

func (s *Sqlite) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	thing := &entity.Thing{}
	err := s.db.GetContext(ctx, thing, "SELECT * FROM THINGS WHERE id = ?", id)

	return thing, err
}

func (s *Sqlite) SaveThing(ctx context.Context, thing *entity.Thing) error {
	const query string = "INSERT INTO THINGS (name) VALUES (?)"
	_, err := s.db.ExecContext(ctx, query, thing.Name)

	return err
}
