package auth

import "github.com/bypepe77/spotify-clone-go/ent"

type authService struct {
	db *ent.Client
}

type AuthServiceInterface interface {
}

func NewAuthService(db *ent.Client) AuthServiceInterface {
	return &authService{
		db: db,
	}
}
