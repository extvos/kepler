package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/extvos/kepler/servlet"
)

type allInOneService struct {
	echo.Echo
	db        *sql.DB
	redis     *redis.Client
	initTasks []servlet.TaskFunc
}

func New(cfg *viper.Viper) (*allInOneService, error) {
	var svr = allInOneService{}
	var e = echo.New()
	svr.Echo = *e
	if nil != cfg {
		if e := svr.Config(cfg); nil != e {
			return nil, e
		}
	}
	return &svr, nil
}

func Config(cfg *viper.Viper) error {
	return builtinService.Config(cfg)
}

func Initialize() error {
	return builtinService.Initialize()
}

func ProbeInit(t servlet.TaskFunc) {
	builtinService.ProbeInit(t)
}

func (svr *allInOneService) Config(cfg *viper.Viper) error {
	return nil
}

func (svr *allInOneService) Initialize() error {
	for _, t := range svr.initTasks {
		if e := t(svr.context(nil)); nil != e {
			return e
		}
	}
	return nil
}

func (svr *allInOneService) ProbeInit(t servlet.TaskFunc) {
	svr.initTasks = append(svr.initTasks, t)
}

func (svr *allInOneService) context(ctx echo.Context) servlet.RequestContext {
	return &allInOneContext{
		svr:     svr,
		Context: ctx,
	}
}

func (svr *allInOneService) handlerFunc(f servlet.HandlerFunc) echo.HandlerFunc {
	var ff = func(ctx echo.Context) error {
		return f(svr.context(ctx))
	}
	return ff
}

func (svr *allInOneService) mw(m servlet.MiddlewareFunc) echo.MiddlewareFunc {
	var ff = func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return svr.handlerFunc(m(func(ctx servlet.RequestContext) error {
			return handlerFunc(ctx)
		}))
	}
	return ff
}

func (svr *allInOneService) middleware(m ...servlet.MiddlewareFunc) []echo.MiddlewareFunc {
	var ms []echo.MiddlewareFunc
	for _, x := range m {
		ms = append(ms, svr.mw(x))
	}
	return ms
}

func Start(address string) error {
	return builtinService.Start(address)
}

func StartServer(s *http.Server) (err error) {
	return builtinService.StartServer(s)
}
func StartAutoTLS(address string) error {
	return builtinService.StartAutoTLS(address)
}
func StartTLS(address string, certFile, keyFile string) (err error) {
	return builtinService.StartTLS(address, certFile, keyFile)
}
func Close() error {
	return builtinService.Close()
}
func Shutdown(ctx context.Context) error {
	return builtinService.Shutdown(ctx)
}

// Pre adds middleware to the chain which is run before echo.Router.
func Pre(middleware ...servlet.MiddlewareFunc) {
	builtinService.Pre(builtinService.middleware(middleware...)...)
}

// Use adds middleware to the chain which is run after echo.Router.
func Use(middleware ...servlet.MiddlewareFunc) {
	builtinService.Use(builtinService.middleware(middleware...)...)
}

// CONNECT registers a new CONNECT echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func CONNECT(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.CONNECT(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// DELETE registers a new DELETE echo.Route for a path with matching handler in the echo.Router
// with optional echo.Route-level middleware.
func DELETE(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.DELETE(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// GET registers a new GET echo.Route for a path with matching handler in the echo.Router
// with optional echo.Route-level middleware.
func GET(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.GET(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// HEAD registers a new HEAD echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func HEAD(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.HEAD(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// OPTIONS registers a new OPTIONS echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func OPTIONS(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.OPTIONS(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// PATCH registers a new PATCH echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func PATCH(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.PATCH(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// POST registers a new POST echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func POST(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.POST(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// PUT registers a new PUT echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func PUT(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.PUT(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// TRACE registers a new TRACE echo.Route for a path with matching handler in the
// echo.Router with optional echo.Route-level middleware.
func TRACE(path string, h servlet.HandlerFunc, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.TRACE(path, builtinService.handlerFunc(h), builtinService.middleware(m...)...)
}

// Any registers a new echo.Route for all HTTP methods and path with matching handler
// in the echo.Router with optional echo.Route-level middleware.
func Any(path string, handler servlet.HandlerFunc, middleware ...servlet.MiddlewareFunc) []*echo.Route {
	return builtinService.Any(path, builtinService.handlerFunc(handler), builtinService.middleware(middleware...)...)
}

// Match registers a new echo.Route for multiple HTTP methods and path with matching
// handler in the echo.Router with optional echo.Route-level middleware.
func Match(methods []string, path string, handler servlet.HandlerFunc, middleware ...servlet.MiddlewareFunc) []*echo.Route {
	return builtinService.Match(methods, path, builtinService.handlerFunc(handler), builtinService.middleware(middleware...)...)
}

// Static registers a new echo.Route with path prefix to serve static files from the
// provided root directory.
func Static(prefix, root string) *echo.Route {
	return builtinService.Static(prefix, root)
}

// File registers a new echo.Route with path to serve a static file with optional echo.Route-level middleware.
func File(path, file string, m ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.File(path, file, builtinService.middleware(m...)...)
}

// Add registers a new echo.Route for an HTTP method and path with matching handler
// in the echo.Router with optional echo.Route-level middleware.
func Add(method, path string, handler servlet.HandlerFunc, middleware ...servlet.MiddlewareFunc) *echo.Route {
	return builtinService.Add(method, path, builtinService.handlerFunc(handler), builtinService.middleware(middleware...)...)
}

// Group creates a new echo.Router group with prefix and optional group-level middleware.
func Group(prefix string, m ...servlet.MiddlewareFunc) (g *echo.Group) {
	return builtinService.Group(prefix, builtinService.middleware(m...)...)
}

// URI generates a URI from handler.
func URI(handler servlet.HandlerFunc, params ...interface{}) string {
	return builtinService.URI(builtinService.handlerFunc(handler), params...)
}

// URL is an alias for `URI` function.
func URL(h servlet.HandlerFunc, params ...interface{}) string {
	return builtinService.URL(builtinService.handlerFunc(h), params...)
}

// Reverse generates an URL from echo.Route name and provided parameters.
func Reverse(name string, params ...interface{}) string {
	return builtinService.Reverse(name, params...)
}

// echo.Routes returns the registered echo.Routes.
func Routes() []*echo.Route {
	return builtinService.Routes()
}

// AcquireContext returns an empty `RequestContext` instance from the pool.
// You must return the context by calling `ReleaseContext()`.
func AcquireContext() echo.Context {
	return builtinService.AcquireContext()
}

// ReleaseContext returns the `RequestContext` instance back to the pool.
// You must call it after `AcquireContext()`.
func ReleaseContext(c echo.Context) {
	builtinService.ReleaseContext(c)
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	builtinService.ServeHTTP(w, r)
}

var builtinService *allInOneService

func init() {
	if s, e := New(viper.GetViper()); nil != e {
		panic(e)
	} else {
		builtinService = s
	}
}
