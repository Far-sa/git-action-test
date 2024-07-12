package service_test

import (
	"testing"
	"user-service/entity"
	"user-service/service"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if res := args.Get(0); res != nil {
		return res.(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Save(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestRegisterUser_Success(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	// mockRepo := port.NewMockUserRepository()
	// mockRepo.EXPECT().FindByEmail(email).Return(nil, errors.New("not found"))
	// mockRepo.EXPECT().Save(gomock.Any()).Return(nil)

	m := new(MockUserRepository)
	m.On("FindByEmail", "test@example.com").Return(nil, nil)      // Expect FindByEmail to return nil (not found)
	m.On("Save", mock.AnythingOfType("*entity.User")).Return(nil) // Expect Save to be called with any user

	userService := service.NewUserService(m)

	username := "testuser"
	email := "test@example.com"
	password := "password123"

	err := userService.RegisterUser(username, email, password)

	assert.NoError(t, err, "Unexpected error during registration")
}

// func TestRegisterUser_EmailAlreadyExists(t *testing.T) {
// 	m := new(MockUserRepository) // Create a mock using testify/mock

// 	existingUser := &entity.User{
// 		Username: "existinguser",
// 		Email:    "test@example.com",
// 		Password: "hashedpassword",
// 	}

// 	// Set expectations for mock methods
// 	m.On("FindByEmail", "test@example.com").Return(existingUser, nil) // Expect FindByEmail to return an existing user
// 	m.On("Save", mock.AnythingOfType("*entity.User")).Times(0)        // Don't expect Save to be called

// 	service := service.NewUserService(m) // Inject the mock into the service

// 	username := "testuser"
// 	email := "test@example.com"
// 	password := "password123"

// 	err := service.RegisterUser(username, email, password)

// 	// Use assert from testify/mock for assertions
// 	assert.Error(t, err, "Expected error for existing email")
// 	// You can optionally check for specific error type or message
// 	// assert.Equal(t, "email already exists", err.Error())
// }
