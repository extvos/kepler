package servlet

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type Publisher interface {
	Publish(topic string, data interface{}) error
}

type Subscriber interface {
	Subscribe(topic string, handler MessageProc, channel ...string) error
}

type Context interface {
	DB(...string) *sql.DB
	Redis(...string) *redis.Client
	Publisher(...string) Publisher
	Subscriber(...string) Subscriber
}

type RequestContext interface {
	echo.Context
	Context
}

type TaskContext interface {
	Context
}
