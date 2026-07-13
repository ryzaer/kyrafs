package app

import (
	"fmt"
	"os"
)

const defaultConfig = `[server]
port=8001

[storage]
access_key=
reserved_free_space=

[volume:main]
path=./storage
`

func Init() error {

	// create temp
	if err := os.MkdirAll(Path("temp"), 0755); err != nil {
		return err
	}

	// create storage
	if err := os.MkdirAll(Path("storage"), 0755); err != nil {
		return err
	}

	cfg := Path("kyrafs.ini")

	if _, err := os.Stat(cfg); os.IsNotExist(err) {

		if err := os.WriteFile(
			cfg,
			[]byte(defaultConfig),
			0644,
		); err != nil {
			return err
		}

		fmt.Println("Configuration created successfully.")
	} else {
		fmt.Println("Configuration already exists.")
	}

	return nil
}
