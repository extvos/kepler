package service

import (
	"github.com/extvos/kepler/servlet"
	"github.com/gofiber/fiber/v2"
)

type keplerContext struct {
	ctx *fiber.Ctx
	svr *KeplerService
}

func (ctx keplerContext) Ctx() *fiber.Ctx {
	return ctx.ctx
}

func (ctx keplerContext) Next() error {
	return ctx.ctx.Next()
}

func (ctx keplerContext) Session() servlet.Session {
	return nil
}

func (ctx keplerContext) SQL(name ...string) servlet.SQL {
	if nil == ctx.svr.dbMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.dbMap[name[0]]
	} else {
		return ctx.svr.dbMap[DefaultName]
	}
}

func (ctx keplerContext) Redis(name ...string) servlet.Redis {
	if nil == ctx.svr.redisMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.redisMap[name[0]]
	} else {
		return ctx.svr.redisMap[DefaultName]
	}
}

func (ctx keplerContext) Publisher(name ...string) servlet.Publisher {
	if nil == ctx.svr.redisMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.pubMap[name[0]]
	} else {
		return ctx.svr.pubMap[DefaultName]
	}
}

func (ctx keplerContext) Subscriber(name ...string) servlet.Subscriber {
	if nil == ctx.svr.subMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.subMap[name[0]]
	} else {
		return ctx.svr.subMap[DefaultName]
	}
}

func (ctx keplerContext) Resource(name string) interface{} {
	return ctx.svr.resMap[name]
}

func (ctx keplerContext) Config(key ...string) servlet.Config {
	if len(key) > 0 {
		return ctx.svr.cfg.Sub(key[0])
	}
	return ctx.svr.cfg
}
