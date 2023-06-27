package servlet

type Session interface {
	ID() string
	Set(key string, val interface{})
	Get(key string) interface{}
}
