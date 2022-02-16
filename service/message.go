package service

import (
	"fmt"

	"github.com/extvos/kepler/servlet"
)

func DefaultPubConnector(cfg servlet.Config) (servlet.Publisher, error) {
	return nil, fmt.Errorf("not implemented")
}

func DefaultSubConnector(cfg servlet.Config) (servlet.Subscriber, error) {
	return nil, fmt.Errorf("not implemented")
}
