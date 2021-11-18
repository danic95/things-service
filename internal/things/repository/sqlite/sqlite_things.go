package sqlite

import (
	"context"

	"github.com/danic95/things-service/internal/things/entity"
)

func (s *Sqlite) GetThings(ctx context.Context) ([]entity.Thing, error) {
	things := []entity.Thing{}
	err := s.db.SelectContext(ctx, &things, "SELECT id, name, email, phone, birth_date, start_day FROM things")

	return things, err
}

func (s *Sqlite) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	thing := &entity.Thing{}
	err := s.db.GetContext(ctx, thing, "SELECT id, name, email, phone, birth_date, start_day FROM things WHERE id = ?", id)

	return thing, err
}

func (s *Sqlite) SaveThing(ctx context.Context, thing *entity.Thing) error {
	const query string = "INSERT INTO things (name, email, phone, birth_date, start_day) VALUES (?,?,?,?,?)"
	_, err := s.db.ExecContext(ctx, query, thing.Name, thing.Email, thing.Phone, thing.Birth_date.Format("2006-01-02"),
		thing.Start_day)

	return err
}
