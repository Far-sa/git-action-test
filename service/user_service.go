package service

import (
	"test-cicd/entity"
	"test-cicd/port"
)

type UserService struct {
	Repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Save(user *entity.User) error {
	return nil
}

func (s *UserService) RegisterUser(username, email, password string) error {
	if _, err := s.Repo.FindByEmail(email); err != nil {
		return err
	}
	user := &entity.User{
		Username: username,
		Email:    email,
		Password: password, // Normally, you should hash the password
	}
	return s.Repo.Save(user)
}
