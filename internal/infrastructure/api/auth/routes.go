package auth

import (
	"github.com/bypepe77/spotify-clone-go/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	service    auth.AuthServiceInterface
	routeGroup gin.RouterGroup
	api        AuthApiInterface
}

func NewAuthRouter(service auth.AuthServiceInterface, r gin.RouterGroup) *AuthRouter {
	return &AuthRouter{
		service:    service,
		api:        NewAuthApi(service),
		routeGroup: r,
	}
}

func (r *AuthRouter) RegisterRoutes() {
	r.routeGroup.POST("/login", r.api.Login)
}
