package repository

import (
	"context"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type PostRepository interface {
	InsertPost(ctx context.Context, post *models.Post) (err error)
	GetPostById(ctx context.Context, id, userId string) (post *models.Post, err error)
	Close() error
}

var postRepoImplementation PostRepository

func SetPostRepository(repository PostRepository) {
	postRepoImplementation = repository
}

func InsertPost(ctx context.Context, post *models.Post) (err error) {
	return postRepoImplementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id, userId string) (post *models.Post, err error) {
	return postRepoImplementation.GetPostById(ctx, id, userId)
}

func ClosePostRepo() error {
	return postRepoImplementation.Close()
}
