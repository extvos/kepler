package service

import (
	"database/sql"
	"fmt"
	"github.com/extvos/kepler/servlet"
)

type DBConnector func(cfg servlet.Config) (*sql.DB, error)

func MySQLConnector(cfg servlet.Config, name ...string) (*sql.DB, error) {
	return nil, fmt.Errorf("not implemented")
}

func PostgresConnector(cfg servlet.Config, name ...string) (*sql.DB, error) {
	return nil, fmt.Errorf("not implemented")
}
