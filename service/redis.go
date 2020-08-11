package service

import (
	"github.com/extvos/kepler/servlet"
	"github.com/go-redis/redis"
)

type RedisConnector func(cfg servlet.Config, name ...string) (*redis.Client, error)
