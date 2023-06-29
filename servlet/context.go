package servlet

import "github.com/gofiber/fiber/v2"

// Context
// Generic context interface
type Context interface {
	// SQL
	// get a SQL instance by name or default one
	SQL(...string) SQL
	// Redis
	// get a redis instance by name or default one
	Redis(...string) Redis
	// Config
	// get the root config or sub config by name
	Config(...string) Config
	// Gear
	// get the registered gear by name
	Gear(name string) interface{}
}

// RequestContext
// Extension context for HTTP Request
type RequestContext interface {
	Context
	// Session
	// get the session object
	Session() Session
	// Ctx
	// the detail context of fiber
	Ctx() *fiber.Ctx
	Next() error
}

// TaskContext
// Extension context for builtin tasks
type TaskContext interface {
	Context
}
