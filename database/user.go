package database

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type UserDatabase struct {
	implementation AbstractDatabase
}

func NewUserDatabase(databaseImplementation AbstractDatabase) *UserDatabase {
	return &UserDatabase{
		implementation: databaseImplementation,
	}
}

func (repo *UserDatabase) InsertUser(ctx context.Context, user *models.User) (err error) {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		user.ID, user.Email, user.Password)

	return err
}

func (repo *UserDatabase) GetUserById(ctx context.Context, id string) (user *models.User, err error) {
	user = &models.User{}
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return nil, err
	}
	row := db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	if err = row.Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserDatabase) GetUserByEmail(ctx context.Context, email string) (user *models.User, err error) {
	user = &models.User{}
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return nil, err
	}
	row := db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	if err = row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserDatabase) Close() error {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return err
	}

	return db.Close()
}
