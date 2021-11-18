package mysql

import (
	"context"

	"github.com/danic95/things-service/internal/things/entity"
)

// GetThing stands for getting all things from db
func (m *Mysql) GetThings(ctx context.Context) ([]entity.Thing, error) {
	things := []entity.Thing{}

	//TODO: Implement mysql calls
	err := m.db.SelectContext(ctx, &things, "SELECT id, name, email, phone, birth_date, start_day FROM things")

	return things, err
}

// GetThing stands for getting a thing from db
func (m *Mysql) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	thing := &entity.Thing{}

	//TODO: Implement mysql calls
	err := m.db.GetContext(ctx, thing, "SELECT id, name, email, phone, birth_date, start_day FROM things WHERE id = ?", id)

	return thing, err
}

func (m *Mysql) SaveThing(ctx context.Context, thing *entity.Thing) error {

	//TODO: Implement mysql calls
	const query string = "INSERT INTO things (name, email, phone, birth_date, start_day) VALUES (?,?,?,?,?)"
	_, err := m.db.ExecContext(ctx, query, thing.Name, thing.Email, thing.Phone, thing.Birth_date.Format("2006-01-02"),
		thing.Start_day)

	return err
}
