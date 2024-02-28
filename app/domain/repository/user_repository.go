package repository

import (
	"database/sql"
)

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(int) (User, error)
	Create(User) (User, error)
	Update(User) (User, error)
	Delete(int) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetByID(id int) (User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Create(user User) (User, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING *")
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Name, user.Email)
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
