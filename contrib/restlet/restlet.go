package restlet

import (
	"encoding/json"
	xql "github.com/archsh/go.xql"
	"github.com/extvos/kepler/servlet"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Handler[T xql.TableIdentified] struct {
	_ds  string
	_dao DAO[T]
}

func NewHandler[T xql.TableIdentified](ds ...string) func(servlet.RequestContext) error {
	var h Handler[T]
	if len(ds) < 1 {
		h._ds = "default"
	} else {
		h._ds = ds[0]
	}
	h._dao, _ = NewDAO[T](nil, ds...)
	var f = func(ctx servlet.RequestContext) error {
		log.Debugln("Handler[T]:> ", ctx.Ctx().Method(), ctx.Ctx().Path())
		switch ctx.Ctx().Method() {
		case fiber.MethodGet:
			if q, e1 := h.buildQuery(ctx); nil != e1 {
				return onFailure(ctx, fiber.StatusBadRequest, e1.Error())
			} else if r, e2 := h.onGet(q); nil != e2 {
				return onFailure(ctx, fiber.StatusInternalServerError, e2.Error())
			} else {
				if r.Code == 0 {
					r.Code = fiber.StatusOK
				}
				return ctx.Ctx().Status(fiber.StatusOK).JSON(r)
			}
		case fiber.MethodPost:
			var obj T
			if e1 := json.Unmarshal(ctx.Ctx().Request().Body(), &obj); nil != e1 {
				return onFailure(ctx, fiber.StatusBadRequest, e1.Error())
			} else if r, e2 := h.onPost(obj); nil != e2 {
				return onFailure(ctx, fiber.StatusInternalServerError, e2.Error())
			} else {
				if r.Code == 0 {
					r.Code = fiber.StatusCreated
				}
				return ctx.Ctx().Status(fiber.StatusCreated).JSON(r)
			}
		case fiber.MethodPut:
		case fiber.MethodDelete:
		case fiber.MethodOptions:
		default:
			return ctx.Ctx().Status(fiber.StatusMethodNotAllowed).JSON(Result[T]{Code: fiber.StatusMethodNotAllowed, Message: "Method Not Allowed"})
		}
		return nil
	}
	return f
}

func (*Handler[T]) buildQuery(ctx servlet.RequestContext) (Query, error) {
	return Query{}, nil
}

func onFailure(ctx servlet.RequestContext, status int, msg string) error {
	log.Warnln("onFailure:>", status, msg)
	return ctx.Ctx().Status(status).JSON(Result[string]{Code: status, Message: msg})
}

func (*Handler[T]) onPost(obj T) (Result[T], error) {
	log.Debugln("onPost:>", obj)
	return Result[T]{Data: obj}, nil
}

func (*Handler[T]) onGet(queries Query) (Result[[]T], error) {
	log.Debugln("onGet:>", queries)
	return Result[[]T]{Data: []T{}}, nil
}

func (*Handler[T]) onPut(obj T, queries Query) (Result[int64], error) {
	log.Debugln("onPut:>", obj, queries)
	return Result[int64]{Data: 0}, nil
}

func (*Handler[T]) onDelete(queries Query) (Result[int64], error) {
	log.Debugln("onDelete:>", queries)
	return Result[int64]{Data: 0}, nil
}
