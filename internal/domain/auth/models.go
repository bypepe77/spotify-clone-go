package auth

type User struct {
	iD       int
	username string
	token    string
}

func (u *User) ID() int {
	return u.iD
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Token() string {
	return u.token
}

func NewUser(iD int, username string, token string) *User {
	return &User{
		iD:       iD,
		username: username,
		token:    token,
	}
}
