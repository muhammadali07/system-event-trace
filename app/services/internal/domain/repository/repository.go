package repository

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateUser(ctx context.Context, name, email string) (int64, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	UpdateUser(ctx context.Context, id int64, name, email string) error
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &userRepository{
		db: db,
	}
}

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *userRepository) CreateUser(ctx context.Context, name, email string) (int64, error) {
	// Implement create user logic here
	return 0, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (User, error) {
	// Implement get user by ID logic here
	return User{}, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int64, name, email string) error {
	// Implement update user logic here
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	// Implement delete user logic here
	return nil
}
