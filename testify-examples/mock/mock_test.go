package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepository interface {
	CreateUser(user *User) (int, error)
	GetUserById(id int) (*User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name string, age int) (*User, error) {
	user := &User{Name: name, Age: age}
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (s *UserService) GetUserById(id int) (*User, error) {
	return s.repo.GetUserById(id)
}

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user *User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *UserRepositoryMock) GetUserById(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
	repo := new(UserRepositoryMock)
	service := NewUserService(repo)

	user := &User{Name: "Alice", Age: 30}
	repo.On("CreateUser", user).Return(1, nil)

	createdUser, err := service.CreateUser(user.Name, user.Age)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdUser.ID)
	assert.Equal(t, "Alice", createdUser.Name)
	assert.Equal(t, 30, createdUser.Age)

	repo.AssertExpectations(t)
}

func TestUserService_GetUserById(t *testing.T) {
	repo := new(UserRepositoryMock)
	service := NewUserService(repo)

	user := &User{ID: 1, Name: "Alice", Age: 30}
	repo.On("GetUserById", 1).Return(user, nil)

	foundUser, err := service.GetUserById(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, foundUser.ID)
	assert.Equal(t, "Alice", foundUser.Name)
	assert.Equal(t, 30, foundUser.Age)

	repo.AssertExpectations(t)
}
