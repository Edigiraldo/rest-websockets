package database

import (
	"context"

	"github.com/Edigiraldo/RestWebSockets/models"
)

type PostDatabase struct {
	implementation AbstractDatabase
}

func NewPostDatabase(databaseImplementation AbstractDatabase) *PostDatabase {
	return &PostDatabase{
		implementation: databaseImplementation,
	}
}

func (repo *PostDatabase) InsertPost(ctx context.Context, post *models.Post) (err error) {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, "INSERT INTO posts (id, content, user_id) VALUES ($1, $2, $3)",
		post.Id, post.Content, post.UserId)

	return err
}

func (repo *PostDatabase) Close() error {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return err
	}

	return db.Close()
}
