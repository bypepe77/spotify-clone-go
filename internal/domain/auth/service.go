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

func (s *AuthService) Login(username string, password string) (string, error) {
	_, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	return "", nil

}
