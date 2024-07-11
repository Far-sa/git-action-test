package port

import (
	"errors"
	"sync"
	"user-service/entity"
)

type MockUserRepository struct {
	users map[string]*entity.User
	mu    sync.Mutex
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (m *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.users[email]
	if !exists {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (m *MockUserRepository) Save(user *entity.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users[user.Email] = user
	return nil
}
