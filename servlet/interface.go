package servlet

import "database/sql"

type Publisher interface {
	Publish(topic string, data interface{}) error
}

type Subscriber interface {
	Subscribe(topic string, handler MessageProc, channel ...string) error
}

type Redis interface {
	// K:V
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int64) error
	Keys(string) []string
	Exists(...string) bool
	Del(string) error

	// LIST
	LPush(string, interface{}) error
	RPush(string, interface{}) error
	LPop(string) (interface{}, error)
	RPop(string) (interface{}, error)
	LRange(string) (int, error)

	// MAP
	MSet(string, string, interface{}, ...int64) error
	MGet(string, string) (interface{}, error)

	// Increase and Decrease
	Incr(string, ...int64) (int64, error)
	Decr(string, ...int64) (int64, error)

	// Expire
	Expire(string, int64) error
	ExpireAt(string, int64) error
}

type SQL interface {
	DB() *sql.DB
	Name() string
	Driver() string
}
