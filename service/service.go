package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"

	"github.com/extvos/kepler/servlet"
)

func New(cfg servlet.Config) (*KeplerService, error) {
	var svr = KeplerService{}
	var fcfg = fiber.Config{}
	svr._app = fiber.New(fcfg)
	if nil != cfg {
		if e := svr.Config(cfg); nil != e {
			return nil, e
		}
	}
	return &svr, nil
}

func Config(cfg servlet.Config) error {
	return builtinService.Config(cfg)
}

func Initialize() error {
	return builtinService.Initialize()
}

func MountInitialization(t servlet.TaskProc) {
	builtinService.MountInitialization(t)
}

func Listen(address string) error {
	return builtinService.Listen(address)
	// return builtinService.Listen(address)
}

func ListenTLS(address string, certFile, keyFile string) (err error) {
	return builtinService.ListenTLS(address, certFile, keyFile)
}
func Shutdown() error {
	return builtinService.Shutdown()
}

// Use adds middleware to the chain which is run after fiber.Router.
func Use(handlers ...interface{}) fiber.Router {
	return builtinService.Use(handlers...)
}

// Connect registers a new Connect fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Connect(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Connect(path, handlers...)
}

// Delete registers a new Delete fiber.Route for a path with matching handler in the fiber.Router
// with optional fiber.Route-level middleware.
func Delete(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Delete(path, handlers...)
}

// Get registers a new Get fiber.Route for a path with matching handler in the fiber.Router
// with optional fiber.Route-level middleware.
func Get(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Get(path, handlers...)
}

// Head registers a new Head fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Head(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Head(path, handlers...)
}

// Options registers a new Options fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Options(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Options(path, handlers...)
}

// Patch registers a new Patch fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Patch(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Patch(path, handlers...)
}

// Post registers a new Post fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Post(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Post(path, handlers...)
}

// Put registers a new Put fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Put(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Put(path, handlers...)
}

// Trace registers a new Trace fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Trace(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Trace(path, handlers...)
}

// All registers a new fiber.Route for all HTTP methods and path with matching handler
// in the fiber.Router with optional fiber.Route-level middleware.
func All(path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.All(path, handlers...)
}

// Match registers a new fiber.Route for multiple HTTP methods and path with matching
// handler in the fiber.Router with optional fiber.Route-level middleware.
// func Match(methods []string, path string, handler servlet.HandlerFunc, middleware ...servlet.MiddlewareFunc) []*fiber.Router {
// 	return builtinService.Match(methods, path, builtinService.handlerFunc(handler), builtinService.middleware(middleware...)...)
// }

// Static registers a new fiber.Route with path prefix to serve static files from the
// provided root directory.
func Static(prefix, root string) fiber.Router {
	return builtinService.Static(prefix, root)
}

// File registers a new fiber.Route with path to serve a static file with optional fiber.Route-level middleware.
// func File(path, file string, m ...servlet.MiddlewareFunc) *fiber.Route {
//     return builtinService.File(path, file, builtinService.middleware(m...)...)
// }

// Add registers a new fiber.Route for an HTTP method and path with matching handler
// in the fiber.Router with optional fiber.Route-level middleware.
func Add(method, path string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Add(method, path, handlers...)
}

// Group creates a new fiber.Router group with prefix and optional group-level middleware.
func Group(prefix string, handlers ...servlet.HandlerFunc) fiber.Router {
	return builtinService.Group(prefix, handlers...)
}

// Mount another service into current service as seperated parts.
func Mount(prefix string, service *KeplerService) fiber.Router {
	return service.Mount(prefix, service)
}

// // URI generates a URI from handler.
// func Uri(handler servlet.HandlerFunc, params ...interface{}) string {
// 	return builtinService.Uri(builtinService.handlerFunc(handler), params...)
// }
//
// // URL is an alias for `URI` function.
// func URL(h servlet.HandlerFunc, params ...interface{}) string {
// 	return builtinService.URL(builtinService.handlerFunc(h), params...)
// }
//
// // Reverse generates an URL from fiber.Route name and provided parameters.
// func Reverse(name string, params ...interface{}) string {
// 	return builtinService.Reverse(name, params...)
// }

// fiber.Routes returns the registered fiber.Routes.
// func Routes() []*fiber.Route {
// 	return builtinService.Routes()
// }

// AcquireContext returns an empty `RequestContext` instance from the pool.
// You must return the context by calling `ReleaseContext()`.
func AcquireContext(ctx *fasthttp.RequestCtx) *fiber.Ctx {
	return builtinService.AcquireContext(ctx)
}

// ReleaseContext returns the `RequestContext` instance back to the pool.
// You must call it after `AcquireContext()`.
func ReleaseContext(c *fiber.Ctx) {
	builtinService.ReleaseContext(c)
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
// func ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	builtinService.Se(w, r)
// }

var builtinService *KeplerService

func init() {
	if s, e := New(servlet.MakeConfig(viper.GetViper())); nil != e {
		panic(e)
	} else {
		builtinService = s
	}
}
