package repository

import (
	"context"
	"errors"

	"github.com/danic95/things-service/goutils/settings"
	"github.com/danic95/things-service/internal/things"
	"github.com/danic95/things-service/internal/things/repository/cache"
	"github.com/danic95/things-service/internal/things/repository/mysql"
	"github.com/danic95/things-service/internal/things/repository/sqlite"
	"github.com/jmoiron/sqlx"
	logger "github.com/sanservices/apilogger/v2"
)

// New constructs the repository
func New(ctx context.Context, cfg *settings.Settings, db *sqlx.DB, cach *cache.Cache) (things.Repository, error) {

	switch cfg.DB.Engine {
	case "mysql":
		return mysql.New(db), nil

	case "sqlite":
		repo := sqlite.New(db, cach)
		return repo, repo.PopulateSchema(ctx)

	default:
		err := errors.New("unsupported or missing database engine")
		logger.Error(ctx, logger.LogCatReadConfig, err)
		return nil, err
	}
}
