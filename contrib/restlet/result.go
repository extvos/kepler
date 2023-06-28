package restlet

type resultCtrl struct {
	Total    int64 `json:"total,omitempty"`
	Count    int64 `json:"count,omitempty"`
	Limit    int64 `json:"limit,omitempty"`
	Offset   int64 `json:"offset,omitempty"`
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
	Pages    int64 `json:"pages,omitempty"`
}

type Result[T any] struct {
	Code    int         `json:"code"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Control *resultCtrl `json:"_ctrl,omitempty"`
}

func (r Result[T]) D(d T) Result[T] {
	r.Data = d
	return r
}

func (r Result[T]) Ctrl(total, limit, offset int64) Result[T] {
	ctrl := resultCtrl{
		Total: total, Limit: limit, Offset: offset,
	}
	ctrl.PageSize = limit
	ctrl.Page = offset / limit
	if total%limit == 0 {
		ctrl.Pages = total / limit
	} else {
		ctrl.Pages = total/limit + 1
	}
	r.Control = &ctrl
	return r
}
