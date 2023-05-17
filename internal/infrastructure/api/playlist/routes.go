package playlist

import (
	"net/http"

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
	r.routeGroup.GET("/playlists", getFakePlaylist)
}

func getFakePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": " Implement me!",
	},
	)
}
