package main

import (
	"itmo/configs"
	"itmo/server"
	"log"
)

func main() {

	if err := configs.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run("8080"); err != nil {
		log.Fatalf("%s", err)
	}

}
