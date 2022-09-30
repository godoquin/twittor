package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/godoquin/twittor/middleware"
	"github.com/godoquin/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middleware.CheckBD(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/seeprofile", middleware.CheckBD(middleware.CheckJWT(routers.SeeProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
