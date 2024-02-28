package repository

import (
	"database/sql"

	model "github.com/muhammadali07/service-grap-go-api/app/domain/models"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(int) (model.User, error)
	Create(model.User) (model.User, error)
	// Update(model.User) (model.User, error)
	// Delete(int) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAll() (response []model.User, err error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetByID(id int) (response model.User, err error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var user model.User
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return response, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Create(user model.User) (response model.User, err error) {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING *")
	if err != nil {
		return response, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Name, user.Email)
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return response, err
	}
	return user, nil
}
