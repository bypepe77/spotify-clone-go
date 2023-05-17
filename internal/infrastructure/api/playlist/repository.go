package playlist

import "github.com/bypepe77/spotify-clone-go/ent"

type playlistRepository struct {
	db *ent.Client
}

type PlaylistRepositoryInterface interface {
}

func NewPlaylistRepository(db *ent.Client) PlaylistRepositoryInterface {
	return &playlistRepository{
		db: db,
	}
}

func (r *playlistRepository) CreatePlaylist() {
	panic("implement me")
}
