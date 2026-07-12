package main

import (
	"log"

	"github.com/ryzaer/kyrafs/internal/app"
	"github.com/ryzaer/kyrafs/internal/config"
	"github.com/ryzaer/kyrafs/internal/server"
)

func main() {

	log.Println("KyraFS Bootstrap")

	cfg, err := config.Load(
		app.Path("kyrafs.ini"),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening :%s\n", cfg.ServerPort())

	s := server.New()

	if err := s.Run(":" + cfg.ServerPort()); err != nil {
		log.Fatal(err)
	}
}
