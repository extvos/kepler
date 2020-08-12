package service

import (
	"fmt"
	"github.com/extvos/kepler/servlet"
	"strings"
)

func DefaultRedisConnector(cfg servlet.Config) (servlet.Redis, error) {
	driver := cfg.GetString("driver", "redis")
	switch strings.ToLower(driver) {
	case "redis":
		return RedisConnector(cfg)
	case "sentinel":
		return RedisSentinelConnector(cfg)
	default:
		return RedisConnector(cfg)
	}
}

func RedisConnector(cfg servlet.Config) (servlet.Redis, error) {
	return nil, fmt.Errorf("not implemented")
}

func RedisSentinelConnector(cfg servlet.Config) (servlet.Redis, error) {
	return nil, fmt.Errorf("not implemented")
}
