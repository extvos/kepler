package service

import (
    "github.com/gofiber/fiber/v2"
    "github.com/spf13/viper"
    "github.com/valyala/fasthttp"

    "github.com/extvos/kepler/servlet"
)

func New(cfg servlet.Config) (*allInOneService, error) {
    var svr = allInOneService{}
    var fcfg = fiber.Config{}
    svr.App = *fiber.New(fcfg)
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

func ProbeInit(t servlet.TaskProc) {
    builtinService.ProbeInit(t)
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
func Use(intfs ...interface{}) fiber.Router {
    return builtinService.Use(intfs...)
}

// Connect registers a new Connect fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Connect(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Connect(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Delete registers a new Delete fiber.Route for a path with matching handler in the fiber.Router
// with optional fiber.Route-level middleware.
func Delete(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Delete(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Get registers a new Get fiber.Route for a path with matching handler in the fiber.Router
// with optional fiber.Route-level middleware.
func Get(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Get(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Head registers a new Head fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Head(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Head(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Options registers a new Options fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Options(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Options(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Patch registers a new Patch fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Patch(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Patch(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Post registers a new Post fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Post(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Post(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Put registers a new Put fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Put(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Put(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Trace registers a new Trace fiber.Route for a path with matching handler in the
// fiber.Router with optional fiber.Route-level middleware.
func Trace(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Trace(path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// All registers a new fiber.Route for all HTTP methods and path with matching handler
// in the fiber.Router with optional fiber.Route-level middleware.
func All(path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.All(path, builtinService.handlerFuncs(handlerFuncs...)...)
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
func Add(method, path string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Add(method, path, builtinService.handlerFuncs(handlerFuncs...)...)
}

// Group creates a new fiber.Router group with prefix and optional group-level middleware.
func Group(prefix string, handlerFuncs ...servlet.HandlerFunc) fiber.Router {
    return builtinService.Group(prefix, builtinService.handlerFuncs(handlerFuncs...)...)
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
func AcquireContext(fctx *fasthttp.RequestCtx) *fiber.Ctx {
    return builtinService.AcquireCtx(fctx)
}

// ReleaseContext returns the `RequestContext` instance back to the pool.
// You must call it after `AcquireContext()`.
func ReleaseContext(c *fiber.Ctx) {
    builtinService.ReleaseCtx(c)
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
// func ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	builtinService.Se(w, r)
// }

var builtinService *allInOneService

func init() {
    if s, e := New(servlet.MakeConfig(viper.GetViper())); nil != e {
        panic(e)
    } else {
        builtinService = s
    }
}
