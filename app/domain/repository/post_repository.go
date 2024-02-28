package repository

import (
  "database/sql"
  "errors"
  "fmt"

  "github.com/muhammadali07/service-grap-go-api/app/domain/model"
)

type PostRepository interface {
  GetAll() ([]model.Post, error)
  GetByID(int) (model.Post, error)
  Create(model.Post) (model.Post, error)
  Update(model.Post) (model.Post, error)
  Delete(int) error
}

type PostRepositoryImpl struct {
  db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
  return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) GetAll() ([]model.Post, error) {
  rows, err := r.db.Query("SELECT * FROM posts")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var posts []model.Post
  for rows.Next() {
    var post model.Post
    err := rows.Scan(&post.ID, &post.Title, &post.Content)
    if err != nil {
      return nil, err
    }
    posts = append(posts, post)
  }

  return posts, nil
}

func (r *PostRepositoryImpl) GetByID(id int) (model.Post, error) {
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

func (r *PostRepositoryImpl) Create(post model.Post) (model.Post, error) {
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

func (r *PostRepositoryImpl) Update(post model.Post) (model.Post, error) {
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

func (r *PostRepositoryImpl) Delete(id int) error {
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
