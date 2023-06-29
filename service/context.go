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
	if nil == ctx.svr.dbMap || len(ctx.svr.dbMap) == 0 {
		panic("no database was mounted")
	}
	if len(name) > 0 {
		return ctx.svr.dbMap[name[0]]
	} else {
		return ctx.svr.dbMap[DefaultName]
	}
}

func (ctx keplerContext) Redis(name ...string) servlet.Redis {
	if nil == ctx.svr.redisMap || len(ctx.svr.redisMap) == 0 {
		panic("no redis was mounted")
	}
	if len(name) > 0 {
		return ctx.svr.redisMap[name[0]]
	} else {
		return ctx.svr.redisMap[DefaultName]
	}
}

func (ctx keplerContext) Gear(name string) interface{} {
	p, b := ctx.svr.gearsMap[name]
	if !b {
		panic("no gear named '" + name + "' was registered")
	}
	return p
}

func (ctx keplerContext) Config(key ...string) servlet.Config {
	if len(key) > 0 {
		return ctx.svr.cfg.Sub(key[0])
	}
	return ctx.svr.cfg
}
