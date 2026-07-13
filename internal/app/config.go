package app

import (
	"log"

	"github.com/ryzaer/kyrafs/internal/config"
)

var Config *config.Config

func LoadConfig(path string) {

	cfg, err := config.Load(path)

	if err != nil {
		log.Fatal(err)
	}

	Config = cfg
}
