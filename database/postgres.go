package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(URI string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", URI)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) (err error) {
	_, err = repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		user.ID, user.Email, user.Password)

	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (user *models.User, err error) {
	user = &models.User{}
	row := repo.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	if err = row.Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (user *models.User, err error) {
	user = &models.User{}
	row := repo.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	if err = row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
