package port

import "test-cicd/entity"

type UserService interface {
	Save(user *entity.User) error
	RegisterUser(username, email, password string) error
}
