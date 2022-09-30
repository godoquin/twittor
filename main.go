package main

import (
	"log"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/handlers"
)

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
