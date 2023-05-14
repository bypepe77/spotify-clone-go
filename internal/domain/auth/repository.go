package auth

import "github.com/bypepe77/spotify-clone-go/ent"

type AuthRepository interface {
	GetUserByUsername(username string) (*ent.User, error)
	GetIfUserExists(username string) (bool, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUserById(id int) (*ent.User, error)
	CreateUser(username, email, password string) (*ent.User, error)
}
