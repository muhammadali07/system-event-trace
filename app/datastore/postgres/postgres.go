package postgres

import (
	"database/sql"
	"errors"

	model "github.com/muhammadali07/service-grap-go-api/app/domain/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetPostByID(id int) (model.Post, error) {
	row := r.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)

	var post model.Post
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Post{}, ErrNotFound
		}
		return model.Post{}, err
	}

	return post, nil
}

func (r *PostgresRepository) CreatePost(post model.Post) (model.Post, error) {
	stmt, err := r.db.Prepare("INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING *")
	if err != nil {
		return model.Post{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(post.Title, post.Content)
	err = row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func (r *PostgresRepository) UpdatePost(post model.Post) (model.Post, error) {
	stmt, err := r.db.Prepare("UPDATE posts SET title = $1, content = $2 WHERE id = $3 RETURNING *")
	if err != nil {
		return model.Post{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(post.Title, post.Content, post.ID)
	err = row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func (r *PostgresRepository) DeletePost(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrNotFound = errors.New("post not found")
)
