package service

import (
	"database/sql"

	"github.com/extvos/kepler/servlet"
)

type DBConnector func(cfg servlet.Config) (*sql.DB, error)
