package auth

type AuthServiceInterface interface {
}

type AuthService struct {
	repository AuthRepository
}

func NewAuthService(repository AuthRepository) AuthServiceInterface {
	return &AuthService{
		repository: repository,
	}
}
