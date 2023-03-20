package bootstrap

import (
	"log"
	"os"
	"regexp"

	"github.com/bypepe77/spotify-clone-go/internal/infrastructure/bootstrap"
	"github.com/joho/godotenv"
)

const enviroment = "dev"
const projectDirName = "spotify-clone-go"

func Run() error {
	if enviroment != "dev" {
		loadEnv()
	}

	config := bootstrap.NewConfig(os.Getenv("APP_NAME"), os.Getenv("PORT"))
	server := bootstrap.NewServer(config)

	return server.Run()
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
