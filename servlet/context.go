package servlet

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Publisher interface {
	Publish(topic string, data interface{}) error
}

type Subscriber interface {
	Subscribe(topic string, handler MessageProc, channel ...string) error
}

type Redis interface {
	// K:V
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int64) error
	Keys(string) []string
	Del(string) error

	// LIST
	LPush(string, interface{}) error
	RPush(string, interface{}) error
	LPop(string) (interface{}, error)
	RPop(string) (interface{}, error)
	LRange(string) (int, error)

	// MAP
	MSet(string, string, interface{}, ...int64) error
	MGet(string, string) (interface{}, error)

	//
	Incr(string, ...int64) (int64, error)
}

type SQL interface {
	DB() *sql.DB
	Name() string
}

type Context interface {
	SQL(...string) SQL
	Redis(...string) Redis
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
