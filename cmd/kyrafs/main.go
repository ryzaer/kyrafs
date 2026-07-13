package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryzaer/kyrafs/internal/app"
	"github.com/ryzaer/kyrafs/internal/server"
)

func main() {

	fmt.Println("KyraFS Bootstrap")

	if len(os.Args) > 1 {

		switch os.Args[1] {

		case "init":

			if err := app.Init(); err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println("Ready for online.")
			fmt.Println()
			fmt.Println("Run:")
			fmt.Println("    kyrafs serve")

			return

		case "serve":

			// ========= SERVE =========

			cfg := app.Path("kyrafs.ini")

			if _, err := os.Stat(cfg); os.IsNotExist(err) {
				log.Fatal("CFG001: configuration file not found.\nRun: kyrafs init")
			}

			app.LoadConfig(cfg)

			log.Printf("Listening :%s\n", app.Config.ServerPort())

			s := server.New()

			s.Run(":" + app.Config.ServerPort())

		default:

			fmt.Println("Unknown command.")
			fmt.Println()
			fmt.Println("Available commands:")
			fmt.Println("    kyrafs init")
			fmt.Println("    kyrafs serve")
			return
		}

	} else {

		fmt.Println("Usage:")
		fmt.Println("    kyrafs init")
		fmt.Println("    kyrafs serve")
		return
	}
}

// func main() {

// 	log.Println("KyraFS Bootstrap")

// 	app.LoadConfig(
// 		app.Path("kyrafs.ini"),
// 	)

// 	log.Printf("Listening :%s\n", app.Config.ServerPort())

// 	s := server.New()

// 	if err := s.Run(":" + app.Config.ServerPort()); err != nil {
// 		log.Fatal(err)
// 	}
// }
