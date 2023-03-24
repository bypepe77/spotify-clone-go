package auth

import (
	"errors"

	"github.com/bypepe77/spotify-clone-go/internal/infrastructure/middlewares"
)

type AuthServiceInterface interface {
	Login(username string, password string) (*User, error)
	Register(username string, password string) (*User, error)
}

type AuthService struct {
	repository AuthRepository
	jwtService middlewares.JWTService
}

func NewAuthService(repository AuthRepository) AuthServiceInterface {
	return &AuthService{
		repository: repository,
		jwtService: middlewares.JWTAuthService(),
	}
}

func (s *AuthService) Login(username string, password string) (*User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}

	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// TODO: Hash password
	if user.Password != password || user.Username != username {
		return nil, errors.New("invalid password or username")
	}

	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return toUserModel(user.ID, user.Username, token), nil
}

func (s *AuthService) Register(username string, password string) (*User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}

	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("username already exists")
	}

	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return toUserModel(user.ID, user.Username, token), nil
}
