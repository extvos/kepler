package servlet

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type Context interface {
	DB(...string) *sql.DB
	Redis(...string) *redis.Client
}

type RequestContext interface {
	echo.Context
	Context
}

type TaskContext interface {
	Context
}
