package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bypepe77/spotify-clone-go/ent"
	"github.com/bypepe77/spotify-clone-go/internal/domain/auth"
	"github.com/bypepe77/spotify-clone-go/internal/domain/playlist"
	authApi "github.com/bypepe77/spotify-clone-go/internal/infrastructure/api/auth"
	playlistApi "github.com/bypepe77/spotify-clone-go/internal/infrastructure/api/playlist"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type server struct {
	config *Config
	engine *gin.Engine
	db     *ent.Client
}

func NewServer(config *Config, db *ent.Client) *server {
	return &server{
		config: config,
		engine: gin.New(),
		db:     db,
	}
}

func (s *server) buildConnectionString() string {
	if s.config.Port == "" {
		s.config.Port = "8080"
	}

	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

func (s *server) awsConnectionCommand() (*s3.S3, error) {
	// move this to aws folder inside bootstrap
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)

	return svc, nil
}

func (s *server) Run() error {
	s.engine.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization", "content-type"},
		AllowHeaders:     []string{"Authorization", "content-type "},
	}))
	/*
		svc, err := s.awsConnectioCommand()
		if err != nil {
			return err
		}

		result, err := svc.ListBuckets(nil)
		if err != nil {
			return err
		}

		fmt.Println("Buckets:")
		for _, b := range result.Buckets {
			fmt.Printf("* %s created on %s\n",
				aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
		}
	*/
	s.registerRoutes()
	return s.engine.Run()
}

func (s *server) registerRoutes() {
	// Initialize Basic routes
	s.engine.GET("/", healthCheck)

	// Initialize Auth routes
	authRepository := authApi.NewAuthRepository(s.db)
	authService := auth.NewAuthService(authRepository)
	authRouter := authApi.NewAuthRouter(authService, *s.engine.Group("/auth"))
	authRouter.RegisterRoutes()

	// Initialize Playlist routes
	playlistRepository := playlistApi.NewPlaylistRepository(s.db)
	playlistService := playlist.NewPlaylistService(playlistRepository)
	playlistRouter := playlistApi.NewPlaylistRouter(*s.engine.Group("/playlist"), playlistService)
	playlistRouter.RegisterRoutes()

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong", "status": true})
}
