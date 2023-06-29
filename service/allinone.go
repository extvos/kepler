package service

import (
	"fmt"
	"github.com/extvos/kepler/servlet"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"reflect"
)

/*
*

	Database connector structure
*/
type dbConnector struct {
	name      string
	connector servlet.SqlConnector
}

/*
*

	Redis connector structure
*/
type redisConnector struct {
	name      string
	connector servlet.RedisConnector
}

/*
*

	MSQ publish connector structure
*/
type pubConnector struct {
	name      string
	connector servlet.PublishConnector
}

/*
*

	MSQ subscribe connector structure
*/
type subConnector struct {
	name      string
	connector servlet.SubscribeConnector
}

// KeplerService structure
type KeplerService struct {
	_app            *fiber.App               // the fiber app
	cfg             servlet.Config           // the configuration
	initTasks       []servlet.TaskHandler    // application initialize tasks
	dbMap           map[string]servlet.SQL   // Database connections
	redisMap        map[string]servlet.Redis // Redis connections
	gearsMap        map[string]interface{}   // Gear instances
	dbConnectors    []dbConnector            // Database connectors
	redisConnectors []redisConnector         // Redis connectors
	pubConnectors   []pubConnector           // Publisher connectors
	subConnectors   []subConnector           // Subscriber connectors
}

const (
	DefaultName = "*" // Default name for connections
)

// RequireDatabase
// Get database connection by connector
func (svr *KeplerService) RequireDatabase(name string, connector ...servlet.SqlConnector) {
	var c = dbConnector{name: name, connector: DefaultDBConnector}
	if len(connector) > 0 {
		c.connector = connector[0]
	}
	svr.dbConnectors = append(svr.dbConnectors, c)
}

// RequireRedis
// Get redis connection by connector
func (svr *KeplerService) RequireRedis(name string, connector ...servlet.RedisConnector) {
	var c = redisConnector{name: name, connector: DefaultRedisConnector}
	if len(connector) > 0 {
		c.connector = connector[0]
	}
	svr.redisConnectors = append(svr.redisConnectors, c)
}

// Initialize
// Run service initialization
func (svr *KeplerService) Initialize() error {
	for _, t := range svr.initTasks {
		if e := t(svr.context(nil)); nil != e {
			return e
		}
	}
	return nil
}

// MountInitialization
// Mount initialization tasks
func (svr *KeplerService) MountInitialization(t servlet.TaskHandler) {
	svr.initTasks = append(svr.initTasks, t)
}

// MountGear
// mount gear instances
func (svr *KeplerService) MountGear(name string, res interface{}) {
	svr.gearsMap[name] = res
}

func (svr *KeplerService) configDatabase() error {
	for i, c := range svr.dbConnectors {
		if p, e := c.connector(svr.cfg.Sub(c.name)); nil != e {
			return e
		} else {
			svr.dbMap[c.name] = p
			if i == 0 {
				svr.dbMap[DefaultName] = p
			}
		}
	}
	return nil
}

func (svr *KeplerService) configRedis() error {
	for i, c := range svr.redisConnectors {
		if p, e := c.connector(svr.cfg.Sub(c.name)); nil != e {
			return e
		} else {
			svr.redisMap[c.name] = p
			if i == 0 {
				svr.redisMap[DefaultName] = p
			}
		}
	}
	return nil
}

// Config
// Configure service with specified configuration.
func (svr *KeplerService) Config(cfg servlet.Config) error {
	svr.cfg = cfg
	svr.gearsMap = make(map[string]interface{})
	if cfg.Has("database") {
		svr.dbMap = make(map[string]servlet.SQL)
		if e := svr.configDatabase(); nil != e {
			return e
		}
	}
	if cfg.Has("redis") {
		svr.redisMap = make(map[string]servlet.Redis)
		if e := svr.configRedis(); nil != e {
			return e
		}
	}
	if cfg.Has("gears") {
		svr.gearsMap = make(map[string]interface{})
	}
	return nil
}

func (svr *KeplerService) context(ctx *fiber.Ctx) servlet.RequestContext {
	return &keplerContext{
		svr: svr,
		ctx: ctx,
	}
}

func (svr *KeplerService) handlerFunc(f servlet.RequestHandler) fiber.Handler {
	var ff = func(ctx *fiber.Ctx) error {
		return f(svr.context(ctx))
	}
	return ff
}

func (svr *KeplerService) handlerFuncEx(f ...servlet.RequestHandler) []fiber.Handler {
	var handlers []fiber.Handler
	for _, h := range f {
		handlers = append(handlers, svr.handlerFunc(h))
	}
	return handlers
}

