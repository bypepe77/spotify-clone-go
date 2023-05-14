package auth

import (
	"context"

	"github.com/bypepe77/spotify-clone-go/ent"
	"github.com/bypepe77/spotify-clone-go/ent/user"
)

type authRepository struct {
	db *ent.Client
}

type AuthRepositoryInterface interface {
	GetUserByUsername(username string) (*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUserById(id int) (*ent.User, error)
	CreateUser(username, email, password string) (*ent.User, error)
	GetIfUserExists(username string) (bool, error)
}

func NewAuthRepository(db *ent.Client) AuthRepositoryInterface {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) GetUserByUsername(username string) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.Username(username)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) GetIfUserExists(username string) (bool, error) {
	exist, err := r.db.User.Query().Where(user.Username(username)).Exist(context.Background())
	if err != nil {
		return true, err
	}
	return exist, nil
}

func (r *authRepository) GetUserByEmail(email string) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.Email(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) GetUserById(id int) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) CreateUser(username, email, password string) (*ent.User, error) {
	user, err := r.db.User.Create().SetUsername(username).SetPassword(password).SetEmail(email).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
