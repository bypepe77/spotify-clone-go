package auth

func toUserModel(ID int, username, token string) *User {
	return &User{
		ID:       ID,
		Username: username,
		Token:    token,
	}
}
