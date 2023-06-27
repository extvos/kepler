package servlet

import "github.com/gofiber/fiber/v2"

// Context
// Generic context interface
type Context interface {
	SQL(...string) SQL
	Redis(...string) Redis
	Publisher(...string) Publisher
	Subscriber(...string) Subscriber
	Config(...string) Config
	Resource(name string) interface{}
}

// RequestContext
// Extension context for HTTP Request
type RequestContext interface {
	Context
	Session() Session
	Ctx() *fiber.Ctx
	Next() error
}

// TaskContext
// Extension context for builtin tasks
type TaskContext interface {
	Context
}
