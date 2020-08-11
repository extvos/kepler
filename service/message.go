package service

import "github.com/extvos/kepler/servlet"

type Publisher interface {
	Publish(topic string, data interface{}) error
}

type Subscriber interface {
	Subscribe(topic string, handler servlet.MessageProc, channel ...string) error
}
