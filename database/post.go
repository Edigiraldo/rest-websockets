package database

import (
	"context"
	"log"

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

func (repo *PostDatabase) ListPosts(ctx context.Context, userId string, page uint64) ([]*models.Post, error) {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, "SELECT id, content, user_id, created_at FROM posts WHERE user_id = $1 LIMIT $2 OFFSET $3", userId, 5, page*5)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var posts []*models.Post
	for rows.Next() {
		var post = models.Post{}
		if err = rows.Scan(&post.Id, &post.Content, &post.UserId, &post.CreatedAt); err == nil {
			posts = append(posts, &post)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
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

func (repo *PostDatabase) GetPostById(ctx context.Context, id, userId string) (post *models.Post, err error) {
	post = &models.Post{}
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return nil, err
	}
	row := db.QueryRowContext(ctx, "SELECT id, content FROM posts WHERE id = $1 AND user_id = $2", id, userId)

	if err = row.Scan(&post.Id, &post.Content); err != nil {
		return nil, err
	}

	return post, nil
}

func (repo *PostDatabase) UpdatePostById(ctx context.Context, content, id, userId string) (post *models.Post, err error) {
	post = &models.Post{}
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return nil, err
	}
	row := db.QueryRowContext(ctx, "UPDATE posts SET content = $1 WHERE id = $2 AND user_id = $3", content, id, userId)

	if err = row.Scan(&post.Id, &post.Content); err != nil {
		return nil, err
	}

	return post, nil
}

func (repo *PostDatabase) DeletePostById(ctx context.Context, id, userId string) (err error) {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return err
	}

	deleteStatement := `DELETE FROM posts WHERE id = $1 AND user_id = $2;`
	_, err = db.Exec(deleteStatement, id, userId)

	return err
}

func (repo *PostDatabase) Close() error {
	db, err := repo.implementation.GetConnection()
	if err != nil {
		return err
	}

	return db.Close()
}
