package service

import (
	"github.com/go-redis/redis"

	"github.com/extvos/kepler/servlet"
)

type RedisConnector func(cfg servlet.Config, name ...string) (*redis.Client, error)
