package repository

import (
	"context"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type PostRepository interface {
	ListPosts(ctx context.Context, userId string, page uint64) ([]*models.Post, error)
	InsertPost(ctx context.Context, post *models.Post) (err error)
	GetPostById(ctx context.Context, id, userId string) (post *models.Post, err error)
	UpdatePostById(ctx context.Context, content, id, userId string) (post *models.Post, err error)
	DeletePostById(ctx context.Context, id, userId string) (err error)
	Close() error
}

var postRepoImplementation PostRepository

func SetPostRepository(repository PostRepository) {
	postRepoImplementation = repository
}

func ListPosts(ctx context.Context, userId string, page uint64) ([]*models.Post, error) {
	return postRepoImplementation.ListPosts(ctx, userId, page)
}

func InsertPost(ctx context.Context, post *models.Post) (err error) {
	return postRepoImplementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id, userId string) (post *models.Post, err error) {
	return postRepoImplementation.GetPostById(ctx, id, userId)
}

func UpdatePostById(ctx context.Context, content, id, userId string) (post *models.Post, err error) {
	return postRepoImplementation.UpdatePostById(ctx, content, id, userId)
}

func DeletePostById(ctx context.Context, id, userId string) (err error) {
	return postRepoImplementation.DeletePostById(ctx, id, userId)
}

func ClosePostRepo() error {
	return postRepoImplementation.Close()
}
