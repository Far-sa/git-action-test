package port

import "user-service/entity"

type UserService interface {
	Save(user *entity.User) error
	RegisterUser(username, email, password string) error
}
