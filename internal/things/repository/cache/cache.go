package cache

import (
	"context"

	"github.com/danic95/things-service/goutils/dbfactory"
	"github.com/danic95/things-service/goutils/settings"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client  *redis.Client
	enabled bool
}

func New(ctx context.Context, sett *settings.Settings) (*Cache, error) {

	var rclient *redis.Client
	var err error

	if sett.Cache.Enabled {
		rclient, err = dbfactory.CreateRedisConnection(ctx, &sett.Cache)
	}

	c := &Cache{
		client:  rclient,
		enabled: true}

	return c, err

}
