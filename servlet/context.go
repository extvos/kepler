package servlet

import "github.com/gofiber/fiber/v2"

type Context interface {
	SQL(...string) SQL
	Redis(...string) Redis
	Publisher(...string) Publisher
	Subscriber(...string) Subscriber
	Config(...string) Config
	Resource(name string) interface{}
}

type RequestContext interface {
	Context
	Ctx() *fiber.Ctx
}

type TaskContext interface {
	Context
}
