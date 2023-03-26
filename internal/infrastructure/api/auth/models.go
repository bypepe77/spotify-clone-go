package auth

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
