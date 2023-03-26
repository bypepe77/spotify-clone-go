package auth

import "github.com/bypepe77/spotify-clone-go/internal/domain/auth"

func toUserResponse(user *auth.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}
}
