package service

import (
	"github.com/extvos/kepler/servlet"
	"github.com/gofiber/fiber/v2"
)

type dbConnector struct {
	name      string
	connector servlet.SqlConnector
}

type redisConnector struct {
	name      string
	connector servlet.RedisConnector
}

type pubConnector struct {
	name      string
	connector servlet.PublishConnector
}

type subConnector struct {
	name      string
	connector servlet.SubscribeConnector
}

type allInOneService struct {
	fiber.App
	cfg             servlet.Config
	initTasks       []servlet.TaskProc
	dbMap           map[string]servlet.SQL
	redisMap        map[string]servlet.Redis
	pubMap          map[string]servlet.Publisher
	subMap          map[string]servlet.Subscriber
	resMap          map[string]interface{}
	dbConnectors    []dbConnector
	redisConnectors []redisConnector
	pubConnectors   []pubConnector
	subConnectors   []subConnector
}

const (
	DefaultName = "*"
)

func (svr *allInOneService) RequireDatabase(name string, connector ...servlet.SqlConnector) {
	var c = dbConnector{name: name, connector: DefaultDBConnector}
	if len(connector) > 0 {
		c.connector = connector[0]
	}
	svr.dbConnectors = append(svr.dbConnectors, c)
}

func (svr *allInOneService) RequireRedis(name string, connector ...servlet.RedisConnector) {
	var c = redisConnector{name: name, connector: DefaultRedisConnector}
	if len(connector) > 0 {
		c.connector = connector[0]
	}
	svr.redisConnectors = append(svr.redisConnectors, c)
}

func (svr *allInOneService) RequirePublisher(name string, connector ...servlet.PublishConnector) {
	var c = pubConnector{name: name, connector: DefaultPubConnector}
	if len(connector) > 0 {
		c.connector = connector[0]
	}
	svr.pubConnectors = append(svr.pubConnectors, c)
}

func (svr *allInOneService) RequireSubscriber(name string, connector ...servlet.SubscribeConnector) {
	var c = subConnector{name: name, connector: DefaultSubConnector}
	if len(connector) > 0 {
		c.connector = connector[0]
	}
	svr.subConnectors = append(svr.subConnectors, c)
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
	svr.resMap[name] = res
}

func (svr *allInOneService) configDatabase() error {
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

func (svr *allInOneService) configRedis() error {
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

func (svr *allInOneService) configPublish() error {
	for i, c := range svr.pubConnectors {
		if p, e := c.connector(svr.cfg.Sub(c.name)); nil != e {
			return e
		} else {
			svr.pubMap[c.name] = p
			if i == 0 {
				svr.pubMap[DefaultName] = p
			}
		}
	}
	return nil
}

func (svr *allInOneService) configSubscribe() error {
	for i, c := range svr.subConnectors {
		if p, e := c.connector(svr.cfg.Sub(c.name)); nil != e {
			return e
		} else {
			svr.subMap[c.name] = p
			if i == 0 {
				svr.subMap[DefaultName] = p
			}
		}
	}
	return nil
}

func (svr *allInOneService) Config(cfg servlet.Config) error {
	svr.cfg = cfg
	svr.dbMap = make(map[string]servlet.SQL)
	svr.redisMap = make(map[string]servlet.Redis)
	svr.pubMap = make(map[string]servlet.Publisher)
	svr.subMap = make(map[string]servlet.Subscriber)
	svr.resMap = make(map[string]interface{})
	if e := svr.configDatabase(); nil != e {
		return e
	}
	if e := svr.configRedis(); nil != e {
		return e
	}
	if e := svr.configPublish(); nil != e {
		return e
	}
	if e := svr.configSubscribe(); nil != e {
		return e
	}
	return nil
}

func (svr *allInOneService) context(ctx *fiber.Ctx) servlet.RequestContext {
	return &allInOneContext{
		svr: svr,
		ctx: ctx,
	}
}

func (svr *allInOneService) handlerFunc(f servlet.HandlerFunc) fiber.Handler {
	var ff = func(ctx *fiber.Ctx) error {
		return f(svr.context(ctx))
	}
	return ff
}

func (svr *allInOneService) handlerFuncs(f ...servlet.HandlerFunc) []fiber.Handler {
	var handlers []fiber.Handler
	for _, h := range f {
		handlers = append(handlers, svr.handlerFunc(h))
	}
	return handlers
}

// func (svr *allInOneService) mw(m servlet.HandlerFunc) fiber.Handler {
// 	var ff = func(ctx *fiber.Ctx) error {
// 		return m(svr.context(ctx))
// 	}
// 	return ff
// }

// func (svr *allInOneService) middleware(m ...servlet.MiddlewareFunc) []fiber.Handler {
// 	var ms []fiber.Handler
// 	for _, x := range m {
// 		ms = append(ms, svr.mw(x))
// 	}
// 	return ms
// }
