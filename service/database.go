package service

import (
	"fmt"
	"strings"

	"github.com/extvos/kepler/servlet"
)

func DefaultDBConnector(cfg servlet.Config) (servlet.SQL, error) {
	driver := cfg.GetString("driver", "postgres")
	switch strings.ToLower(driver) {
	case "mysql":
		return MySQLConnector(cfg)
	case "sqlite":
		return SQLiteConnector(cfg)
	default:
		return PostgresConnector(cfg)
	}
}

func SQLiteConnector(cfg servlet.Config) (servlet.SQL, error) {
	return nil, fmt.Errorf("not implemented")
}

func MySQLConnector(cfg servlet.Config) (servlet.SQL, error) {
	return nil, fmt.Errorf("not implemented")
}

func PostgresConnector(cfg servlet.Config) (servlet.SQL, error) {
	return nil, fmt.Errorf("not implemented")
}
