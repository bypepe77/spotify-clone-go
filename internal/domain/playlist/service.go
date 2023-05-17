package playlist

type PlayListServiceInterface interface {
	CreatePlaylist()
}

type PlaylistService struct {
	repository PlaylistRepositoryInterface
}

func NewPlaylistService(repository PlaylistRepositoryInterface) PlayListServiceInterface {
	return &PlaylistService{
		repository: repository,
	}
}

func (s *PlaylistService) CreatePlaylist() {
	panic("implement me")
}
