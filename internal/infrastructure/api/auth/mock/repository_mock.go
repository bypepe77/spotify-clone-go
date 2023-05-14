package mock

import (
	"github.com/bypepe77/spotify-clone-go/ent"
	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func NewAuthRepositoryMock() *AuthRepositoryMock {
	return &AuthRepositoryMock{}
}

func (m *AuthRepositoryMock) GetUserByUsername(username string) (*ent.User, error) {
	args := m.Called(username)
	return args.Get(0).(*ent.User), args.Error(1)
}

func (m *AuthRepositoryMock) GetIfUserExists(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func (m *AuthRepositoryMock) GetUserByEmail(email string) (*ent.User, error) {
	args := m.Called(email)
	return args.Get(0).(*ent.User), args.Error(1)
}

func (m *AuthRepositoryMock) GetUserById(id int) (*ent.User, error) {
	args := m.Called(id)
	return args.Get(0).(*ent.User), args.Error(1)
}

func (m *AuthRepositoryMock) CreateUser(username, email, password string) (*ent.User, error) {
	args := m.Called(username, email, password)
	return args.Get(0).(*ent.User), args.Error(1)
}
