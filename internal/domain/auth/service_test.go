package auth

import (
	"testing"

	"github.com/bypepe77/spotify-clone-go/ent"
	"github.com/bypepe77/spotify-clone-go/internal/infrastructure/api/auth/mock"
	"github.com/stretchr/testify/assert"
)

func TestForT(t *testing.T) {
	t.Log("Test")
}

func Test_Login_Ok(t *testing.T) {
	authRepo := mock.NewAuthRepositoryMock()
	authService := NewAuthService(authRepo)
	expectedUser := &ent.User{
		ID:       1,
		Username: "username",
		Password: "password",
	}

	authRepo.On("GetUserByUsername", "username").Return(expectedUser, nil)

	user, err := authService.Login("username", "password")
	assert.NoError(t, err)

	assert.Equal(t, expectedUser.ID, user.ID())
	assert.Equal(t, expectedUser.Username, user.Username())
}

func Test_Login_Failed_Wrong_Password(t *testing.T) {
	authRepo := mock.NewAuthRepositoryMock()
	authService := NewAuthService(authRepo)
	expectedUser := &ent.User{
		ID:       1,
		Username: "username",
		Password: "No password",
	}

	authRepo.On("GetUserByUsername", "username").Return(expectedUser, nil)

	_, err := authService.Login("username", "password")
	assert.Error(t, err)
}

func Test_Login_Failed_Wrong_Username(t *testing.T) {
	authRepo := mock.NewAuthRepositoryMock()
	authService := NewAuthService(authRepo)
	expectedUser := &ent.User{
		ID:       1,
		Username: "No username",
		Password: "password",
	}

	authRepo.On("GetUserByUsername", "username").Return(expectedUser, nil)

	_, err := authService.Login("username", "password")
	assert.Error(t, err)
}

func Test_Login_Failed_Password_And_Username_Empty(t *testing.T) {
	authRepo := mock.NewAuthRepositoryMock()
	authService := NewAuthService(authRepo)
	expectedUser := &ent.User{
		ID:       1,
		Username: "No username",
		Password: "password",
	}

	authRepo.On("GetUserByUsername", "username").Return(expectedUser, nil)

	_, err := authService.Login("", "")
	assert.Error(t, err)
}

func Test_Register_Ok(t *testing.T) {
	authRepo := mock.NewAuthRepositoryMock()
	authService := NewAuthService(authRepo)

	authRepo.On("GetIfUserExists", "username").Return(false, nil).Once()
	userCreated := &ent.User{
		ID:       1,
		Username: "username",
		Email:    "email@test.com",
		Password: "password",
	}
	authRepo.On("CreateUser", "username", "email@test.com", "password").Return(userCreated, nil).Once()

	user, err := authService.Register("username", "password", "email@test.com")
	assert.NoError(t, err)
	assert.Equal(t, "username", user.Username())
}

func Test_Register_Fail_Username_Exist(t *testing.T) {
	authRepo := mock.NewAuthRepositoryMock()
	authService := NewAuthService(authRepo)

	authRepo.On("GetIfUserExists", "username").Return(true, nil).Once()

	userCreated := &ent.User{
		ID:       1,
		Username: "username",
		Email:    "email@test.com",
		Password: "password",
	}
	authRepo.On("CreateUser", "username", "email@test.com", "password").Return(userCreated, nil).Once()

	_, err := authService.Register("username", "password", "email@test.com")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "username already exists")
}