func (svr *KeplerService) Listen(address string) error {
	return svr._app.Listen(address)
	// return builtinService.Listen(address)
}

func (svr *KeplerService) ListenTLS(address string, certFile, keyFile string) (err error) {
	return svr._app.ListenTLS(address, certFile, keyFile)
}
func (svr *KeplerService) Shutdown() error {
	return svr._app.Shutdown()
}

// Use adds middleware to the chain which is run after fiber.Router.
func (svr *KeplerService) Use(args ...interface{}) fiber.Router {
	var parameters []interface{}
	for _, arg := range args {
		switch a := arg.(type) {
		case string:
			parameters = append(parameters, a)
		case []string:
			for _, s := range a {
				parameters = append(parameters, s)
			}
		case servlet.RequestHandler:
			parameters = append(parameters, svr.handlerFunc(a))
		default:
			panic(fmt.Sprintf("use: invalid handler %v\n", reflect.TypeOf(a)))
		}
	}
	return svr._app.Use(parameters...)
}

// Connect registers a new Connect fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Connect(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Connect(path, svr.handlerFuncEx(handlers...)...)
}

// Delete registers a new Delete fiber.Route for a path with matching handler in the fiber.Router
// with optional fiber.Route-level middleware.
func (svr *KeplerService) Delete(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Delete(path, svr.handlerFuncEx(handlers...)...)
}

// Get registers a new Get fiber.Route for a path with matching handler in the fiber.Router
// with optional fiber.Route-level middleware.
func (svr *KeplerService) Get(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Get(path, svr.handlerFuncEx(handlers...)...)
}

// Head registers a new Head fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Head(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Head(path, svr.handlerFuncEx(handlers...)...)
}

// Options registers a new Options fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Options(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Options(path, svr.handlerFuncEx(handlers...)...)
}

// Patch registers a new Patch fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Patch(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Patch(path, svr.handlerFuncEx(handlers...)...)
}

// Post registers a new Post fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Post(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Post(path, svr.handlerFuncEx(handlers...)...)
}

// Put registers a new Put fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Put(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Put(path, svr.handlerFuncEx(handlers...)...)
}

// Trace registers a new Trace fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Trace(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Trace(path, svr.handlerFuncEx(handlers...)...)
}

// All registers a new fiber.Route for all HTTP methods and path with matching handler
// in the fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) All(path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.All(path, svr.handlerFuncEx(handlers...)...)
}

// Match registers a new fiber.Route for multiple HTTP methods and path with matching
// handler in the fiber.Router with optional fiber.Route-level middleware.
// func Match(methods []string, path string, handler servlet.RequestHandler, middleware ...servlet.RequestMiddleware) []*fiber.Router {
// 	return builtinService.Match(methods, path, builtinService.handlerFunc(handler), builtinService.middleware(middleware...)...)
// }

// Static registers a new fiber.Route with path prefix to serve static files from the
// provided root directory.
func (svr *KeplerService) Static(prefix, root string) fiber.Router {
	return svr._app.Static(prefix, root)
}

// File registers a new fiber.Route with path to serve a static file with optional fiber.Route-level middleware.
// func File(path, file string, m ...servlet.RequestMiddleware) *fiber.Route {
//     return svr._app.File(path, file, svr.middleware(m...)...)
// }

// Add registers a new fiber.Route for an HTTP method and path with matching handler
// in the fiber.Router with optional fiber.Route-level middleware.
func (svr *KeplerService) Add(method, path string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Add(method, path, svr.handlerFuncEx(handlers...)...)
}

// Group creates a new fiber.Router group with prefix and optional group-level middleware.
func (svr *KeplerService) Group(prefix string, handlers ...servlet.RequestHandler) fiber.Router {
	return svr._app.Group(prefix, svr.handlerFuncEx(handlers...)...)
}

// Mount another service into current service as seperated parts.
func (svr *KeplerService) Mount(prefix string, service *KeplerService) fiber.Router {
	return svr._app.Mount(prefix, service._app)
}

// AcquireContext returns an empty `RequestContext` instance from the pool.
// You must return the context by calling `ReleaseContext()`.
func (svr *KeplerService) AcquireContext(ctx *fasthttp.RequestCtx) *fiber.Ctx {
	return svr._app.AcquireCtx(ctx)
}

// ReleaseContext returns the `RequestContext` instance back to the pool.
// You must call it after `AcquireContext()`.
func (svr *KeplerService) ReleaseContext(c *fiber.Ctx) {
	svr._app.ReleaseCtx(c)
}
