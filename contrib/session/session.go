package session

import (
	"github.com/extvos/kepler/servlet"
	log "github.com/sirupsen/logrus"
)

func ProbeHandler(ctx servlet.RequestContext) error {
	log.Debugln("ProbeHandler:> ", ctx.Ctx().Method(), ctx.Ctx().Path())
	ctx.Ctx().Cookies("K_SESS_ID")
	return ctx.Next()
}
