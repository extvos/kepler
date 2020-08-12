package service

import (
	"github.com/extvos/kepler/servlet"
)

type SqlConnector func(cfg servlet.Config) (servlet.SQL, error)
