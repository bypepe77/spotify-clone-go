package playlist

import (
	"github.com/bypepe77/spotify-clone-go/internal/domain/playlist"
	"github.com/gin-gonic/gin"
)

type PlaylistRouter struct {
	routeGroup gin.RouterGroup
	api        PlayListApiInterface
}

func NewPlaylistRouter(r gin.RouterGroup, service playlist.PlayListServiceInterface) *PlaylistRouter {
	return &PlaylistRouter{
		api:        NewPlaylistApi(service),
		routeGroup: r,
	}
}

func (r *PlaylistRouter) RegisterRoutes() {
	panic("implement me")
}
