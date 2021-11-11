package sqlite

import (
	"context"
	_ "embed"

	"github.com/danic95/things-service/internal/things/repository/cache"
	"github.com/jmoiron/sqlx"
)

type Sqlite struct {
	cache *cache.Cache
	db    *sqlx.DB
}

//go:embed schema.sql
var createSchemaStmt string

func New(db *sqlx.DB, cach *cache.Cache) *Sqlite {
	return &Sqlite{
		cache: cach,
		db:    db,
	}
}

func (sl Sqlite) PopulateSchema(ctx context.Context) error {
	_, err := sl.db.ExecContext(ctx, createSchemaStmt)
	return err
}
