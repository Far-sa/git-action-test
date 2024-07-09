package repository

import (
	"test-cicd/entity"
	"test-cicd/port"
)

type MySQLUserRepository struct {
	db port.Database
}

func NewMySQLUserRepository(db port.Database) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	return err
}

func (r *MySQLUserRepository) FindByEmail(email string) (*entity.User, error) {
	// row := r.db.Get("SELECT id, username, email, password FROM users WHERE email = ?", email)
	// user := &entity.User{}
	// if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, nil
	// 	}
	// 	return nil, err
	// }
	return &entity.User{}, nil
}
