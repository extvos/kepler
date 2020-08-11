package service

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"

	"github.com/extvos/kepler/servlet"
)

type allInOneService struct {
	echo.Echo
	//db        *sql.DB
	//redis     *redis.Client
	initTasks []servlet.TaskProc
	dbMap     map[string]*sql.DB
	redisMap  map[string]*redis.Client
	pubMap    map[string]servlet.Publisher
	subMap    map[string]servlet.Subscriber
	resMap    map[string]interface{}

	dbConnectors []struct {
		name      string
		connector DBConnector
	}
	redisConnectors []struct {
		name      string
		connector RedisConnector
	}
	pubConnectors []struct {
		name      string
		connector PublishConnector
	}
	subConnectors []struct {
		name      string
		connector SubscribeConnector
	}
}

const (
	DefaultName = "*"
)

func (svr *allInOneService) Config(cfg servlet.Config) error {
	return nil
}

func (svr *allInOneService) RequireDatabase(name string, connector ...DBConnector) {

}

func (svr *allInOneService) RequireRedis(name string, connector ...RedisConnector) {

}

func (svr *allInOneService) RequirePublisher(name string, connector ...PublishConnector) {

}

func (svr *allInOneService) RequireSubscriber(name string, connector ...SubscribeConnector) {

}

func (svr *allInOneService) Initialize() error {
	for _, t := range svr.initTasks {
		if e := t(svr.context(nil)); nil != e {
			return e
		}
	}
	return nil
}

func (svr *allInOneService) ProbeInit(t servlet.TaskProc) {
	svr.initTasks = append(svr.initTasks, t)
}

func (svr *allInOneService) ProbeResource(name string, res interface{}) {

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
