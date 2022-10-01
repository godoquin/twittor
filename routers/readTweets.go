package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/godoquin/twittor/database"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	response, ok := database.ReadTweet(ID, pag)
	if !ok {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
