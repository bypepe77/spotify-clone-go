package auth

import "github.com/bypepe77/spotify-clone-go/ent"

type AuthRepository interface {
	GetUserByUsername(username string) (*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUserById(id int) (*ent.User, error)
	CreateUser(username string, email string, password string) (*ent.User, error)
}