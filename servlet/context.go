package servlet

import (
	"github.com/labstack/echo/v4"
)

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
