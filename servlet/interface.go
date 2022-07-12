package servlet

import "database/sql"

// Publisher defines the interface for publishing topics.
type Publisher interface {
	Publish(topic string, data interface{}) error
}

// Subscriber  defines the interface for subscribing messages from topics.
type Subscriber interface {
	Subscribe(topic string, handler MessageProc, channel ...string) error
}

// Redis defines the interface for Redis operations.
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
	LLen(string) int

	// HASH
	HSet(string, string, interface{}, ...int64) error
	HGet(string, string) (interface{}, error)
	HDel(string, string) error
	HExists(string, string) bool
	HLen(string) int
	HKeys(string, string) []string

	// Increase and Decrease
	Incr(string, ...int64) (int64, error)
	Decr(string, ...int64) (int64, error)

	// Expire
	Expire(string, int64) error
	ExpireAt(string, int64) error
}

// SQL defines the interface for SQL Database operations.
type SQL interface {
	DB() *sql.DB
	Name() string
	Driver() string
}

// Cache defines the interface for cache operations.
type Cache interface {
	Bucket(string) error
}
