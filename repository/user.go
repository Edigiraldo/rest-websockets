package repository

import (
	"context"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) (err error)
	GetUserById(ctx context.Context, id string) (user *models.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user *models.User, err error)
	Close() error
}

var userRepoImplementation UserRepository

func SetUserRepository(repository UserRepository) {
	userRepoImplementation = repository
}

func InsertUser(ctx context.Context, user *models.User) (err error) {
	return userRepoImplementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (user *models.User, err error) {
	return userRepoImplementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (user *models.User, err error) {
	return userRepoImplementation.GetUserByEmail(ctx, email)
}

func CloseUserRepo() error {
	return userRepoImplementation.Close()
}
