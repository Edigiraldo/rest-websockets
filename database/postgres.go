package database

import (
	"database/sql"
	"errors"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewDatabase(URI string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", URI)
	if err != nil {
		return nil, err
	}

	return &PostgresDatabase{db: db}, nil
}

func (repo *PostgresDatabase) GetConnection() (*sql.DB, error) {
	if repo.db == nil {
		return nil, errors.New("connection has not been created yet")
	}

	return repo.db, nil
}

func (repo *PostgresDatabase) Close() error {
	return repo.db.Close()
}
