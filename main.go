package main

import (
	"log"

	"github.com/godoquin/twittor/bd"
	"github.com/godoquin/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
