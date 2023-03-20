package main

import (
	"log"

	"github.com/bypepe77/spotify-clone-go/cmd/bootstrap"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
