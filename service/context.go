package service

import (
	"github.com/labstack/echo/v4"

	"github.com/extvos/kepler/servlet"
)

type allInOneContext struct {
	echo.Context
	svr *allInOneService
}

func (ctx allInOneContext) SQL(name ...string) servlet.SQL {
	if nil == ctx.svr.dbMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.dbMap[name[0]]
	} else {
		return ctx.svr.dbMap[DefaultName]
	}
}

func (ctx allInOneContext) Redis(name ...string) servlet.Redis {
	if nil == ctx.svr.redisMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.redisMap[name[0]]
	} else {
		return ctx.svr.redisMap[DefaultName]
	}
}

func (ctx allInOneContext) Publisher(name ...string) servlet.Publisher {
	if nil == ctx.svr.redisMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.pubMap[name[0]]
	} else {
		return ctx.svr.pubMap[DefaultName]
	}
}

func (ctx allInOneContext) Subscriber(name ...string) servlet.Subscriber {
	if nil == ctx.svr.subMap {
		return nil
	}
	if len(name) > 0 {
		return ctx.svr.subMap[name[0]]
	} else {
		return ctx.svr.subMap[DefaultName]
	}
}
