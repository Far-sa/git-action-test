package repository

import (
	"database/sql"
	"errors"
	"user-service/entity"
	"user-service/port"
)

type MySQLUserRepository struct {
	db port.Database
}

func NewMySQLUserRepository(db port.Database) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user *entity.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	return err
}

func (r *MySQLUserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	if err := r.db.Get(user, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
