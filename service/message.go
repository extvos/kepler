package service

import "github.com/extvos/kepler/servlet"

type PublishConnector func(cfg servlet.Config) (servlet.Publisher, error)
type SubscribeConnector func(cfg servlet.Config) (servlet.Subscriber, error)
