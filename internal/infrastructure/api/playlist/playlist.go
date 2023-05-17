package playlist

import (
	"github.com/bypepe77/spotify-clone-go/internal/domain/playlist"
	"github.com/gin-gonic/gin"
)

type PlayListApiInterface interface{}

type PlaylistApi struct {
	service playlist.PlayListServiceInterface
}

func NewPlaylistApi(service playlist.PlayListServiceInterface) PlayListApiInterface {
	return &PlaylistApi{
		service: service,
	}
}

func (a *PlaylistApi) CreatePlaylist(c *gin.Context) {
	panic("implement me")
}
