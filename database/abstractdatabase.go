package database

import "database/sql"

type AbstractDatabase interface {
	GetConnection() (*sql.DB, error)
	Close() error
}
