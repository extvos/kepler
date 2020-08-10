package service

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type allInOneContext struct {
	echo.Context
	svr *allInOneService
}

func (ctx allInOneContext) DB(name ...string) *sql.DB {
	return ctx.svr.db
}

func (ctx allInOneContext) Redis(name ...string) *redis.Client {
	return ctx.svr.redis
}
