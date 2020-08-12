package service

import (
	"github.com/extvos/kepler/servlet"
)

type RedisConnector func(cfg servlet.Config, name ...string) (servlet.Redis, error)
