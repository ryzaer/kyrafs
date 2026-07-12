package main

import (
	"log"

	"github.com/ryzaer/kyrafs/internal/server"
)

func main() {

	log.Println("KyraFS Bootstrap")

	s := server.New()

	if err := s.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
