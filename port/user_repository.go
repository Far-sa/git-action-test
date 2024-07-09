package port

import (
	"database/sql"
	"test-cicd/entity"
)

type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
}

type UserRepository interface {
	Save(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
