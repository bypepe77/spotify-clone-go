package auth

import (
	"net/http"

	"github.com/bypepe77/spotify-clone-go/internal/domain/auth"
	"github.com/bypepe77/spotify-clone-go/internal/infrastructure/responses"
	"github.com/gin-gonic/gin"
)

type AuthApiInterface interface {
	Login(c *gin.Context)
}

type AuthApi struct {
	service auth.AuthServiceInterface
}

func NewAuthApi(service auth.AuthServiceInterface) AuthApiInterface {
	return &AuthApi{
		service: service,
	}
}

func (a *AuthApi) Login(c *gin.Context) {
	var payload *UserInput
	err := c.BindJSON(&payload)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, "Invalid payload")
		return
	}

	if payload.Username == "" || payload.Password == "" {
		responses.Error(http.StatusBadRequest, c, "Username or password is empty")
		return
	}

	user, err := a.service.Login(payload.Username, payload.Password)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, err.Error())
		return
	}

	responses.Success(http.StatusOK, c, toUserResponse(user))
}
