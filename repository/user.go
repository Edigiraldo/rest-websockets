package repository

import (
	"context"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) (err error)
	GetUserById(ctx context.Context, id int64) (user *models.User, err error)
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) (err error) {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (user *models.User, err error) {
	return implementation.GetUserById(ctx, id)
}
